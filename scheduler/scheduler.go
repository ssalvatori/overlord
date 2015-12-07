package scheduler

import "github.com/ch3lo/overlord/service"

// Scheduler es una interfaz que debe implementar cualquier scheduler de Servicios
// Para un ejemplo ir a swarm.SwarmScheduler
type Scheduler interface {
	IsAlive(id string) (bool, error)
	GetInstances(filter FilterInstances) ([]*service.ServiceInstance, error)
}

// FilterInstances es una estrutura para encapsular los requerimientos
// que se utilizaran para filtrar instancias de servicios
type FilterInstances struct {
	imageName string
	imageTag  string
}
