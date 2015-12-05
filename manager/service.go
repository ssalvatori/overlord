package manager

import (
	"fmt"
	"sync"

	"github.com/ch3lo/overlord/configuration"
	"github.com/ch3lo/overlord/manager/types"
	"github.com/ch3lo/overlord/monitor"
)

type ServiceManager struct {
	serviceMux     sync.Mutex
	clusters       map[string]*types.Cluster
	serviceUpdater interface{}
}

func newServiceManager(config *configuration.Configuration) *ServiceManager {
	sm := &ServiceManager{}
	sm.setupClusters(config)

	return sm
}

// setupClusters inicia el cluster, mapeando el cluster el id del cluster como key
func (sm *ServiceManager) setupClusters(config *configuration.Configuration) {
	sm.clusters = make(map[string]*types.Cluster)

	for key, _ := range config.Clusters {
		sm.clusters[key] = types.NewCluster(key, config.Clusters[key])
	}
}

func (sm *ServiceManager) RegisterService(clusterKey string, service *types.ServiceParameters) (*types.ServiceContainer, error) {
	sm.serviceMux.Lock()
	defer sm.serviceMux.Unlock()

	for key, _ := range sm.clusters {
		if key == clusterKey {
			return sm.clusters[key].RegisterService(service)
		}
	}

	return nil, &ClusterDoesntExitsError{name: clusterKey}
}

// ClusterDoesntExitsError error generado cuando un cluster no existe
type ClusterDoesntExitsError struct {
	name string
}

func (err ClusterDoesntExitsError) Error() string {
	return fmt.Sprintf("El cluster no existe: %s", err.name)
}

func RegisterMonitor(service *types.ServiceParameters) {
	mon := &monitor.HttpMonitor{}
	//var mon monitor.Monitor
	mon.SetExpected(".*")
	mon.SetRetries(-1)
	mon.SetRequest("/")
	//service.monitor = mon
	//servicesBag[service.Id].GetMonitor() = mon
}
