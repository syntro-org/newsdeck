package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/syntro-org/newsdeck/internal/article"
)

func postsHandler(c *gin.Context) {
	qParams := c.Request.URL.Query()
	var articles []article.Article

	if qParams != nil {
		q := qParams.Get("q")

		var sorted []article.Article

		for _, v := range article.Articles {
			if strings.Contains(v.Title, q) || strings.Contains(v.Description, q) || strings.Contains(v.Content, q) {
				sorted = append(sorted, v)
			}
		}

		articles = sorted
	} else {
		articles = article.Articles
	}

	c.HTML(http.StatusOK, "posts.html", gin.H{
		"Header":   "публикации",
		"Articles": articles,
	})
}
