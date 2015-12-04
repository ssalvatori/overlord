package manager

import (
	"github.com/ch3lo/overlord/monitor"
	"github.com/ch3lo/overlord/util"
)

type Instance struct {
	Id      string
	Address string
}

type ServiceVersion struct {
	Version   string
	Instances []*Instance
}

type Service struct {
	Id       string
	Versions []*ServiceVersion
}

var servicesBag []*Service = make([]*Service, 0)

func ServicesBag() []*Service {
	return servicesBag
}

func RegisterService(service *Service) error {
	for _, srv := range servicesBag {
		if srv.Id == service.Id {
			return &util.ElementAlreadyExists{}
		}
	}

	servicesBag = append(servicesBag, service)
	return nil
}

func RegisterServiceVersion(serviceId string, sv *ServiceVersion) error {
	for i, srv := range servicesBag {
		if srv.Id == serviceId {
			for _, ver := range srv.Versions {
				if ver.Version == sv.Version {
					return &util.ElementAlreadyExists{}
				}
			}
			servicesBag[i].Versions = append(servicesBag[i].Versions, sv)
			return nil
		}
	}
	return &util.ServiceNotFound{}
}

func GetService(cluster string, id string) (bool, error) {
	return overlord.Cluster[cluster].Scheduler.IsAlive(id)
}

func RegisterMonitor(service *Service) {
	mon := &monitor.HttpMonitor{}
	//var mon monitor.Monitor
	mon.SetExpected(".*")
	mon.SetRetries(-1)
	mon.SetRequest("/")
	//service.monitor = mon
	//servicesBag[service.Id].GetMonitor() = mon
}

func (s *Service) GetMonitor() monitor.Monitor {
	return nil //s.monitor
}
