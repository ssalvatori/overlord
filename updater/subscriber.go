package updater

type Subscriber interface {
	id() string
	update()
}
