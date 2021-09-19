package factory

import (
	. "github.com/v8tix/factory/pkg/config"
	"github.com/v8tix/factory/pkg/container"
	. "github.com/v8tix/factory/pkg/models/vehicle"
	. "github.com/v8tix/factory/pkg/queue"
)

type Factory struct {
	CarContainer *container.CarContainer
	SrvCfg       *SrvCfg
	Queue        *Queue
}

func New(srvCfg *SrvCfg, carContainer *container.CarContainer, queue *Queue) *Factory {
	factory := &Factory{
		SrvCfg:       srvCfg,
		CarContainer: carContainer,
		Queue:        queue,
	}

	return factory
}

// StartAssemblingProcess HINT: this function is currently not returning anything, make it return right away every single vehicle once assembled,
//(Do not wait for all of them to be assembled to return them all, send each one ready over to main)
func (f *Factory) StartAssemblingProcess(chunkSize int) {
	done := make(chan interface{})
	defer close(done)
	vehiclesSlice := f.CarContainer.Container
	f.Queue.BulkAdd(chunkSize, vehiclesSlice)
}

func (f *Factory) testCar(car *Car) string {
	logs := ""

	log, err := car.StartEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.MoveForwards(10)
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnLeft()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.TurnRight()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	log, err = car.StopEngine()
	if err == nil {
		logs += log + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}
