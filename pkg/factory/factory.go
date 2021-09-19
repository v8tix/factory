package factory

import (
	"fmt"
	. "github.com/v8tix/factory/pkg/config"
	"github.com/v8tix/factory/pkg/container"
	"github.com/v8tix/factory/pkg/models/vehicle"
	. "github.com/v8tix/factory/pkg/queue"
	"log"
	"time"
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
func (f *Factory) StartAssemblingProcess() {
	done := make(chan interface{})
	defer close(done)
	vehiclesSlice := f.CarContainer.Container
	for i, cars := range vehiclesSlice {
		for j, car := range cars {
			log.Println(fmt.Sprintf("(%d, %d) = %#v", i, j, car))
		}
	}
}

func (f *Factory) fanOutAssembleVehicles(
	done <-chan interface{},
	carsStream <-chan *vehicle.Car,
	assemblySpots int,
) []<-chan *vehicle.Car {
	assemblers := make([]<-chan *vehicle.Car, assemblySpots)
	for i := 0; i < assemblySpots; i++ {
		assemblers[i] = f.assembleVehicles(done, carsStream)
	}
	return assemblers
}

func (f *Factory) assembleVehicles(
	done <-chan interface{},
	carsStream <-chan *vehicle.Car,
) <-chan *vehicle.Car {
	outCarsStream := make(chan *vehicle.Car)
	go func() {
		defer close(outCarsStream)
		for car := range carsStream {
			time.Sleep(1 * time.Millisecond)
			select {
			case <-done:
				return
			case outCarsStream <- car:
			}
		}
	}()
	return outCarsStream
}

func generateVehicleStream(vehicles []*vehicle.Car) <-chan *vehicle.Car {
	vehicleStream := make(chan *vehicle.Car, len(vehicles))
	go func(vehicleStream chan *vehicle.Car) {
		defer close(vehicleStream)
		for i := 0; i < len(vehicles); i++ {
			vehicleStream <- vehicles[i]
		}
	}(vehicleStream)
	return vehicleStream
}

func (f *Factory) testCar(car *vehicle.Car) string {
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
