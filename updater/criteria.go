package updater

import "github.com/ch3lo/overlord/scheduler"

type ChangeCriteria interface {
	MeetCriteria(status map[string]*scheduler.ServiceInformation) map[string]*scheduler.ServiceInformation
}

type ImageNameCriteria struct {
	Name string
}

func (c *ImageNameCriteria) MeetCriteria(elements map[string]*scheduler.ServiceInformation) map[string]*scheduler.ServiceInformation {
	var filtered map[string]*scheduler.ServiceInformation
	for k, v := range elements {
		if c.Name == v.Image {
			filtered[k] = elements[k]
		}
	}
	return filtered
}

type StatusCriteria struct {
	Status string
}

func (c *StatusCriteria) MeetCriteria(elements map[string]*scheduler.ServiceInformation) map[string]*scheduler.ServiceInformation {
	var filtered map[string]*scheduler.ServiceInformation
	for k, v := range elements {
		if c.Status == v.Status {
			filtered[k] = elements[k]
		}
	}
	return filtered
}

type AndCriteria struct {
	criteria      ChangeCriteria
	otherCriteria ChangeCriteria
}

func (c *AndCriteria) MeetCriteria(elements map[string]*scheduler.ServiceInformation) map[string]*scheduler.ServiceInformation {
	filtered := c.criteria.MeetCriteria(elements)
	return c.otherCriteria.MeetCriteria(filtered)
}

type OrCriteria struct {
	criteria      ChangeCriteria
	otherCriteria ChangeCriteria
}

func (c *OrCriteria) MeetCriteria(elements map[string]*scheduler.ServiceInformation) map[string]*scheduler.ServiceInformation {
	filtered := c.criteria.MeetCriteria(elements)
	others := c.otherCriteria.MeetCriteria(elements)

	for k, _ := range others {
		if filtered[k] == nil {
			filtered[k] = others[k]
		}
	}

	return filtered
}
