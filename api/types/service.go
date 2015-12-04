package types

type Instance struct {
	Id      string `json:"id"`
	Address string `json:"address"`
}

type ServiceVersion struct {
	Version   string              `json:"version,omitempty"`
	Status    string              `json:"status"`
	Instances map[string]Instance `json:"instances,omitempty"`
}

type Service struct {
	Id       string                    `json:"id,omitempty"`
	Versions map[string]ServiceVersion `json:"versions,omitempty"`
}
