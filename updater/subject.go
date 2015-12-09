package updater

import (
	"sync"

	"github.com/ch3lo/overlord/util"
)

type Subject struct {
	subscriberMux sync.Mutex
	subscribers   map[string]Subscriber
}

func (s *Subject) Subscribe(sub Subscriber) {
	s.subscriberMux.Lock()
	defer s.subscriberMux.Unlock()

	for subscriberId, _ := range s.subscribers {
		if subscriberId == sub.Id() {
			return
		}
	}

	s.subscribers[sub.Id()] = sub
	util.Log.Infof("Se agregó el subscriptor: %s", sub.Id())
}

func (s *Subject) Unsubscribe(sub Subscriber) {
	s.subscriberMux.Lock()
	defer s.subscriberMux.Unlock()

	for k, v := range s.subscribers {
		if v.Id() == sub.Id() {
			delete(s.subscribers, k)
			util.Log.Infof("Se removio el subscriptor: %s", sub.Id())
			return
		}
	}
}

func (s *Subject) Notify() {
	s.subscriberMux.Lock()
	defer s.subscriberMux.Unlock()
	util.Log.Debugln("Notificando cambios")
	for k, _ := range s.subscribers {
		s.subscribers[k].Update()
	}
}
