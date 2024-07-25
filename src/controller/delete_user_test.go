package controller

import (
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/tests/mocks"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("id_is_invalid_returns_error",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context := GetTestGinContext(recorder)

			param := []gin.Param{
				{
					Key:   "userId",
					Value: "TEST_ERROR",
				},
			}

			MakeRequest(context, param, url.Values{}, "GET", nil)
			controller.DeleteUser(context)

			assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

		})

	t.Run("id_is_valid_service_ returns_error",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context := GetTestGinContext(recorder)

			id := primitive.NewObjectID().Hex()

			param := []gin.Param{
				{
					Key:   "userId",
					Value: id,
				},
			}

			service.EXPECT().DeleteUserServices(id).Return(
				rest_err.NewInternalServerError("error test"))

			MakeRequest(context, param, url.Values{}, "DELETE", nil)
			controller.DeleteUser(context)

			assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)

		})

	t.Run("id_is_valid_service_ returns_success",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context := GetTestGinContext(recorder)

			id := primitive.NewObjectID().Hex()

			param := []gin.Param{
				{
					Key:   "userId",
					Value: id,
				},
			}

			service.EXPECT().DeleteUserServices(id).Return(nil)

			MakeRequest(context, param, url.Values{}, "DELETE", nil)
			controller.DeleteUser(context)

			assert.EqualValues(t, http.StatusOK, recorder.Code)

		})
}
