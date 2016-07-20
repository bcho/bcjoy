package rtm

import (
	"fmt"

	"github.com/bcho/bcjoy/bearychat/model"
)

type ChannelService struct {
	rtm *RTM
}

func newChannelService(rtm *RTM) error {
	rtm.Channel = &ChannelService{rtm}

	return nil
}

func (c *ChannelService) Info(channelId string) (*model.Channel, error) {
	client := setupClient(c.rtm.API, c.rtm.token)

	channel := new(model.Channel)
	resource := fmt.Sprintf("v1/channel.info?channel_id=%s", channelId)
	_, err := client.Get(resource, channel)

	return channel, err
}
