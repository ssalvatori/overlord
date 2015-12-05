package types

import "time"

type ServiceParameters struct {
	Id           string
	Version      string
	ImageName    string
	ImageTag     string
	MinInstances int
}

type Instance struct {
	Id      string
	Address string
	Status  string
}

type ServiceVersion struct {
	Version      string
	CreateDate   time.Time
	ImageName    string
	ImageTag     string
	Instances    []*Instance
	MinInstances int
}

type ServiceContainer struct {
	Id         string
	CreateDate time.Time
	Container  map[string]*ServiceVersion
}

func (s *ServiceContainer) registerServiceVersion(params *ServiceParameters) error {
	for key, _ := range s.Container {
		if key == params.Version {
			return &ServiceVersionAlreadyExist{service: s.Id, version: params.Version}
		}
	}

	sv := &ServiceVersion{
		Version:      params.Version,
		CreateDate:   time.Now(),
		ImageName:    params.ImageName,
		ImageTag:     params.ImageTag,
		Instances:    make([]*Instance, 0),
		MinInstances: params.MinInstances,
	}

	s.Container[params.Version] = sv
	return nil
}

func newServiceContainer(id string) *ServiceContainer {
	container := &ServiceContainer{
		Id:         id,
		CreateDate: time.Now(),
		Container:  make(map[string]*ServiceVersion),
	}

	return container
}
