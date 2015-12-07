package manager

import (
	"sync"

	"github.com/ch3lo/overlord/configuration"
	"github.com/ch3lo/overlord/manager/cluster"
	"github.com/ch3lo/overlord/scheduler"
	"github.com/ch3lo/overlord/service"
	"github.com/ch3lo/overlord/updater"
	"github.com/ch3lo/overlord/util"
	//Necesarios para que funcione el init()
	_ "github.com/ch3lo/overlord/scheduler/marathon"
	_ "github.com/ch3lo/overlord/scheduler/swarm"
)

var overlordApp *Overlord

type Overlord struct {
	config         *configuration.Configuration
	serviceMux     sync.Mutex
	clusters       map[string]*cluster.Cluster
	serviceUpdater *updater.ServiceUpdater
}

func NewApp(config *configuration.Configuration) {
	app := &Overlord{
		config: config,
	}

	app.setupClusters(config)
	app.setupUpdater()

	overlordApp = app
}

func GetAppInstance() *Overlord {
	return overlordApp
}

// setupClusters inicia el cluster, mapeando el cluster el id del cluster como key
func (o *Overlord) setupClusters(config *configuration.Configuration) {
	o.clusters = make(map[string]*cluster.Cluster)

	for key, _ := range config.Clusters {
		c, err := cluster.NewCluster(key, config.Clusters[key])
		if err != nil {
			util.Log.Infof(err.Error())
			continue
		}

		o.clusters[key] = c
	}

	if len(o.clusters) == 0 {
		util.Log.Fatalln("Al menos debe existir un cluster")
	}
}

func (o *Overlord) setupUpdater() {
	var schedulers []scheduler.Scheduler
	for k, _ := range o.clusters {
		schedulers = append(schedulers, o.clusters[k].GetScheduler())
	}

	su := updater.NewServiceUpdater(schedulers)
	cm := &updater.ChangeManager{}
	su.Attach(cm)

	su.Monitor()

	o.serviceUpdater = su
}

func (o *Overlord) RegisterService(clusterKey string, srv *service.ServiceParameters) (*service.ServiceContainer, error) {
	o.serviceMux.Lock()
	defer o.serviceMux.Unlock()

	for key, _ := range o.clusters {
		if key == clusterKey {
			return o.clusters[key].RegisterService(srv)
		}
	}

	return nil, &cluster.ClusterDoesntExits{Name: clusterKey}
}
