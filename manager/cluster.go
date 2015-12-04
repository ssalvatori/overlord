package manager

import (
	"github.com/ch3lo/overlord/configuration"
	"github.com/ch3lo/overlord/scheduler"
	"github.com/ch3lo/overlord/scheduler/factory"
	"github.com/ch3lo/overlord/util"
)

type Cluster struct {
	Name      string
	Scheduler scheduler.Scheduler
}

func NewCluster(name string, config configuration.Cluster) *Cluster {
	cluster := &Cluster{
		Name: name,
	}

	var err error
	cluster.Scheduler, err = factory.Create(config.Scheduler.Type(), config.Scheduler.Parameters())
	if err != nil {
		util.Log.Fatalf("Error al crear el scheduler %s en %s. %s", config.Scheduler.Type(), name, err.Error())
	}

	return cluster
}
