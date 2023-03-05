package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func agreementHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "agreement.html", gin.H{
		"Header": "пользовательское соглашение",
	})
}
