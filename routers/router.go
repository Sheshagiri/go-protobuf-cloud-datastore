package routers

import (
	apiv1 "github.com/Sheshagiri/go-protobuf-cloud-datastore/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	v1 := r.Group(("/api/v1"))

	v1.GET("/users", apiv1.GetUsers)
	v1.GET("/users/:username", apiv1.GetUser)
	v1.POST("/users", apiv1.AddUser)
	v1.PUT("/users/:username", apiv1.AddUser)
	v1.DELETE("/users/:username", apiv1.DeleteUser)

	return r
}
