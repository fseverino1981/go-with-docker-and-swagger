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

func TestUserDomainService_UpdateUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)
	t.Run("when_user_updates_returns_success",
		func(t *testing.T) {
			id := primitive.NewObjectID().Hex()

			userDomain := model.NewUserDomain("teste@teste.com", "test", "Teste", 20)
			userDomain.SetID(id)

			repository.EXPECT().UpdateUser(id, userDomain).
				Return(nil)

			err := service.UpdateUserServices(id, userDomain)

			assert.Nil(t, err)

		})

	t.Run("when_user_updates_returns_error",
		func(t *testing.T) {
			id := primitive.NewObjectID().Hex()

			userDomain := model.NewUserDomain("teste@teste.com", "test", "Teste", 20)
			userDomain.SetID(id)

			repository.EXPECT().UpdateUser(id, userDomain).
				Return(rest_err.NewInternalServerError("error try to update user"))

			err := service.UpdateUserServices(id, userDomain)

			assert.NotNil(t, err)
			assert.EqualValues(t, "error try to update user", err.Message)

		})

}
