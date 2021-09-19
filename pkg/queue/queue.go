package queue

import (
	. "github.com/v8tix/factory/pkg/config"
	"github.com/v8tix/factory/pkg/models/vehicle"
	"github.com/v8tix/factory/pkg/service/assemblyspot"
	. "github.com/v8tix/factory/pkg/utils"
	"log"
	"sync"
)

type Queue struct {
	srvCfg  *SrvCfg
	mu      sync.Mutex
	data    []*vehicle.Car
	latched *vehicle.Car
}

func NewQueue(capacity int, srvCfg *SrvCfg) *Queue {
	cars := make([]*vehicle.Car, 0, capacity)
	return &Queue{
		data:   cars,
		srvCfg: srvCfg,
		mu:     sync.Mutex{},
	}
}

func (q *Queue) Enqueue(vehicle *vehicle.Car) error {
	q.mu.Lock()
	q.data = append(q.data, vehicle)
	q.mu.Unlock()
	return nil
}

func (q *Queue) PendingData() bool {
	q.mu.Lock()
	pending := len(q.data) != 0
	q.mu.Unlock()
	return pending
}

func (q *Queue) DiscardData() error {
	q.mu.Lock()
	q.data = q.data[:0]
	q.mu.Unlock()
	return nil
}

func (q *Queue) Next() bool {
	q.mu.Lock()
	qLen := len(q.data)
	if qLen == 0 {
		q.mu.Unlock()
		return false
	}
	q.latched = q.data[qLen-1]
	q.data = q.data[:qLen-1]
	q.mu.Unlock()
	return true
}

func (q *Queue) Message() *vehicle.Car {
	q.mu.Lock()
	msg := q.latched
	q.mu.Unlock()
	return msg
}

func (q *Queue) BulkAdd(chunkSize int, value [][]*vehicle.Car) {
	YXAxisChunking(chunkSize, value, q.fillXAxis, q.srvCfg)
	q.srvCfg.Wg.Wait()
}

func YXAxisChunking(
	chunkSize int,
	container [][]*vehicle.Car,
	fn func(chunkSize int, a [][]*vehicle.Car, srvCfg *SrvCfg),
	srvCfg *SrvCfg,
) {
	for i := 0; i <= chunkSize; i += chunkSize {
		start := (i / chunkSize) * chunkSize
		end := (i/chunkSize + 1) * chunkSize
		a := container[start:end]
		SleepRandomTime()
		fn(chunkSize, a, srvCfg)
	}
}

func (q *Queue) fillXAxis(chunkSize int, a [][]*vehicle.Car, srvCfg *SrvCfg) {
	background2D(chunkSize, a, srvCfg, q.chunkX)
}

func background2D(
	chunkSize int,
	a [][]*vehicle.Car,
	srvCfg *SrvCfg,
	fun func(chunkSize int, a [][]*vehicle.Car, srvCfg *SrvCfg),
) {
	srvCfg.Wg.Add(1)
	go func(ac [][]*vehicle.Car) {
		defer srvCfg.Wg.Done()
		//fmt.Println(GetGoRoutineId())
		fun(chunkSize, a, srvCfg)
	}(a)
}

func (q *Queue) chunkX(chunkSize int, a [][]*vehicle.Car, srvCfg *SrvCfg) {
	for index := range a {
		XAxisChunking(chunkSize, a[index], srvCfg, q.fillYAxis)
	}
}

func XAxisChunking(
	chunkSize int,
	array []*vehicle.Car,
	srvCfg *SrvCfg,
	fn func(a []*vehicle.Car, srvCfg *SrvCfg),

) {
	for i := 0; i <= chunkSize; i += chunkSize {
		start := (i / chunkSize) * chunkSize
		end := (i/chunkSize + 1) * chunkSize
		a := array[start:end]
		SleepRandomTime()
		fn(a, srvCfg)
	}
}

func (q *Queue) fillYAxis(a []*vehicle.Car, srvCfg *SrvCfg) {
	background1D(a, srvCfg, q.chunkY)
}

func background1D(
	a []*vehicle.Car,
	srvCfg *SrvCfg,
	fun func(a []*vehicle.Car),
) {
	srvCfg.Wg.Add(1)
	go func(ac []*vehicle.Car) {
		defer srvCfg.Wg.Done()
		//fmt.Println(GetGoRoutineId())
		fun(a)
	}(a)
}

func (q *Queue) chunkY(a []*vehicle.Car) {
	for _, car := range a {
		SleepRandomTime()
		as := assemblyspot.AssemblySpot{}
		as.SetVehicle(car)
		if assembleVehicle, err := as.AssembleVehicle(); err != nil {

			log.Println(err)

		} else {

			err := q.Enqueue(assembleVehicle)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
