package main

import (
	. "github.com/v8tix/factory/pkg/config"
	"github.com/v8tix/factory/pkg/container"
	"github.com/v8tix/factory/pkg/factory"
	. "github.com/v8tix/factory/pkg/queue"
	"log"
	"sync"
	"time"
)

const (
	carsAmount = 100
	capacity   = 5
	chunkSize  = 5
)

func main() {

	var srvCfg SrvCfg
	var wg sync.WaitGroup
	srvCfg.Wg = &wg
	queue := NewQueue(capacity, &srvCfg)
	carContainer, _ := container.NewCarContainer(carsAmount, chunkSize, &srvCfg)
	fcty := factory.New(&srvCfg, carContainer, queue)

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	var wgr sync.WaitGroup
	listener(queue, &wgr)
	fcty.StartAssemblingProcess(chunkSize)
	wgr.Wait()
}

func listener(queue *Queue, wgr *sync.WaitGroup) {
	wgr.Add(1)
	go func() {
		defer wgr.Done()
		for {
			if queue.Next() {
				log.Println(queue.Message().AssembleLog)
				log.Println(queue.Message().TestingLog)
			}
			time.Sleep(1 * time.Second)
		}
	}()
}
