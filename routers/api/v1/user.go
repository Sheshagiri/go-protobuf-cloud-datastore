package v1

import (
	"github.com/Sheshagiri/go-protobuf-cloud-datastore/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"log"
)

func AddUser(c *gin.Context) {
	/*data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("unable to read all the data from http request",err)
	}
	fmt.Printf("data received is: %v", data)*/

	user := models.User{}

	if err := jsonpb.Unmarshal(c.Request.Body, &user); err != nil {
		log.Println("unable to unmarshall to user", err)
	}
	log.Println("user is: ", user)
	proto.DiscardUnknown(&user)
	models.AddUser(&user)
}

func GetUsers(c *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, users)
}
