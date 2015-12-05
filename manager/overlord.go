package manager

import (
	"github.com/ch3lo/overlord/configuration"
	//Necesarios para que funcione el init()
	_ "github.com/ch3lo/overlord/scheduler/marathon"
	_ "github.com/ch3lo/overlord/scheduler/swarm"
)

var overlordApp *Overlord

type Overlord struct {
	config         *configuration.Configuration
	serviceManager *ServiceManager
}

func NewApp(config *configuration.Configuration) {
	app := &Overlord{
		config: config,
	}

	app.serviceManager = newServiceManager(config)

	overlordApp = app
}

func GetAppInstance() *Overlord {
	return overlordApp
}
