## The goal here is to see if we can directly persist the protobuf's upon unmarshalling to Google Cloud Datastore.

## How to run

1. Set the you google project id to DATASTORE_PROJECT_ID environment variable

```bash
$ go run main.go
```

## How to test

```bash
$ curl --request POST \
    --url http://localhost:8080/api/v1/users \
    --header 'content-type: application/json' \
    --data '{
  	"username":"Steve",
  	"first_name":"Steve",
  	"last_name":"Jobs"
  }'
```

## How to verify

You should see an entry in google cloud firestore with `users` as the key.


