package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syntro-org/newsdeck/internal/article"
)

func postHandler(c *gin.Context) {
	title := c.Param("title")
	var obj article.Article

	for _, v := range article.Articles {
		if v.Title == title {
			obj = v
			break
		}
	}

	c.HTML(http.StatusOK, "post.html", gin.H{
		"Header":  obj.Title,
		"Article": obj,
	})
}
