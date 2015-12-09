package updater

import (
	"reflect"
	"time"

	"github.com/ch3lo/overlord/manager/cluster"
	"github.com/ch3lo/overlord/scheduler"
	"github.com/ch3lo/overlord/util"
)

type ServiceUpdater struct {
	Subject
	clusters map[string]*cluster.Cluster
	services map[string]*scheduler.ServiceInformation
}

func NewServiceUpdater(clusters map[string]*cluster.Cluster) *ServiceUpdater {
	if clusters == nil || len(clusters) == 0 {
		util.Log.Fatalln("Al menos se debe monitorear un cluster")
	}

	s := &ServiceUpdater{
		Subject: Subject{
			subscribers: make(map[string]Subscriber),
		},
		services: make(map[string]*scheduler.ServiceInformation),
	}
	s.clusters = clusters

	return s
}

// Monitor comienza el monitoreo de los servicios de manera desatachada
func (su *ServiceUpdater) Monitor() {
	go su.detachedMonitor()
}

// detachedMonitor loop que permite monitorear los servicios de los schedulers
func (su *ServiceUpdater) detachedMonitor() {
	for {
		mustNotify := false

		for clusterKey, c := range su.clusters {
			srvs, err := c.GetScheduler().GetInstances(scheduler.FilterInstances{})
			if err != nil {
				util.Log.Errorln("No se pudieron obtener instancias del cluster %s con scheduler %. Motivo: %s", clusterKey, c.GetScheduler().Id(), err.Error())
				continue
			}
			mustNotify = su.attachServices(srvs)
		}

		if mustNotify {
			su.Notify()
		}

		time.Sleep(time.Second * 10)
	}
}

func (su *ServiceUpdater) attachServices(services []*scheduler.ServiceInformation) bool {
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
