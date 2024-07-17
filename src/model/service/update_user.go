package service

import (
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices(userID string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init updateUser services", zap.String("jorney", "updateUser"))

	err := ud.userRepository.UpdateUser(userID, userDomain)
	if err != nil {
		logger.Error("Error try to call repository", err, zap.String("jorney", "updateUser"))
		return err
	}

	logger.Info("UpdateUser service executed successfully", zap.String("userID", userID),
		zap.String("journey", "updateUser"))

	return nil

}
