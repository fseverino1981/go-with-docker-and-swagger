package service

import (
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser model", zap.String("jorney", "createUser"))
	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Error try to call repository", err, zap.String("jorney", "createUser"))
		return nil, err
	}

	logger.Info("CreateUser service executed successfully", zap.String("userID", userDomainRepository.GetID()),
		zap.String("journey", "createUser"))

	return userDomainRepository, nil

}
