package service

import (
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_DeleteUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)
	t.Run("when_user_deletes_returns_success",
		func(t *testing.T) {
			id := primitive.NewObjectID().Hex()

			repository.EXPECT().DeleteUser(id).
				Return(nil)

			err := service.DeleteUserServices(id)

			assert.Nil(t, err)

		})

	t.Run("when_user_deletes_returns_error",
		func(t *testing.T) {
			id := primitive.NewObjectID().Hex()

			repository.EXPECT().DeleteUser(id).
				Return(rest_err.NewInternalServerError("error try to update user"))

			err := service.DeleteUserServices(id)

			assert.NotNil(t, err)
			assert.EqualValues(t, "error try to update user", err.Message)

		})

}
