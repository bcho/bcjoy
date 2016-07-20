package rtm

import (
	"fmt"

	"github.com/bcho/bcjoy/bearychat/model"
)

type UserService struct {
	rtm *RTM
}

func newUserService(rtm *RTM) error {
	rtm.User = &UserService{rtm}

	return nil
}

func (c *UserService) Info(userId string) (*model.User, error) {
	client := setupClient(c.rtm.API, c.rtm.token)

	user := new(model.User)
	resource := fmt.Sprintf("v1/user.info?user_id=%s", userId)
	_, err := client.Get(resource, user)

	return user, err
}
