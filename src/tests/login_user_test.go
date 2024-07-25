package tests

import (
	"encoding/json"
	"fmt"
	"go-with-docker-and-swagger/src/controller/model/request"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {

	t.Run("user_and_password_is_valid", func(t *testing.T) {

		recorderCreate := httptest.NewRecorder()
		ctxCreate := GetTestGinContext(recorderCreate)

		recorderLogin := httptest.NewRecorder()
		ctxLogin := GetTestGinContext(recorderLogin)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d@", rand.Int())

		fmt.Println("Email:", email)
		fmt.Println("PassWord:", password)

		userCreateRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test User",
			Age:      20,
		}

		bCreate, _ := json.Marshal(userCreateRequest)
		stringReaderCreate := io.NopCloser(strings.NewReader(string(bCreate)))
		MakeRequest(ctxCreate, []gin.Param{}, url.Values{}, "POST", stringReaderCreate)
		UserController.CreateUser(ctxCreate)

		userLoginRequest := request.UserLogin{
			Email:    email,
			Password: password,
		}

		bLogin, _ := json.Marshal(userLoginRequest)
		stringReaderLogin := io.NopCloser(strings.NewReader(string(bLogin)))
		MakeRequest(ctxLogin, []gin.Param{}, url.Values{}, "POST", stringReaderLogin)
		UserController.LoginUser(ctxLogin)

		assert.EqualValues(t, http.StatusOK, recorderLogin.Result().StatusCode)
		assert.NotEmpty(t, recorderLogin.Result().Header.Get("Authorization"))

	})

	t.Run("user_and_password_is_not_valid", func(t *testing.T) {

		recorderCreate := httptest.NewRecorder()
		ctxCreate := GetTestGinContext(recorderCreate)

		recorderLogin := httptest.NewRecorder()
		ctxLogin := GetTestGinContext(recorderLogin)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d@", rand.Int())

		fmt.Println("Email:", email)
		fmt.Println("PassWord:", password)

		userCreateRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test User",
			Age:      20,
		}

		bCreate, _ := json.Marshal(userCreateRequest)
		stringReaderCreate := io.NopCloser(strings.NewReader(string(bCreate)))
		MakeRequest(ctxCreate, []gin.Param{}, url.Values{}, "POST", stringReaderCreate)
		UserController.CreateUser(ctxCreate)

		userLoginRequest := request.UserLogin{
			Email:    "test_invalid@tes.com",
			Password: "p@sswordincorrect",
		}

		bLogin, _ := json.Marshal(userLoginRequest)
		stringReaderLogin := io.NopCloser(strings.NewReader(string(bLogin)))
		MakeRequest(ctxLogin, []gin.Param{}, url.Values{}, "POST", stringReaderLogin)
		UserController.LoginUser(ctxLogin)

		assert.EqualValues(t, http.StatusForbidden, recorderLogin.Result().StatusCode)
		assert.Empty(t, recorderLogin.Result().Header.Get("Authorization"))

	})

}
