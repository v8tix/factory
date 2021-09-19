package main

import (
	"fmt"
	. "github.com/v8tix/factory/pkg/config"
	"github.com/v8tix/factory/pkg/container"
	"github.com/v8tix/factory/pkg/factory"
	"github.com/v8tix/factory/pkg/models/vehicle"
	. "github.com/v8tix/factory/pkg/queue"
	"log"
	"os"
	"sync"
	"time"
)

const (
	carsAmount = 100
	capacity   = 5
	chunkSize  = 5
)

func main() {
	cars := make([]*vehicle.Car, 0, 5)
	queue := NewQueue(cars, capacity)
	var srvCfg SrvCfg
	infoLog, errorLog := setupLoggers()
	srvCfg.Log.InfoLog = infoLog
	srvCfg.Log.ErrorLog = errorLog
	var wg sync.WaitGroup
	srvCfg.Wg = wg

	carContainer, _ := container.NewCarContainer(carsAmount, chunkSize, &srvCfg)
	fcty := factory.New(&srvCfg, carContainer, queue)

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	//reader(queue)
	fcty.StartAssemblingProcess()
	//srvCfg.Wg.Wait()
}

func reader(queue *Queue) {
	go func() {

		for {
			car, result := queue.Remove(1 * time.Second)
			if result {

				log.Println(fmt.Sprintf("Removed from queue %#v\n", car))

			} else {

				log.Println("couldn't remove car")
			}
		}
	}()
}

func setupLoggers() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}
