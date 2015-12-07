package service

import (
	"time"
)

// ServiceParameters es una estructura que encapsula los parametros
// de configuración de un nuevo servicio
type ServiceParameters struct {
	Id           string
	Version      string
	ImageName    string
	ImageTag     string
	MinInstances int
}

// ServiceInstance contiene la información de una instancia de un servicio
type ServiceInstance struct {
	Id      string
	Address string
	Status  string
}

// ServiceVersion es una estructura que contiene la información de una
// version de un servicio.
type ServiceVersion struct {
	Version      string
	CreateDate   time.Time
	ImageName    string
	ImageTag     string
	Instances    []*ServiceInstance
	MinInstances int
}

// ServiceContainer agrupa un conjuntos de versiones de un servicio bajo el parametro Container
type ServiceContainer struct {
	Id         string
	CreateDate time.Time
	Container  map[string]*ServiceVersion
}

// RegisterServiceVersion registra una nueva version de servicio en el contenedor
// Si la version ya existia se retornara un error ServiceVersionAlreadyExist
func (s *ServiceContainer) RegisterServiceVersion(params *ServiceParameters) error {
	for key, _ := range s.Container {
		if key == params.Version {
			return &ServiceVersionAlreadyExist{Service: s.Id, Version: params.Version}
		}
	}

	sv := &ServiceVersion{
		Version:      params.Version,
		CreateDate:   time.Now(),
		ImageName:    params.ImageName,
		ImageTag:     params.ImageTag,
		Instances:    make([]*ServiceInstance, 0),
		MinInstances: params.MinInstances,
	}

	s.Container[params.Version] = sv
	return nil
}

// NewServiceContainer crea un nuevo contenedor de servicios con la fecha actual
func NewServiceContainer(id string) *ServiceContainer {
	container := &ServiceContainer{
		Id:         id,
		CreateDate: time.Now(),
		Container:  make(map[string]*ServiceVersion),
	}

	return container
}
