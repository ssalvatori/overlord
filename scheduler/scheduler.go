package scheduler

type Scheduler interface {
	IsAlive(id string) (bool, error)
}
