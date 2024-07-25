package tests

import (
	"context"
	"encoding/json"
	"go-with-docker-and-swagger/src/controller/model/request"
	"go-with-docker-and-swagger/src/model/entity"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateUser(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx := GetTestGinContext(recorder)
	ID := primitive.NewObjectID()

	_, err := Database.
		Collection(os.Getenv("MONGODB_USER_COLLECTION")).
		InsertOne(context.Background(), bson.M{"_id": ID,
			"name":  "Old Name",
			"age":   10,
			"email": "flavio@test.com"})
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

	userRequest := request.UserUpdateRequest{
		Name: "Flavio New Name",
		Age:  30,
	}

	b, _ := json.Marshal(userRequest)
	stringReader := io.NopCloser(strings.NewReader(string(b)))

	MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)
	UserController.UpdateUser(ctx)

	assert.EqualValues(t, http.StatusOK, recorder.Result().StatusCode)

	userEntity := entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: ID}}
	_ = Database.
		Collection(os.Getenv("MONGODB_USER_COLLECTION")).
		FindOne(context.Background(), filter).Decode(&userEntity)

	assert.EqualValues(t, userRequest.Name, userEntity.Name)
	assert.EqualValues(t, userRequest.Age, userEntity.Age)

}
