package service

import (
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/configuration/rest_err"

	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserServices(userId string) *rest_err.RestErr {
	logger.Info("Init deleteUser services", zap.String("jorney", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error try to call repository", err, zap.String("jorney", "deleteUser"))
		return err
	}

	logger.Info("DeleteUser service executed successfully", zap.String("userId", userId),
		zap.String("journey", "deleteUser"))

	return nil
}
