package queue

import (
	"github.com/v8tix/factory/pkg/models/vehicle"
	"github.com/v8tix/factory/pkg/service/assemblyspot"
	"sync"
	"time"
)

type Queue struct {
	capacity  int
	condition *sync.Cond
	data      []*vehicle.Car
}

func NewQueue(cars []*vehicle.Car, capacity int) *Queue {
	cond := sync.NewCond(&sync.Mutex{})
	return &Queue{
		condition: cond,
		data:      cars,
		capacity:  capacity,
	}
}

func (q *Queue) Add(value *vehicle.Car) bool {
	q.condition.L.Lock()
	for q.Size() == q.capacity {
		q.condition.Wait()
	}
	as := &assemblyspot.AssemblySpot{}
	as.SetVehicle(value)
	car, _ := as.AssembleVehicle()
	q.data = append(q.data, car)
	q.condition.L.Unlock()
	q.condition.Signal()
	return true
}

func (q *Queue) Remove(delay time.Duration) (*vehicle.Car, bool) {
	time.Sleep(delay)
	q.condition.L.Lock()

	var value *vehicle.Car
	for q.Size() == 0 {
		q.condition.Wait()
	}

	q.data = q.data[1:]
	q.condition.L.Unlock()
	q.condition.Signal()
	return value, true
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Size() int {
	return len(q.data)
}
