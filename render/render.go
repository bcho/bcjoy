package render

import "github.com/gin-gonic/contrib/renders/multitemplate"

func Make() multitemplate.Render {
	r := multitemplate.New()

	r.AddFromFiles("join", "template/join.html")

	return r
}
