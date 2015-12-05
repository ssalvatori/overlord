package types

import "fmt"

// ServiceAlreadyExist sucede cuando se esta agregando un servicio que ya existe
type ServiceAlreadyExist struct {
	Name string
}

func (err ServiceAlreadyExist) Error() string {
	return fmt.Sprintf("El servicio ya existe: %s", err.Name)
}

// ServiceVersionAlreadyExist sucede cuando se esta agregando una version de servicio que ya existe
type ServiceVersionAlreadyExist struct {
	service string
	version string
}

func (err ServiceVersionAlreadyExist) Error() string {
	return fmt.Sprintf("La version % del servicio %s ya existe", err.version, err.service)
}
