package team

import (
	"sync"

	"github.com/bcho/bcjoy/bearychat/model"
	"github.com/bcho/bcjoy/bearychat/rtm"
)

type Team struct {
	rtm      *rtm.RTM
	rtmToken string
	rtmUser  *model.User

	joinURL string

	lock               sync.RWMutex
	totalMembersCount  int
	onlineMembersCount int
	team               *model.Team
}

func New(rtmToken string, joinURL string) (*Team, error) {
	rtm, err := rtm.New(rtmToken)
	if err != nil {
		return nil, err
	}

	return &Team{
		rtm:      rtm,
		rtmToken: rtmToken,

		joinURL: joinURL,

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

func (t *Team) Team() (*model.Team, error) {
	t.lock.RLock()
	defer t.lock.RUnlock()

	return t.team, nil
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

	totalMembersCount := 0
	onlineMembersCount := 0
	for _, member := range members {
		if !member.IsNormal() {
			continue
		}

		totalMembersCount = totalMembersCount + 1
		if member.IsOnline() {
			onlineMembersCount = onlineMembersCount + 1
		}
	}

	t.lock.Lock()
	t.totalMembersCount = totalMembersCount
	t.onlineMembersCount = onlineMembersCount
	t.lock.Unlock()

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

func (t *Team) JoinURL() (string, error) {
	return t.joinURL, nil
}
