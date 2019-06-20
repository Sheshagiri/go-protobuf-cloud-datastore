package routers

import (
	"github.com/gin-gonic/gin"
	apiv1 "github.com/Sheshagiri/go-protobuf-cloud-datastore/routers/api/v1"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group(("/api/v1"))

	//v1.GET("/users", v1.GetUsers)
	v1.POST("/users", apiv1.AddUser)
	//v1.PUT("/users/:id", v1.EditUser)
	//v1.DELETE("/users/:id", v1.DeleteUser)

	return r
}