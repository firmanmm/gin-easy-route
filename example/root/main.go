package main

import (
	"net/http"

	router "github.com/firmanmm/gin-easy-route"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	//We inject gin.Engine instance to the router
	routeBuilder := router.NewRouter(engine)

	//Lets make *simple* authorization check
	routeBuilder.AddMiddleware("/", func(ctx *gin.Context) {
		if len(ctx.GetHeader("Authorization")) == 0 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Missing Authorization Header",
			})
		}
	})

	//And another url
	routeBuilder.AddRoute(http.MethodGet, "/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Authenticated Main",
		})
	})

	//And another url
	routeBuilder.AddRoute(http.MethodGet, "/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Authenticated Hello",
		})
	})

	//And another url
	routeBuilder.AddRoute(http.MethodGet, "/hello/:msg", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Authenticated Hello " + ctx.Param("msg"),
		})
	})

	//Or try using post?
	routeBuilder.AddRoute(http.MethodPost, "/hello", func(ctx *gin.Context) {
		ctx.Request.ParseForm()
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello",
			"form":    ctx.Request.PostForm,
		})
	})

	//And trigger all the build
	routeBuilder.Build()
	engine.Run()
}
