package repository

import (
	"context"
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/model"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	MONGODB_USER_DB = "MONGODB_USER_DB "
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

func (ur userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init userCreate repository", zap.String("journey", "createUser"))
	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))
	return nil, nil
}
