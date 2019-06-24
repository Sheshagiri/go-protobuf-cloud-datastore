package models

import (
	"cloud.google.com/go/datastore"
	"github.com/Sheshagiri/go-protobuf-cloud-datastore/helpers"
	"github.com/golang/protobuf/ptypes"
	"log"
)

// AddUser adds a use to the cloud data store
func AddUser(user *User) (err error) {
	u := datastore.NameKey(helpers.USERS, user.Username, nil)
	user.CreatedOn = ptypes.TimestampNow()
	if _, err = db.Put(ctx, u, user); err != nil {
		log.Fatalf("failed to save the user: %v", err)
		return
	}
	log.Printf("saved %s", user.Username)
	return
}

// GetUsers will get all the users present in the data store
func GetUsers() (users []User, err error) {
	query := datastore.NewQuery(helpers.USERS)
	if _, err := db.GetAll(ctx, query, &users); err != nil {
		return nil, err
	}
	return
}

// GetUser will get the user for the given username
func GetUser(username string) (user User, err error) {
	searchKey := datastore.NameKey(helpers.USERS, username, nil)
	if err = db.Get(ctx, searchKey, &user); err != nil {
		log.Printf("unable to query user: %s", username)
		return
	}
	return
}

// DeleteUser will delete the user if matched to the given username
func DeleteUser(username string) (err error) {
	deleteKey := datastore.NameKey(helpers.USERS, username, nil)
	err = db.Delete(ctx, deleteKey)
	return
}

// UpdateUser adds a use to the cloud data store
func UpdateUser(user *User) error {
	if ok, err := GetUser(user.Username); err != nil {
		return err
	} else {
		user.CreatedOn = ok.CreatedOn
	}
	u := datastore.NameKey(helpers.USERS, user.Username, nil)
	user.LastUpdatedOn = ptypes.TimestampNow()
	if _, err := db.Put(ctx, u, user); err != nil {
		log.Fatalf("failed to save the user: %v", err)
		return err
	}
	log.Printf("saved %s", user.Username)
	return nil
}
