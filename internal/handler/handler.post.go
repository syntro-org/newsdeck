package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syntro-org/newsdeck/internal/article"
)

func postHandler(c *gin.Context) {
	title := c.Param("title")
	obj := article.GetArticleByTitle(title)

	c.HTML(http.StatusOK, "post.html", gin.H{
		"Header":  obj.Title,
		"Article": obj,
	})
}
