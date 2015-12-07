package updater

import (
	"reflect"
	"sync"
	"time"

	"github.com/ch3lo/overlord/scheduler"
	"github.com/ch3lo/overlord/service"
	"github.com/ch3lo/overlord/util"
)

type ServiceUpdater struct {
	subscriberMux     sync.Mutex
	clusterSchedulers []scheduler.Scheduler
	subscribers       map[string]Subscriber
	services          map[string]*service.ServiceInstance
}

func NewServiceUpdater(schedulers []scheduler.Scheduler) *ServiceUpdater {
	if schedulers == nil {
		util.Log.Fatalln("Al menos se debe monitorear un scheduler")
	}

	s := &ServiceUpdater{
		subscribers: make(map[string]Subscriber),
		services:    make(map[string]*service.ServiceInstance),
	}
	s.clusterSchedulers = schedulers

	return s
}

func (su *ServiceUpdater) Attach(s Subscriber) {
	su.subscriberMux.Lock()
	defer su.subscriberMux.Unlock()

	for subscriberId, _ := range su.subscribers {
		if subscriberId == s.id() {
			return
		}
	}

	su.subscribers[s.id()] = s
	util.Log.Infof("Se agregó el subscriptor: %s", s.id())
}

func (su *ServiceUpdater) Detach(s Subscriber) {
	su.subscriberMux.Lock()
	defer su.subscriberMux.Unlock()

	for k, v := range su.subscribers {
		if v.id() == s.id() {
			delete(su.subscribers, k)
			util.Log.Infof("Se removio el subscriptor: %s", s.id())
			return
		}
	}
}

// Monitor comienza el monitoreo de los servicios de manera desatachada
func (su *ServiceUpdater) Monitor() {
	go su.detachedMonitor()
}

// detachedMonitor loop que permite monitorear los servicios de los schedulers
func (su *ServiceUpdater) detachedMonitor() {
	for {
		mustNotify := false

		for _, sched := range su.clusterSchedulers {
			srvs, _ := sched.GetInstances(scheduler.FilterInstances{})
			su.attachServices(sched, srvs)
		}

		if mustNotify {
			su.notify()
		}

		time.Sleep(time.Second * 10)
	}
}

func (su *ServiceUpdater) attachServices(sched scheduler.Scheduler, services []*service.ServiceInstance) bool {
	updated := false

	for k, v := range services {
		util.Log.Debugf("Comparando servicio %+v <-> %+v", su.services[v.Id], services[k])
		if su.services[v.Id] != nil {
			if reflect.DeepEqual(su.services[v.Id], services[k]) {
				util.Log.Debugln("Servicio sin cambios")
				continue
			}
			util.Log.Debugln("Servicio tuvo un cambio")
		} else {
			util.Log.Debugln("Monitoreando nuevo servicio")
		}

		su.services[v.Id] = services[k]
		updated = true
	}

	return updated
}

func (su *ServiceUpdater) notify() {
	su.subscriberMux.Lock()
	defer su.subscriberMux.Unlock()

	for k, _ := range su.subscribers {
		su.subscribers[k].update()
	}
}
