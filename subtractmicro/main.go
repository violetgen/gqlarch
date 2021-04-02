package main

import (
	"github.com/gin-gonic/gin"
	"github.com/parikshitg/gqlarch/gqlgateway/graph/model"
)

func Minus(a, b int) int {
	return a - b
}

func Subtract(c *gin.Context) {
	var request model.SubtractRequest
	var response = &model.SubtractResponse{}

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, err)
		return
	}

	sub := Minus(request.A, request.B)
	response.Result = sub
	c.JSON(200, response)
}

func main() {
	r := gin.Default()
	r.POST("/subtract", Subtract)
	r.Run(":8082")
}
