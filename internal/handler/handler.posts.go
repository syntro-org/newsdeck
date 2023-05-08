package handler

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/syntro-org/newsdeck/internal/article"
)

func postsHandler(c *gin.Context) {
	qParams := c.Request.URL.Query()
	var articles []article.Article

	if qParams != nil {
		q := qParams.Get("q")
		d, err := getDateFromQuery(qParams)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": err.Error(),
			})
			return
		}

		articles = searchArticles(q, d)
	} else {
		articles = article.Articles
	}

	c.HTML(http.StatusOK, "posts.html", gin.H{
		"Header":   "публикации",
		"Articles": articles,
	})
}

func searchArticles(q string, d time.Time) []article.Article {
	var results []article.Article

	for _, v := range article.Articles {
		if q != "" && !containsContent(q, v) {
			continue
		}
		if !d.IsZero() && !matchesDate(d, v) {
			continue
		}
		results = append(results, v)
	}

	return results
}

func containsContent(q string, v article.Article) bool {
	if strings.Contains(v.Title, q) || strings.Contains(v.Description, q) || strings.Contains(v.Content, q) {
		return true
	}
	return false
}

func matchesDate(d time.Time, v article.Article) bool {
	layout := "2006-01-02"
	date, err := time.Parse(layout, v.PublishedAt)
	if err != nil {
		return false
	}
	return date.Equal(d)
}

func getDateFromQuery(query url.Values) (time.Time, error) {
	if strings.Contains(query.Get("d"), "-") {
		rawDate := strings.Split(query.Get("d"), "-")
		year, err := strconv.Atoi(strings.ReplaceAll(rawDate[2], "\\", ""))

		if err != nil {
			return time.Time{}, err
		}

		monthInt, err := strconv.Atoi(rawDate[1])

		if err != nil {
			return time.Time{}, err
		}

		day, err := strconv.Atoi(rawDate[0])

		if err != nil {
			return time.Time{}, err
		}

		return time.Date(year, time.Month(monthInt), day, 0, 0, 0, 0, time.UTC), nil
	}

	return time.Time{}, nil
}
