package route

import (
	"github.com/bcho/bcjoy/team"
	"github.com/gin-gonic/gin"
)

func join(c *gin.Context) {
	t := team.FromContext(c)
	team, err := t.Team()
	if err != nil {
		panic(err)
	}
	onlineMembers, err := t.OnlineMembersCount()
	if err != nil {
		panic(err)
	}
	totalMembers, err := t.TotalMembersCount()
	if err != nil {
		panic(err)
	}

	c.HTML(200, "join", gin.H{
		"team":          team,
		"onlineMembers": onlineMembers,
		"totalMembers":  totalMembers,
	})
}

func joinApply(c *gin.Context) {
	t := team.FromContext(c)

	joinUrl, err := t.JoinURL()
	if err != nil {
		panic(err)
	}
	c.Redirect(301, joinUrl)
}
