package v1

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"fmt"
	"github.com/Sheshagiri/go-protobuf-cloud-datastore/models"
	"github.com/gogo/protobuf/proto"
)

func AddUser(c *gin.Context) {
	data, err := ioutil.ReadAll(c.Request.Body)
	user := models.User{}

	if err != nil {
		fmt.Println(err)
	}
	if err := proto.Unmarshal(data, &user); err != nil {
		fmt.Println(err)
	}
	models.AddUser(&user)
}
