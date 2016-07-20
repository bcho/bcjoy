package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bcho/bcjoy/route"
	"github.com/bcho/bcjoy/team"
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

	http.ListenAndServe(os.Getenv("BCJOY_BIND"), r)
}

func makeTeam() *team.Team {
	t, err := team.New(
		os.Getenv("BCJOY_RTM_TOKEN"),
		os.Getenv("BCJOY_JOIN_URL"),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := t.UpdateMeta(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second) // ratelimit
	if err := t.StartRTM(); err != nil {
		log.Fatal(err)
	}

	return t
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
