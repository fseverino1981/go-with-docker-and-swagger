package model

import (
	"fmt"
	"go-with-docker-and-swagger/src/configuration/logger"
	"go-with-docker-and-swagger/src/configuration/rest_err"

	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("jorney", "createUser"))
	ud.EncryptPassword()

	fmt.Print(ud)

	return nil

}
