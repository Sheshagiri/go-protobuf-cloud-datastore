package models

import (
	"cloud.google.com/go/datastore"
	"log"
)

// AddUser adds a use to the cloud data store
func AddUser(user *User) (err error) {
	u := datastore.NameKey("users", user.Username, nil)
	if _, err = db.Put(ctx, u, &user); err != nil {
		log.Fatalf("failed to save the user: %v", err)
		return
	}
	log.Printf("saved %s", user.Username)
	return
}
