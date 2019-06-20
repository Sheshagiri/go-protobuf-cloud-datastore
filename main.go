package main

import (
	"github.com/Sheshagiri/go-protobuf-cloud-datastore/models"
	"github.com/Sheshagiri/go-protobuf-cloud-datastore/routers"
	"log"
	"net/http"
)

func init() {
	models.Setup()
}

func main() {
	routers := routers.InitRouters()
	address := ":8080"
	server := &http.Server{
		Addr:    address,
		Handler: routers,
	}

	log.Printf("starting http server listening on %s", address)
	server.ListenAndServe()
}
