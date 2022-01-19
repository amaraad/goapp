package rest

import (
	"net/http"

	"github.com/amaraad/goapp/graph/model"
	"github.com/gin-gonic/gin"
)

func getRequest(c *gin.Context) {
	question := model.Question{}
	c.IndentedJSON(http.StatusOK, question)
}

func postRequest(c *gin.Context) {
	question := model.Question{}
	c.IndentedJSON(http.StatusCreated, question)
}
