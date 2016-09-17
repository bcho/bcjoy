package route

import (
	"log"
	"strings"

	"github.com/bcho/bcjoy/badge"
	"github.com/bcho/bcjoy/team"
	"github.com/gin-gonic/gin"
)

var badgeColors = map[string]string{
	"green": "#85c158",
	"red":   "#d6604a",
}
var defaultBadgeColor = "#85c158"

func showDefaultBadge(c *gin.Context) {
	team := team.FromContext(c)

	totalMembers, err := team.TotalMembersCount()
	if err != nil {
		panic(err)
	}
	onlineMembers, err := team.OnlineMembersCount()
	if err != nil {
		panic(err)
	}

	showBadge(c, badge.New(totalMembers, onlineMembers))
}

func showStyledBadge(c *gin.Context) {
	team := team.FromContext(c)

	totalMembers, err := team.TotalMembersCount()
	if err != nil {
		panic(err)
	}
	onlineMembers, err := team.OnlineMembersCount()
	if err != nil {
		panic(err)
	}

	title, color := parseStyle(c.Params.ByName("style"))
	b := badge.New(
		totalMembers,
		onlineMembers,
		badge.BadgeTitle(title),
		badge.BadgeColor(color),
	)
	showBadge(c, b)
}

func showBadge(c *gin.Context, badge *badge.Badge) {
	c.Header("Content-Type", "image/svg+xml; charset=utf-8")
	c.Header("Cache-Control", "max-age=0, no-cache")
	c.Header("Pragma", "no-cache")

	c.String(200, "%s", badge)
}

func parseStyle(rawstyle string) (title, color string) {
	log.Printf("%s", rawstyle)
	parts := strings.Split(rawstyle, "-")
	log.Printf("%+v", parts)
	if len(parts) != 2 {
		title = ""
		color = defaultBadgeColor
		return
	}

	title = parts[0]
	if _, present := badgeColors[parts[1]]; !present {
		color = defaultBadgeColor
	} else {
		color = badgeColors[parts[1]]
	}

	return title, color
}
