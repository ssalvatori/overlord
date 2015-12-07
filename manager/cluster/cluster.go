package cluster

import (
	"github.com/ch3lo/overlord/configuration"
	"github.com/ch3lo/overlord/scheduler"
	"github.com/ch3lo/overlord/scheduler/factory"
	"github.com/ch3lo/overlord/service"
	"github.com/ch3lo/overlord/util"
)

type Cluster struct {
	name             string
	scheduler        scheduler.Scheduler
	serviceContainer map[string]*service.ServiceContainer
}

// RegisterService registra un nuevo servicio en un contenedor de servicios
// Si el contenedor ya existia se omite su creación y se procede a registrar
// las versiones de los servicios.
// Si no se puede registrar una nueva version se retornara un error.
func (c *Cluster) RegisterService(params *service.ServiceParameters) (*service.ServiceContainer, error) {
	container := service.NewServiceContainer(params.Id)

	for key, _ := range c.serviceContainer {
		if key == params.Id {
			container = c.serviceContainer[key]
		}
	}

	c.serviceContainer[params.Id] = container

	return container, container.RegisterServiceVersion(params)
}

// NewCluster crea un nuevo cluster a partir de un id y parametros de configuracion
// La configuración es necesaria para configurar el scheduler
func NewCluster(name string, config configuration.Cluster) (*Cluster, error) {
	if config.Disabled {
		return nil, &ClusterDisabled{Name: name}
	}

	clusterScheduler, err := factory.Create(config.Scheduler.Type(), config.Scheduler.Parameters())
	if err != nil {
		util.Log.Fatalf("Error al crear el scheduler %s en %s. %s", config.Scheduler.Type(), name, err.Error())
	}

	c := &Cluster{
		name:             name,
		serviceContainer: make(map[string]*service.ServiceContainer),
		scheduler:        clusterScheduler,
	}

	return c, nil
}

// GetScheduler retorna el scheduler que utiliza el cluster
func (c *Cluster) GetScheduler() scheduler.Scheduler {
	return c.scheduler
}
