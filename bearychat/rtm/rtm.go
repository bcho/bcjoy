// RTM API
package rtm

import "github.com/bearyinnovative/bcjoy/bearychat/model"

const (
	DEFAULT_RTM_API = "https://rtm.bearychat.com"
)

type optSetter func(*RTM) error

type RTM struct {
	// rtm token
	token string

	// rtm api base, defaults to `https://rtm.bearychat.com`
	API string

	CurrentTeam *CurrentTeamService
	User        *UserService
	Channel     *ChannelService
}

var services = []optSetter{
	newCurrentTeamService,
	newUserService,
	newChannelService,
}

func New(token string, setters ...optSetter) (*RTM, error) {
	rtm := &RTM{
		token: token,
		API:   DEFAULT_RTM_API,
	}

	for _, setter := range setters {
		if err := setter(rtm); err != nil {
			return nil, err
		}
	}

	for _, serviceSetter := range services {
		if err := serviceSetter(rtm); err != nil {
			return nil, err
		}
	}

	return rtm, nil
}

func (r *RTM) Start() (*model.User, string, error) {
	client := setupClient(r.API, r.token)

	userAndWSHost := new(struct {
		User   *model.User `json:"user"`
		WSHost string      `json:"ws_host"`
	})
	_, err := client.Post("start", nil, userAndWSHost)

	return userAndWSHost.User, userAndWSHost.WSHost, err
}

func API(api string) optSetter {
	return func(rtm *RTM) error {
		rtm.API = api
		return nil
	}
}
