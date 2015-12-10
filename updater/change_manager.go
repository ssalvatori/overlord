package updater

import (
	"sync"

	"github.com/ch3lo/overlord/scheduler"
	"github.com/ch3lo/overlord/util"
)

const changeManagerId string = "ChangeManager"

type ServiceChangeManagerData struct {
	serviceInformation map[string]*scheduler.ServiceInformation
}

type ServiceChangeManager struct {
	subscriberMux      sync.Mutex
	subscribers        map[string]Subscriber
	subscriberCriteria map[string]ChangeCriteria
	status             *ServiceChangeManagerData // DECORATOR?
	services           map[string]*scheduler.ServiceInformation
}

func NewChangeManager() *ServiceChangeManager {
	cm := &ServiceChangeManager{
		subscribers:        make(map[string]Subscriber),
		subscriberCriteria: make(map[string]ChangeCriteria),
		services:           make(map[string]*scheduler.ServiceInformation),
	}
	return cm
}

func (cm *ServiceChangeManager) Id() string {
	return changeManagerId
}

func (cm *ServiceChangeManager) Update(updatedServices map[string]ServiceUpdaterData) {
	cm.Notify()
}

func (cm *ServiceChangeManager) Notify() {
	cm.subscriberMux.Lock()
	defer cm.subscriberMux.Unlock()

	util.Log.Debugln("Notificando cambios")
	for k, _ := range cm.subscribers {
		//cm.subscriberCriteria[k].MeetCriteria()
		cm.subscribers[k].Update()
	}
}

func (cm *ServiceChangeManager) Remove(sub Subscriber) {
	cm.subscriberMux.Lock()
	defer cm.subscriberMux.Unlock()

	for k, v := range cm.subscribers {
		if v.Id() == sub.Id() {
			delete(cm.subscriberCriteria, k)
			delete(cm.subscribers, k)
			util.Log.Infof("Se removio el subscriptor: %s", sub.Id())
			return
		}
	}
}

func (cm *ServiceChangeManager) SubscribeWithCriteria(sub Subscriber, cc ChangeCriteria) {
	cm.subscriberMux.Lock()
	defer cm.subscriberMux.Unlock()

	for subscriberId, _ := range cm.subscribers {
		if subscriberId == sub.Id() {
			return
		}
	}

	cm.subscriberCriteria[sub.Id()] = cc
	cm.subscribers[sub.Id()] = sub
	util.Log.Infof("Se agregó el subscriptor: %s", sub.Id())
}
