package service

import (
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/model"
	"go-with-docker-and-swagger/src/tests/mocks"
	"os"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_LoginUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := &userDomainService{repository}
	t.Run("when_calling_repository_returtns_error",
		func(t *testing.T) {
			id := primitive.NewObjectID().Hex()

			userDomain := model.NewUserDomain("teste@teste.com", "test", "Teste", 20)
			userDomain.SetID(id)

			userDomainMock := model.NewUserDomain(
				userDomain.GetEmail(),
				userDomain.GetPassword(),
				userDomain.GetName(),
				userDomain.GetAge())
			userDomainMock.EncryptPassword()

			repository.EXPECT().FindUserByEmailAndPassword(
				userDomain.GetEmail(), userDomainMock.GetPassword()).
				Return(nil, rest_err.NewInternalServerError("error trying to find user by email and password"))

			user, token, err := service.LoginUserServices(userDomain)

			assert.NotNil(t, err)
			assert.Empty(t, token)
			assert.Nil(t, user)
			assert.EqualValues(t, "error trying to find user by email and password", err.Message)

		})

	t.Run("when_calling_create_token_returns_error",
		func(t *testing.T) {

			userDomainMock := mocks.NewMockUserDomainInterface(ctrl)

			userDomainMock.EXPECT().GetEmail().Return("teste@teste.com")
			userDomainMock.EXPECT().GetPassword().Return("test")
			userDomainMock.EXPECT().EncryptPassword()

			userDomainMock.EXPECT().GeneralToken().Return("",
				rest_err.NewInternalServerError("error trying to create token"))

			repository.EXPECT().FindUserByEmailAndPassword(
				"teste@teste.com", "test").Return(userDomainMock, nil)

			user, token, err := service.LoginUserServices(userDomainMock)

			assert.NotNil(t, err)
			assert.Empty(t, token)
			assert.Nil(t, user)
			assert.EqualValues(t, "error trying to create token", err.Message)

		})

	t.Run("when_user_and_password)is_valid_return_success",
		func(t *testing.T) {
			id := primitive.NewObjectID().Hex()
			secret := "test"
			os.Setenv("JWT_SECRET_KEY", secret)
			defer os.Clearenv()

			userDomain := model.NewUserDomain("teste@teste.com", "test", "Teste", 20)
			userDomain.SetID(id)

			repository.EXPECT().FindUserByEmailAndPassword(
				userDomain.GetEmail(), gomock.Any()).
				Return(userDomain, nil)

			user, token, err := service.LoginUserServices(userDomain)

			assert.Nil(t, err)
			assert.NotEmpty(t, token)
			assert.EqualValues(t, userDomain.GetEmail(), user.GetEmail())
			assert.EqualValues(t, userDomain.GetName(), user.GetName())
			assert.EqualValues(t, userDomain.GetAge(), user.GetAge())
			assert.EqualValues(t, userDomain.GetID(), user.GetID())
			assert.EqualValues(t, userDomain.GetPassword(), user.GetPassword())

			tokenReturn, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
					return []byte(secret), nil
				}

				return nil, rest_err.NewBadRequestError("Invalid token")
			})

			_, ok := tokenReturn.Claims.(jwt.MapClaims)
			if !ok || !tokenReturn.Valid {
				t.FailNow()
				return
			}
		})

}
