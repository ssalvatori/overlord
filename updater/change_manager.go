package updater

import "github.com/ch3lo/overlord/scheduler"

type ChangeCriteria int

const (
	ALWAYS = 1 + iota
)

var chageCriteria = [...]string{
	"ALWAYS",
}

func (cc ChangeCriteria) String() string {
	return chageCriteria[cc-1]
}

const changeManagerId string = "ChangeManager"

type ChangeManager struct {
	Subject
	services map[string]*scheduler.ServiceInformation
}

func NewChangeManager() *ChangeManager {
	cm := &ChangeManager{
		Subject: Subject{
			subscribers: make(map[string]Subscriber),
		},
		services: make(map[string]*scheduler.ServiceInformation),
	}
	return cm
}

func (cm *ChangeManager) Id() string {
	return changeManagerId
}

func (cm *ChangeManager) Update() {
	cm.Notify()
}

func (s *Subject) SubscribeWithCriteria(sub Subscriber, cc ChangeCriteria) {
	s.Subscribe(sub)
}
