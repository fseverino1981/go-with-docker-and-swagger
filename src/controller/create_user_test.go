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
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("validation_got_error",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context := GetTestGinContext(recorder)

			userRequest := request.UserRequest{
				Email:    "test@test",
				Password: "teste",
				Name:     "Teste User",
				Age:      0,
			}

			b, _ := json.Marshal(userRequest)
			stringReader := io.NopCloser(strings.NewReader(string(b)))

			MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
			controller.CreateUser(context)

			assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
		})

	t.Run("object_is_valid_but_service_returns_error",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context := GetTestGinContext(recorder)

			userRequest := request.UserRequest{
				Email:    "test@test.com",
				Password: "test@ndo",
				Name:     "Test User",
				Age:      20,
			}

			domain := model.NewUserDomain(
				userRequest.Email,
				userRequest.Password,
				userRequest.Name,
				userRequest.Age,
			)

			b, _ := json.Marshal(userRequest)
			stringReader := io.NopCloser(strings.NewReader(string(b)))

			service.EXPECT().CreateUserServices(domain).Return(
				nil, rest_err.NewInternalServerError("error test"))

			MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
			controller.CreateUser(context)

			assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)

		})
	t.Run("object_is_valid_but_service_returns_success",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context := GetTestGinContext(recorder)

			userRequest := request.UserRequest{
				Email:    "test@test.com",
				Password: "test@ndo",
				Name:     "Test User",
				Age:      20,
			}

			domain := model.NewUserDomain(
				userRequest.Email,
				userRequest.Password,
				userRequest.Name,
				userRequest.Age,
			)

			b, _ := json.Marshal(userRequest)
			stringReader := io.NopCloser(strings.NewReader(string(b)))

			service.EXPECT().CreateUserServices(domain).Return(
				domain, nil)

			MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
			controller.CreateUser(context)

			assert.EqualValues(t, http.StatusOK, recorder.Code)
			assert.EqualValues(t, userRequest.Email, domain.GetEmail())
			assert.EqualValues(t, userRequest.Password, domain.GetPassword())
			assert.EqualValues(t, userRequest.Name, domain.GetName())
			assert.EqualValues(t, userRequest.Age, domain.GetAge())

		})

}
