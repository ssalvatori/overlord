package updater

import (
	"time"

	"github.com/ch3lo/overlord/scheduler"
	"github.com/ch3lo/overlord/service"
)

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

type serviceUpdate struct {
	scheduler  scheduler.Scheduler
	service    *service.ServiceVersion
	lastUpdate time.Time
}

func (u *serviceUpdate) sync() {
	u.scheduler.IsAlive("asd")
}

const changeManagerId string = "ChangeManager"

type ChangeManager struct {
	subscribers []Subscriber
}

func (cm *ChangeManager) id() string {
	return changeManagerId
}

func (cm *ChangeManager) Attach(s Subscriber, criteria ChangeCriteria) {
	// TODO manejar subscribers ya existentes
	cm.subscribers = append(cm.subscribers, s)
}

func (cm *ChangeManager) Detach(s Subscriber) {

}

func (cm *ChangeManager) update() {
	for k, _ := range cm.subscribers {
		cm.subscribers[k].update()
	}
}
