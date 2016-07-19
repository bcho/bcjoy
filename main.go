package main

import (
	"log"
	"os"

	"github.com/bearyinnovative/bcjoy/route"
	"github.com/bearyinnovative/bcjoy/team"
	"github.com/gin-gonic/gin"
)

func main() {
	t := makeTeam()

	r := route.Make(
		func(c *gin.Context) {
			team.ToContext(c, t)
			c.Next()
		},
	)

	r.Run(os.Getenv("BCJOY_BIND"))
}

func makeTeam() *team.Team {
	t, err := team.New(os.Getenv("BCJOY_RTM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	if err := t.UpdateMeta(); err != nil {
		log.Fatal(err)
	}

	if err := t.StartRTM(); err != nil {
		log.Fatal(err)
	}

	return t
}
