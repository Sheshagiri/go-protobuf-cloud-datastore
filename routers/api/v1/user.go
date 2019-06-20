package v1

import (
	"fmt"
	"github.com/Sheshagiri/go-protobuf-cloud-datastore/models"
	"github.com/gin-gonic/gin"
	"github.com/gogo/protobuf/proto"
	"io/ioutil"
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
