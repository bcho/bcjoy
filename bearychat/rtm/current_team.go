package rtm

import "github.com/bearyinnovative/bcjoy/bearychat/model"

type CurrentTeamService struct {
	rtm *RTM
}

func newCurrentTeamService(rtm *RTM) error {
	rtm.CurrentTeam = &CurrentTeamService{rtm}

	return nil
}

func (c *CurrentTeamService) Info() (*model.Team, error) {
	client := setupClient(c.rtm.API, c.rtm.token)

	team := new(model.Team)
	_, err := client.Get("v1/current_team.info", team)

	return team, err
}

func (c *CurrentTeamService) Members() ([]*model.User, error) {
	client := setupClient(c.rtm.API, c.rtm.token)

	members := []*model.User{}
	_, err := client.Get("v1/current_team.members?all=true", &members)

	return members, err
}

func (c *CurrentTeamService) Channels() ([]*model.Channel, error) {
	client := setupClient(c.rtm.API, c.rtm.token)

	channels := []*model.Channel{}
	_, err := client.Get("v1/current_team.channels", &channels)

	return channels, err
}
