package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"go-with-docker-and-swagger/src/controller/model/request"
	"go-with-docker-and-swagger/src/model/entity"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUserTest(t *testing.T) {

	t.Run("user_already_registred_with_this_email",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			ctx := GetTestGinContext(recorder)
			email := fmt.Sprintf("%d@test.com", rand.Int())

			_, err := Database.
				Collection(os.Getenv("MONGODB_USER_COLLECTION")).
				InsertOne(context.Background(), bson.M{"name": t.Name(), "email": email})
			if err != nil {
				t.Fatal(err)
				return
			}

			userRequest := request.UserRequest{
				Email:    email,
				Password: "test@ndo",
				Name:     "Test User",
				Age:      20,
			}

			b, _ := json.Marshal(userRequest)
			stringReader := io.NopCloser(strings.NewReader(string(b)))

			MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
			UserController.CreateUser(ctx)

			assert.EqualValues(t, http.StatusBadRequest, recorder.Code)

		})

	t.Run("user_is_not_registred_in_database",
		func(t *testing.T) {
			recorder := httptest.NewRecorder()
			ctx := GetTestGinContext(recorder)
			email := fmt.Sprintf("%d@test.com", rand.Int())

			_, err := Database.
				Collection(os.Getenv("MONGODB_USER_COLLECTION")).
				InsertOne(context.Background(), bson.M{"name": t.Name(), "email": email})
			if err != nil {
				t.Fatal(err)
				return
			}

			userRequest := request.UserRequest{
				Email:    email,
				Password: "test@ndo",
				Name:     "Test User",
				Age:      20,
			}

			b, _ := json.Marshal(userRequest)
			stringReader := io.NopCloser(strings.NewReader(string(b)))

			MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
			UserController.CreateUser(ctx)

			userEntity := entity.UserEntity{}

			filter := bson.D{{Key: "email", Value: email}}
			_ = Database.
				Collection(os.Getenv("MONGODB_USER_COLLECTION")).
				FindOne(context.Background(), filter).Decode(userEntity)

			assert.EqualValues(t, http.StatusOK, recorder.Code)
			assert.EqualValues(t, userRequest.Email, userEntity.Email)
			assert.EqualValues(t, userRequest.Age, userEntity.Age)
			assert.EqualValues(t, userRequest.Name, userEntity.Name)
		})

}
