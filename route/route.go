package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Make(middlewares ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())

	e.Use(middlewares...)

	e.GET("badge.svg", badge)

	return e
}
