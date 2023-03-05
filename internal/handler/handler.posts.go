package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syntro-org/newsdeck/internal/article"
)

func postsHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "posts.html", gin.H{
		"Header":   "публикации",
		"Articles": article.Articles,
	})
}
