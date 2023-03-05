package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func aboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{
		"Header": "о нас",
	})
}
