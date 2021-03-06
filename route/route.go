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

	e.GET("badge.svg", showDefaultBadge)
	e.GET("badge/:style", showStyledBadge)
	e.GET("join", join)
	e.GET("join/apply", joinApply)

	return e
}
