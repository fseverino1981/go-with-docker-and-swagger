package repository

import (
	"context"
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/model"
	"go-with-docker-and-swagger/src/model/entity/converter"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur userRepository) UpdateUser(userID string, userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init updateUser repository", zap.String("journey", "updateUser"))
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)
	userIDHex, _ := primitive.ObjectIDFromHex(userID)

	filter := bson.D{{Key: "_id", Value: userIDHex}}
	update := bson.D{{Key: "$set", Value: value}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		logger.Error("Error trying to update user", err, zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("UpdateUser repository executed successfully", zap.String("userId", userID),
		zap.String("journey", "updateUser"))

	return nil
}
