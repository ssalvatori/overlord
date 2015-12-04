package factory

import (
	"fmt"

	"github.com/ch3lo/overlord/scheduler"
	"github.com/ch3lo/overlord/util"
)

// schedulerFactories almacena una mapeo entre un identificador de scheduler y su constructor
var schedulerFactories = make(map[string]SchedulerFactory)

// SchedulerFactory es una interfaz para crear un Scheduler
// Cada Scheduler debe implementar estar interfaz y además llamar el metodo Register
type SchedulerFactory interface {
	Create(parameters map[string]interface{}) (scheduler.Scheduler, error)
}

// Register permite a una implementación de Scheduler estar disponible mediante un nombre
func Register(name string, factory SchedulerFactory) {
	if factory == nil {
		util.Log.Fatal("Se debe pasar como argumento un SchedulerFactory")
	}
	_, registered := schedulerFactories[name]
	if registered {
		util.Log.Fatalf("SchedulerFactory %s ya está registrado", name)
	}

	schedulerFactories[name] = factory
}

// Create crea un Scheduler a partir de su nombre.
// Si el Scheduler no estaba registrado se retornará un InvalidScheduler
func Create(name string, parameters map[string]interface{}) (scheduler.Scheduler, error) {
	schedulerFactory, ok := schedulerFactories[name]
	if !ok {
		return nil, InvalidScheduler{name}
	}
	return schedulerFactory.Create(parameters)
}

// InvalidScheduler sucede cuando se instenta construir un Scheduler no registrado
type InvalidScheduler struct {
	Name string
}

func (err InvalidScheduler) Error() string {
	return fmt.Sprintf("Scheduler no esta registrado: %s", err.Name)
}
