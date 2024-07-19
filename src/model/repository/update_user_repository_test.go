package repository

import (
	"go-with-docker-and-swagger/src/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDb.Run("when_sending_a_valid_user_return_success",
		func(mt *mtest.T) {

			mt.AddMockResponses(bson.D{
				{Key: "ok", Value: 1},
				{Key: "n", Value: 1},
				{Key: "acknowledged", Value: true},
			})

			userDomain := model.NewUserDomain(
				"test@test.com", "test", "test", 90)
			userDomain.SetID(primitive.NewObjectID().Hex())
			databaseMock := mt.Client.Database(databaseName)

			repo := NewUserRepository(databaseMock)
			err := repo.UpdateUser(userDomain.GetID(), userDomain)

			assert.Nil(t, err)

		})

	mtestDb.Run("return_error_from_database",
		func(mt *mtest.T) {
			mt.AddMockResponses(bson.D{
				{Key: "ok", Value: 0},
			})

			userDomain := model.NewUserDomain(
				"test@test.com", "test", "test", 90)
			userDomain.SetID(primitive.NewObjectID().Hex())

			databaseMock := mt.Client.Database(databaseName)

			repo := NewUserRepository(databaseMock)
			err := repo.UpdateUser(userDomain.GetID(), userDomain)

			assert.NotNil(t, err)
		})
}
