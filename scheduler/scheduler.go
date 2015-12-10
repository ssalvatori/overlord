package scheduler

// Scheduler es una interfaz que debe implementar cualquier scheduler de Servicios
// Para un ejemplo ir a swarm.SwarmScheduler
type Scheduler interface {
	Id() string
	IsAlive(id string) (bool, error)
	GetInstances(filter FilterInstances) ([]ServiceInformation, error)
}

// FilterInstances es una estrutura para encapsular los requerimientos
// que se utilizaran para filtrar instancias de servicios
type FilterInstances struct {
	imageName string
	imageTag  string
}

// SchedulerUser es una interfaz que deben implementar las clases que usen un scheduler
type SchedulerUser interface {
	GetScheduler() Scheduler
}
