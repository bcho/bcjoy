package route

import (
	"github.com/bcho/bcjoy/badge"
	"github.com/bcho/bcjoy/team"
	"github.com/gin-gonic/gin"
)

func showBadge(c *gin.Context) {
	team := team.FromContext(c)

	totalMembers, err := team.TotalMembersCount()
	if err != nil {
		panic(err)
	}
	onlineMembers, err := team.OnlineMembersCount()
	if err != nil {
		panic(err)
	}

	c.Header("Content-Type", "image/svg+xml; charset=utf-8")
	c.Header("Cache-Control", "max-age=0, no-cache")
	c.Header("Pragma", "no-cache")

	b := badge.New(totalMembers, onlineMembers)
	c.String(200, "%s", b)
}
