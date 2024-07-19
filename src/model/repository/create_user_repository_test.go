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

func TestUserRepository_CreateUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	os.Setenv("MONGODB_USER_DB", collectionName)
	defer os.Clearenv()

	mtestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	//defer mtestDb.Close()

	mtestDb.Run("when_sending_a_valid_domain_returns_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"test@test.com", "test", "test", 90)
		userDomain, err := repo.CreateUser(domain)

		_, errID := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errID)
		assert.EqualValues(t, userDomain.GetEmail(), domain.GetEmail())
		assert.EqualValues(t, userDomain.GetName(), domain.GetName())
		assert.EqualValues(t, userDomain.GetAge(), domain.GetAge())
		assert.EqualValues(t, userDomain.GetPassword(), domain.GetPassword())
	})

	mtestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"test@test.com", "test", "test", 90)
		userDomain, err := repo.CreateUser(domain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
