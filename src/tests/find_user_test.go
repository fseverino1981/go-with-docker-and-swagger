package test

import (
	"context"
	"go-with-docker-and-swagger/src/controller"
	"go-with-docker-and-swagger/src/model/repository"
	"go-with-docker-and-swagger/src/model/service"
	"go-with-docker-and-swagger/src/tests/connection"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UserController controller.UserControllerInterface
	Database       *mongo.Database
)

func TestMain(m *testing.M) {
	err := os.Setenv("MONGODB_USER_DB", "test_user")
	err = os.Setenv("MONGODB_USER_COLLECTION", "test_user_collection")
	if err != nil {
		return
	}

	closeConnection := func() {}
	Database, closeConnection = connection.OpenConnection()

	repo := repository.NewUserRepository(Database)
	userService := service.NewUserDomainService(repo)
	UserController = controller.NewUserControllerInterface(userService)

	defer func() {
		os.Clearenv()
		closeConnection()
	}()

	os.Exit(m.Run())

}

func TestFindUserByEmail(t *testing.T) {

	t.Run("user_not_found_with_this_email",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context := GetTestGinContext(recorder)

			param := []gin.Param{
				{
					Key:   "userEmail",
					Value: "test@test.com",
				},
			}

			MakeRequest(context, param, url.Values{}, "GET", nil)
			UserController.FindUserByEmail(context)

			assert.EqualValues(t, http.StatusNotFound, recorder.Code)

		})

	t.Run("user_found_with_especified_email",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			ctxt := GetTestGinContext(recorder)
			ID := primitive.NewObjectID().Hex()

			_, err := Database.
				Collection(os.Getenv("MONGODB_USER_COLLECTION")).
				InsertOne(context.Background(), bson.M{"_id": ID, "name": t.Name(), "email": "flavio@test.com"})
			if err != nil {
				t.Fatal(err)
				return
			}

			param := []gin.Param{
				{
					Key:   "userEmail",
					Value: "flavio@test.com",
				},
			}

			MakeRequest(ctxt, param, url.Values{}, "GET", nil)
			UserController.FindUserByEmail(ctxt)

			assert.EqualValues(t, http.StatusOK, recorder.Code)

		})

}

func TestFindUserByID(t *testing.T) {

	t.Run("user_not_found_with_this_id",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context := GetTestGinContext(recorder)
			ID := primitive.NewObjectID().Hex()

			param := []gin.Param{
				{
					Key:   "userId",
					Value: ID,
				},
			}

			MakeRequest(context, param, url.Values{}, "GET", nil)
			UserController.FindUserById(context)

			assert.EqualValues(t, http.StatusNotFound, recorder.Code)

		})

	t.Run("user_found_with_especified_id",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			ctxt := GetTestGinContext(recorder)
			ID := primitive.NewObjectID()

			_, err := Database.
				Collection(os.Getenv("MONGODB_USER_COLLECTION")).
				InsertOne(context.Background(), bson.M{"_id": ID, "name": t.Name(), "email": "flavio@test.com"})
			if err != nil {
				t.Fatal(err)
				return
			}

			param := []gin.Param{
				{
					Key:   "userId",
					Value: ID.Hex(),
				},
			}

			MakeRequest(ctxt, param, url.Values{}, "GET", nil)
			UserController.FindUserById(ctxt)

			assert.EqualValues(t, http.StatusOK, recorder.Code)

		})

}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param

	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
