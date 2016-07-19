package team

import (
	"golang.org/x/net/context"
)

const KEY = "team"

type Setter interface {
	Set(string, interface{})
}

func FromContext(c context.Context) *Team {
	return c.Value(KEY).(*Team)
}

func ToContext(c Setter, team *Team) {
	c.Set(KEY, team)
}
