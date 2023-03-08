package handler

import "github.com/gin-gonic/gin"

func Init() *gin.Engine {
	r := gin.New()
	r.LoadHTMLGlob("pages/*")
	r.Static("/static", "static/")
	r.NoRoute(notFoundHandler)

	r.GET("/", indexHandler)
	r.GET("/posts", postsHandler)
	r.GET("/posts/:title", postHandler)
	r.GET("/about", aboutHandler)
	r.GET("/support", supportHandler)
	r.GET("/agreement", agreementHandler)

	return r
}
