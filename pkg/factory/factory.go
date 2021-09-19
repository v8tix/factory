package factory

import (
	"fmt"
	. "github.com/v8tix/factory/pkg/config"
	"github.com/v8tix/factory/pkg/models/vehicle"
	"github.com/v8tix/factory/pkg/service/assemblyspot"
	"log"
	"math/rand"
	"time"
)

const (
	assemblySpots int = 5
)

type Factory struct {
	AssemblingSpots []*assemblyspot.AssemblySpot
	SrvCfg          *SrvCfg
}

func New(srvCfg *SrvCfg) *Factory {
	factory := &Factory{
		AssemblingSpots: make([]*assemblyspot.AssemblySpot, assemblySpots),
		SrvCfg:          srvCfg,
	}

	return factory
}

// StartAssemblingProcess HINT: this function is currently not returning anything, make it return right away every single vehicle once assembled,
//(Do not wait for all of them to be assembled to return them all, send each one ready over to main)
func (f *Factory) StartAssemblingProcess(amountOfVehicles int) {
	done := make(chan interface{})
	defer close(done)
	vehiclesSlice := createVehiclesSlice(amountOfVehicles, assemblySpots)
	for i, cars := range vehiclesSlice {
		log.Println(fmt.Sprintf("starting iteration %d", i))
		vehicleLotsStream := generateVehicleStream(cars)
		finders := fanOutAssembleVehicles(done, vehicleLotsStream, assemblySpots)
		for _, finder := range finders {
			car := <-finder
			log.Println(car)
		}
	}
}

func fanOutAssembleVehicles(
	done <-chan interface{},
	carsStream <-chan *vehicle.Car,
	assemblySpots int,
) []<-chan *vehicle.Car {
	assemblers := make([]<-chan *vehicle.Car, assemblySpots)
	for i := 0; i < assemblySpots; i++ {
		assemblers[i] = assembleVehicles(done, carsStream)
	}
	return assemblers
}

func assembleVehicles(
	done <-chan interface{},
	carsStream <-chan *vehicle.Car,
) <-chan *vehicle.Car {
	outCarsStream := make(chan *vehicle.Car)
	go func() {
		defer close(outCarsStream)
		for car := range carsStream {
			time.Sleep(1 * time.Millisecond)
			//TODO: write to queue
			select {
			case <-done:
				return
			case outCarsStream <- car:
			}
		}
	}()
	return outCarsStream
}

func createVehiclesSlice(amountOfVehicles, assemblySpots int) [][]*vehicle.Car {
	size := amountOfVehicles / assemblySpots
	a := make([][]*vehicle.Car, size)
	for i := 0; i < len(a); i++ {
		a[i] = make([]*vehicle.Car, assemblySpots)
		for j := 0; j < assemblySpots; j++ {
			car := vehicle.New(rand.Intn(amountOfVehicles) + 1)
			a[i][j] = car
		}
	}
	return a
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
