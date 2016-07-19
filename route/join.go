package route

import (
	"github.com/bearyinnovative/bcjoy/team"
	"github.com/gin-gonic/gin"
)

func join(c *gin.Context) {
	t := team.FromContext(c)
	team, err := t.Team()
	if err != nil {
		panic(err)
	}
	joinURL, err := t.JoinURL()
	if err != nil {
		panic(err)
	}

	c.HTML(200, "join", gin.H{
		"team":    team,
		"joinURL": joinURL,
	})
}