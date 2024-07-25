package controller

import (
	"encoding/json"
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/controller/model/request"
	"go-with-docker-and-swagger/src/model"
	"go-with-docker-and-swagger/src/tests/mocks"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("validation_body_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserUpdateRequest{
			Name: "",
			Age:  -1,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "PUT", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_userId_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserUpdateRequest{
			Name: "teste user",
			Age:  10,
		}

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "test",
			},
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, param, url.Values{}, "PUT", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("object_is_valid_but_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		userRequest := request.UserUpdateRequest{
			Name: "Test User test",
			Age:  10,
		}

		domain := model.NewUserUpdateDomain(
			userRequest.Name,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().UpdateUserServices(id, domain).Return(rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "PUT", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("object_is_valid_and_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		userRequest := request.UserUpdateRequest{
			Name: "Test User test",
			Age:  10,
		}

		domain := model.NewUserUpdateDomain(
			userRequest.Name,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().UpdateUserServices(id, domain).Return(nil)

		MakeRequest(context, param, url.Values{}, "PUT", stringReader)
		controller.UpdateUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})

}
