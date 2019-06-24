package v1

import (
	"cloud.google.com/go/datastore"
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

func GetUser(c *gin.Context) {
	username := c.Param("username")
	log.Printf("get %s details", username)
	user, err := models.GetUser(username)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			c.JSON(404, "no such user")
			return
		}
		c.JSON(500, err.Error())
	}
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	username := c.Param("username")
	log.Printf("deleting %s", username)
	if err := models.DeleteUser(username); err != nil {
		c.JSON(500, err.Error())
	}
	c.JSON(200, nil)
}

func UpdateUser(c *gin.Context) {
	user := models.User{}

	if err := jsonpb.Unmarshal(c.Request.Body, &user); err != nil {
		log.Println("unable to unmarshall to user", err)
	}
	log.Printf("updating details of %s: ", user.Username)
	proto.DiscardUnknown(&user)
	models.UpdateUser(&user)
}
