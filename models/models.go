package models

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"os"
)

var db *datastore.Client
var ctx context.Context

func Setup() {
	var err error
	ctx = context.Background()
	projectID := os.Getenv("DATASTORE_PROJECT_ID")
	if projectID == "" {
		log.Fatal(`set the environment variable "DATASTORE_PROJECT_ID"`)
	}

	if db, err = datastore.NewClient(ctx, projectID); err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

}
