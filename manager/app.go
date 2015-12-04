package manager

import (
	"github.com/ch3lo/overlord/configuration"
	//Necesarios para que funcione el init()
	_ "github.com/ch3lo/overlord/scheduler/marathon"
	_ "github.com/ch3lo/overlord/scheduler/swarm"
)

var overlord *App = nil

type App struct {
	Config  *configuration.Configuration
	Cluster map[string]*Cluster
}

func NewApp(config *configuration.Configuration) {
	app := &App{
		Config: config,
	}

	app.setupClusters(config)

	overlord = app
}

func (app *App) setupClusters(config *configuration.Configuration) {
	app.Cluster = make(map[string]*Cluster)
	for key, _ := range config.Clusters {
		app.Cluster[key] = NewCluster(key, config.Clusters[key])
	}
}
