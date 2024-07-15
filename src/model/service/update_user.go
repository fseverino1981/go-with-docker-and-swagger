package service

import (
	"go-with-docker-and-swagger/src/configuration/rest_err"
	"go-with-docker-and-swagger/src/model"
)

func (ud *userDomainService) UpdateUser(userID string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	return nil
}
