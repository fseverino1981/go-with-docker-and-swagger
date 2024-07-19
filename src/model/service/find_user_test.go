package service

import (
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/model"
	"go-with-docker-and-swagger/src/test/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.uber.org/mock/gomock"
)

func TestUserDomainService_FindUserByIDServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success",
		func(t *testing.T) {
			id := primitive.NewObjectID().Hex()

			userDomain := model.NewUserDomain("teste@teste.com", "test", "Teste", 20)
			userDomain.SetID(id)

			repository.EXPECT().FindUserByID(id).Return(userDomain, nil)

			userDomainReturn, err := service.FindUserByIDServices(id)

			assert.Nil(t, err)
			assert.EqualValues(t, userDomainReturn.GetID(), id)
			assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
			assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
			assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
			assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())

		})

	t.Run("when_does_not_exists_an_user_returns_error",
		func(t *testing.T) {
			id := primitive.NewObjectID().Hex()

			repository.EXPECT().FindUserByID(id).Return(nil, rest_err.NewNotFoundError("user not found"))
			userDomainReturn, err := service.FindUserByIDServices(id)

			assert.Nil(t, userDomainReturn)
			assert.NotNil(t, err)
			assert.EqualValues(t, err.Message, "user not found")
		})
}

func TestUserDomainService_FindUserByEmailServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success",
		func(t *testing.T) {
			id := primitive.NewObjectID().Hex()
			email := "teste@teste.com"

			userDomain := model.NewUserDomain("teste@teste.com", "test", "Teste", 20)
			userDomain.SetID(id)

			repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

			userDomainReturn, err := service.FindUserByEmailServices(email)

			assert.Nil(t, err)
			assert.EqualValues(t, userDomainReturn.GetID(), id)
			assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
			assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
			assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
			assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())

		})

	t.Run("when_does_not_exists_an_user_returns_error",
		func(t *testing.T) {
			email := "teste@teste.com"

			repository.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("user not found"))
			userDomainReturn, err := service.FindUserByEmailServices(email)

			assert.Nil(t, userDomainReturn)
			assert.NotNil(t, err)
			assert.EqualValues(t, err.Message, "user not found")
		})
}

func TestUserDomainService_FindUserByEmailAndPasswordServices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_returns_success",
		func(t *testing.T) {
			id := primitive.NewObjectID().Hex()
			email := "teste@teste.com"
			password := "test"

			userDomain := model.NewUserDomain("teste@teste.com", "test", "Teste", 20)
			userDomain.SetID(id)

			repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(userDomain, nil)

			userDomainReturn, err := service.findUserByEmailAndPasswordServices(email, password)

			assert.Nil(t, err)
			assert.EqualValues(t, userDomainReturn.GetID(), id)
			assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
			assert.EqualValues(t, userDomainReturn.GetName(), userDomain.GetName())
			assert.EqualValues(t, userDomainReturn.GetPassword(), userDomain.GetPassword())
			assert.EqualValues(t, userDomainReturn.GetAge(), userDomain.GetAge())

		})

	t.Run("when_does_not_exists_an_user_returns_error",
		func(t *testing.T) {
			email := "teste@teste.com"
			password := "test"

			repository.EXPECT().FindUserByEmailAndPassword(email, password).Return(nil, rest_err.NewNotFoundError("user not found"))
			userDomainReturn, err := service.findUserByEmailAndPasswordServices(email, password)

			assert.Nil(t, userDomainReturn)
			assert.NotNil(t, err)
			assert.EqualValues(t, err.Message, "user not found")
		})
}
