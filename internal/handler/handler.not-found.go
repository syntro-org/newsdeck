package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func notFoundHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "not-found.html", gin.H{
		"Header": "страница не найдена",
	})
}
