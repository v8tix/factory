package container

import (
	"errors"
	"fmt"
	. "github.com/v8tix/factory/pkg/config"
	"github.com/v8tix/factory/pkg/models/vehicle"
	. "github.com/v8tix/factory/pkg/utils"
	"math/rand"
	"time"
)

const factor = 10

type CarContainer struct {
	Container [][]*vehicle.Car
	SrvCfg    *SrvCfg
}

func NewCarContainer(amountOfVehicles, chunkSize int, srvCfg *SrvCfg) (*CarContainer, error) {
	if !(amountOfVehicles/factor == factor) {
		return nil, errors.New("the remainder of the division between the number of cars and 10 should be 10")
	}
	a := createEmptySquareArray(factor)
	now := time.Now()
	YXAxisChunking(amountOfVehicles, chunkSize, a, fillXAxis, srvCfg)
	srvCfg.Wg.Wait()
	fmt.Printf("container creation {elapsed time: %v}\n", time.Since(now))
	return &CarContainer{
		Container: a,
		SrvCfg:    srvCfg,
	}, nil
}

func YXAxisChunking(
	amountOfVehicles,
	chunkSize int,
	emptyArray [][]*vehicle.Car,
	fn func(amountOfVehicles, chunkSize int, a [][]*vehicle.Car, srvCfg *SrvCfg),
	srvCfg *SrvCfg,
) {
	for i := 0; i <= chunkSize; i += chunkSize {
		start := (i / chunkSize) * chunkSize
		end := (i/chunkSize + 1) * chunkSize
		a := emptyArray[start:end]
		SleepRandomTime()
		fn(amountOfVehicles, chunkSize, a, srvCfg)
	}
}

func fillXAxis(amountOfVehicles, chunkSize int, a [][]*vehicle.Car, srvCfg *SrvCfg) {
	background2D(amountOfVehicles, chunkSize, a, srvCfg, chunkX)
}

func background2D(
	amountOfVehicles,
	chunkSize int,
	a [][]*vehicle.Car,
	srvCfg *SrvCfg,
	fun func(amountOfVehicles, chunkSize int, a [][]*vehicle.Car, srvCfg *SrvCfg),
) {
	srvCfg.Wg.Add(1)
	go func(ac [][]*vehicle.Car) {
		defer srvCfg.Wg.Done()
		fun(amountOfVehicles, chunkSize, a, srvCfg)
	}(a)
}

func chunkX(amountOfVehicles int, chunkSize int, a [][]*vehicle.Car, srvCfg *SrvCfg) {
	for index := range a {
		XAxisChunking(amountOfVehicles, chunkSize, a[index], srvCfg, fillYAxis)
	}
}

func XAxisChunking(
	amountOfVehicles int,
	chunkSize int,
	array []*vehicle.Car,
	srvCfg *SrvCfg,
	fn func(amountOfVehicles int, a []*vehicle.Car, srvCfg *SrvCfg),

) {
	for i := 0; i <= chunkSize; i += chunkSize {
		start := (i / chunkSize) * chunkSize
		end := (i/chunkSize + 1) * chunkSize
		a := array[start:end]
		SleepRandomTime()
		fn(amountOfVehicles, a, srvCfg)
	}
}

func fillYAxis(amountOfVehicles int, a []*vehicle.Car, srvCfg *SrvCfg) {
	background1D(amountOfVehicles, a, srvCfg, chunkY)
}

func background1D(
	amountOfVehicles int,
	a []*vehicle.Car,
	srvCfg *SrvCfg,
	fun func(amountOfVehicles int, a []*vehicle.Car),
) {
	srvCfg.Wg.Add(1)
	go func(ac []*vehicle.Car) {
		defer srvCfg.Wg.Done()
		fun(amountOfVehicles, a)
	}(a)
}

func chunkY(amountOfVehicles int, a []*vehicle.Car) {
	for index := range a {
		SleepRandomTime()
		a[index] = vehicle.NewCar(rand.Intn(amountOfVehicles) + 1)
	}
}

func createEmptySquareArray(n int) [][]*vehicle.Car {
	a := make([][]*vehicle.Car, n)
	for i := 0; i < len(a); i++ {
		a[i] = make([]*vehicle.Car, n)
	}
	return a
}
