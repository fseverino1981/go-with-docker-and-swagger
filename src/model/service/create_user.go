package service

import (
	"fmt"
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("jorney", "createUser"))
	userDomain.EncryptPassword()

	fmt.Print(userDomain.GetPassword())

	return nil

}
