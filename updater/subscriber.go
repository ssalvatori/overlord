package updater

type Subscriber interface {
	Id() string
	Update()
}
