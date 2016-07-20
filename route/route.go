package route

import (
	"net/http"

	"github.com/bcho/bcjoy/render"
	"github.com/gin-gonic/gin"
)

func Make(middlewares ...gin.HandlerFunc) http.Handler {
	e := gin.New()
	e.HTMLRender = render.Make()

	e.Use(gin.Recovery())

	e.Use(middlewares...)

	e.GET("badge.svg", showBadge)
	e.GET("join", join)

	return e
}
