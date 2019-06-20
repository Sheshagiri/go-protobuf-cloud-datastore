package helpers

import "github.com/gin-gonic/gin"

type Gin struct {
	 C *gin.Context
}

type Response struct {
	Code int `json: "code"`
	Message string `json:"string"`
	Data interface{} `json:"data"`
}

func (g *Gin) Response(httpCode int, errMessage string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code:httpCode,
		Message:errMessage,
		Data: data,
	})
	return
}