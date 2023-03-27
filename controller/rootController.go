package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(ctx *gin.Context) {
	name, isName := ctx.GetQuery("name")

	var message string

	if isName {
		message = fmt.Sprintf("Hello %s welcome to Books API", name)
	} else {
		message = "Hello welcome to Books API"
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": message,
	})
}
