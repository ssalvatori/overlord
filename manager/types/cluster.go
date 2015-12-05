package types

import (
	"github.com/ch3lo/overlord/configuration"
	"github.com/ch3lo/overlord/scheduler"
	"github.com/ch3lo/overlord/scheduler/factory"
	"github.com/ch3lo/overlord/util"
)

type Cluster struct {
	name             string
	scheduler        scheduler.Scheduler
	serviceContainer map[string]*ServiceContainer
}

func (c *Cluster) RegisterService(params *ServiceParameters) (*ServiceContainer, error) {
	container := newServiceContainer(params.Id)

	for key, _ := range c.serviceContainer {
		if key == params.Id {
			container = c.serviceContainer[key]
		}
	}

	c.serviceContainer[params.Id] = container

	return container, container.registerServiceVersion(params)
}

func NewCluster(name string, config configuration.Cluster) *Cluster {
	clusterScheduler, err := factory.Create(config.Scheduler.Type(), config.Scheduler.Parameters())
	if err != nil {
		util.Log.Fatalf("Error al crear el scheduler %s en %s. %s", config.Scheduler.Type(), name, err.Error())
	}

	c := &Cluster{
		name:             name,
		serviceContainer: make(map[string]*ServiceContainer),
		scheduler:        clusterScheduler,
	}

	return c
}
