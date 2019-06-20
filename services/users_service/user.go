package users_service

import "github.com/Sheshagiri/go-protobuf-cloud-datastore/models"

type U struct {
	User models.User
}

func (user *U) Add() error {
	//return models.AddTag(t.Name, t.State, t.CreatedBy)
	return models.AddUser(&user.User)
}
