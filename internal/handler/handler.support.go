package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func supportHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "support.html", gin.H{
		"Header": "поддержать",
	})
}
