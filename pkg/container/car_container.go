package container

import (
	"errors"
	. "github.com/v8tix/factory/pkg/config"
	. "github.com/v8tix/factory/pkg/models/vehicle"
	. "github.com/v8tix/factory/pkg/utils"
	"math/rand"
)

const factor = 10

type CarContainer struct {
	Container [][]*Car
	SrvCfg    *SrvCfg
}

func NewCarContainer(amountOfVehicles, chunkSize int, srvCfg *SrvCfg) (*CarContainer, error) {
	if !(amountOfVehicles/factor == factor) {
		return nil, errors.New("the remainder of the division between the number of cars and 10 should be 10")
	}
	a := createEmptySquareArray(factor)
	YXAxisChunking(amountOfVehicles, chunkSize, a, fillXAxis, srvCfg)
	srvCfg.Wg.Wait()
	return &CarContainer{
		Container: a,
		SrvCfg:    srvCfg,
	}, nil
}

func YXAxisChunking(
	amountOfVehicles,
	chunkSize int,
	emptyArray [][]*Car,
	fn func(amountOfVehicles, chunkSize int, a [][]*Car, srvCfg *SrvCfg),
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

func fillXAxis(amountOfVehicles, chunkSize int, a [][]*Car, srvCfg *SrvCfg) {
	background2D(amountOfVehicles, chunkSize, a, srvCfg, chunkX)
}

func background2D(
	amountOfVehicles,
	chunkSize int,
	a [][]*Car,
	srvCfg *SrvCfg,
	fun func(amountOfVehicles, chunkSize int, a [][]*Car, srvCfg *SrvCfg),
) {
	srvCfg.Wg.Add(1)
	go func(ac [][]*Car) {
		defer srvCfg.Wg.Done()
		fun(amountOfVehicles, chunkSize, a, srvCfg)
	}(a)
}

func chunkX(amountOfVehicles int, chunkSize int, a [][]*Car, srvCfg *SrvCfg) {
	for index := range a {
		XAxisChunking(amountOfVehicles, chunkSize, a[index], srvCfg, fillYAxis)
	}
}

func XAxisChunking(
	amountOfVehicles int,
	chunkSize int,
	array []*Car,
	srvCfg *SrvCfg,
	fn func(amountOfVehicles int, a []*Car, srvCfg *SrvCfg),

) {
	for i := 0; i <= chunkSize; i += chunkSize {
		start := (i / chunkSize) * chunkSize
		end := (i/chunkSize + 1) * chunkSize
		a := array[start:end]
		SleepRandomTime()
		fn(amountOfVehicles, a, srvCfg)
	}
}

func fillYAxis(amountOfVehicles int, a []*Car, srvCfg *SrvCfg) {
	background1D(amountOfVehicles, a, srvCfg, chunkY)
}

func background1D(
	amountOfVehicles int,
	a []*Car,
	srvCfg *SrvCfg,
	fun func(amountOfVehicles int, a []*Car),
) {
	srvCfg.Wg.Add(1)
	go func(ac []*Car) {
		defer srvCfg.Wg.Done()
		fun(amountOfVehicles, a)
	}(a)
}

func chunkY(amountOfVehicles int, a []*Car) {
	for index := range a {
		SleepRandomTime()
		a[index] = NewCar(rand.Intn(amountOfVehicles) + 1)
	}
}

func createEmptySquareArray(n int) [][]*Car {
	a := make([][]*Car, n)
	for i := 0; i < len(a); i++ {
		a[i] = make([]*Car, n)
	}
	return a
}
