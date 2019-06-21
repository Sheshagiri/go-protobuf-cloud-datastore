package models

import (
	"cloud.google.com/go/datastore"
	"github.com/Sheshagiri/go-protobuf-cloud-datastore/helpers"
	"log"
)

// AddUser adds a use to the cloud data store
func AddUser(user *User) (err error) {
	u := datastore.NameKey(helpers.USERS, user.Username, nil)
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
