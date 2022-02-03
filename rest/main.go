package rest

import (
	"net/http"

	"github.com/amaraad/goapp/graph/model"
	"github.com/gin-gonic/gin"
)

func GetRequest(c *gin.Context) {
	question := model.Question{}
	c.IndentedJSON(http.StatusOK, question)
}

func PostRequest(c *gin.Context) {
	question := model.Question{}
	c.IndentedJSON(http.StatusCreated, question)
}
