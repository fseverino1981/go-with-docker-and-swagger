package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteUser(t *testing.T) {
	recorder := httptest.NewRecorder()
	ctx := GetTestGinContext(recorder)
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

	MakeRequest(ctx, param, url.Values{}, "GET", nil)
	UserController.DeleteUser(ctx)

	UserController.FindUserById(ctx)

	assert.EqualValues(t, http.StatusNotFound, recorder.Code)

	filter := bson.D{{Key: "_id", Value: ID}}
	result := Database.
		Collection(os.Getenv("MONGODB_USER_COLLECTION")).
		FindOne(context.Background(), filter)

	assert.NotNil(t, result.Err())

}
