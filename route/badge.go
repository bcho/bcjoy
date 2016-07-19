package route

import (
	"github.com/bearyinnovative/bcjoy/team"
	"github.com/gin-gonic/gin"
)

func badge(c *gin.Context) {
	team := team.FromContext(c)

	totalMembers, err := team.TotalMembersCount()
	if err != nil {
		failed(err, w, r)
		return
	}
	onlineMembers, err := team.OnlineMembersCount()
	if err != nil {
		failed(err, w, r)
		return
	}

	c.Header("Content-Type", "image/svg+xml; charset=utf-8")
	c.Header("Cache-Control", "max-age=0, no-cache")
	c.Header("Pragma", "no-cache")

	badge := NewBadge(totalMembers, onlineMembers)
	c.String(200, badge)
}
