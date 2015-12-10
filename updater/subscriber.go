package updater

type ServiceUpdaterSubscriber interface {
	Id() string
	Update(map[string]ServiceUpdaterData)
}

type Subscriber interface {
	Id() string
	Update()
}
