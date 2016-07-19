package team

import (
	"sync"

	"github.com/bearyinnovative/bcjoy/bearychat/model"
	"github.com/bearyinnovative/bcjoy/bearychat/rtm"
)

const DEFAULT_MESSAGE_BACKLOG = 1024

type Team struct {
	rtm      *rtm.RTM
	rtmToken string
	rtmUser  *model.User

	lock               sync.RWMutex
	totalMembersCount  int
	onlineMembersCount int
	team               *model.Team
}

func New(rtmToken string) (*Team, error) {
	rtm, err := rtm.New(rtmToken)
	if err != nil {
		return nil, err
	}

	return &Team{
		rtm:      rtm,
		rtmToken: rtmToken,

		lock:               sync.RWMutex{},
		totalMembersCount:  0,
		onlineMembersCount: 0,
	}, nil
}

func (t *Team) StartRTM() error {
	user, _, err := t.rtm.Start()
	if err != nil {
		return err
	}

	t.rtmUser = user

	rtmPoller := Poller{
		Team: t,
	}
	if err := rtmPoller.Poll(); err != nil {
		return err
	}
	if err := rtmPoller.Start(); err != nil {
		return err
	}

	return nil
}

func (t *Team) Name() (string, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	if t.team != nil {
		return t.team.Name, nil
	}
	return "unknown", nil
}

func (t *Team) UpdateMeta() error {
	team, err := t.rtm.CurrentTeam.Info()
	if err != nil {
		return err
	}

	t.lock.Lock()
	t.team = team
	t.lock.Unlock()

	return nil
}

func (t *Team) UpdateMembers() error {
	members, err := t.rtm.CurrentTeam.Members()
	if err != nil {
		return err
	}

	onlineMembersCount := 0
	for _, member := range members {
		if member.IsOnline() {
			onlineMembersCount = onlineMembersCount + 1
		}
	}

	t.lock.Lock()
	t.totalMembersCount = len(members)
	t.onlineMembersCount = onlineMembersCount
	defer t.lock.Unlock()

	return nil
}

func (t *Team) TotalMembersCount() (int, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.totalMembersCount, nil
}

func (t *Team) OnlineMembersCount() (int, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()
	return t.onlineMembersCount, nil
}
