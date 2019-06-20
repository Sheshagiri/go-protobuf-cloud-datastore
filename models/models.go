package models

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
)

var db *datastore.Client
var ctx context.Context

type Model struct {
	ID int
	createdOn int
	modifiedOn int
	deletedOn int
}

func Setup() {
	var err error
	ctx = context.Background()
	projectID := "abcd"

	if db, err = datastore.NewClient(ctx, projectID); err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

}
