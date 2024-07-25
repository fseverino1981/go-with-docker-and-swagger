package service

import (
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/model"
	"go-with-docker-and-swagger/src/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_CreateUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_user_already_exists_returns_error",
		func(t *testing.T) {
			id := primitive.NewObjectID().Hex()

			userDomain := model.NewUserDomain("teste@teste.com", "test", "Teste", 20)
			userDomain.SetID(id)

			repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

			user, err := service.CreateUserServices(userDomain)

			assert.Nil(t, user)
			assert.NotNil(t, err)
			assert.EqualValues(t, "Email is already registered in another account", err.Message)

		})

	t.Run("when_user_is_not_returns", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("teste@teste.com", "test", "Teste", 20)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(nil, rest_err.NewInternalServerError("error try to create user"))

		user, err := service.CreateUserServices(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, "error try to create user", err.Message)

	})

	t.Run("when_user_is_not_registered_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("teste@teste.com", "test", "Teste", 20)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, err := service.CreateUserServices(userDomain)

		assert.NotNil(t, user)
		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetEmail(), user.GetEmail())
		assert.EqualValues(t, userDomain.GetName(), user.GetName())
		assert.EqualValues(t, userDomain.GetAge(), user.GetAge())
		assert.EqualValues(t, userDomain.GetID(), user.GetID())
		assert.EqualValues(t, userDomain.GetPassword(), user.GetPassword())
	})

}
