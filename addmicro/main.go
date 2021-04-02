package main

import (
	"github.com/gin-gonic/gin"
	"github.com/parikshitg/gqlarch/gqlgateway/graph/model"
)

func Plus(a, b int) int {
	return a + b
}

func Add(c *gin.Context) {
	var request model.SumRequest
	var response = &model.SumResponse{}

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(400, err)
		return
	}

	sum := Plus(request.A, request.B)
	response.Result = sum
	c.JSON(200, response)
}

func main() {
	r := gin.Default()
	r.POST("/add", Add)
	r.Run(":8081")
}
