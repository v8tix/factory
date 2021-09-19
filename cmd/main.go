package main

import (
	. "github.com/v8tix/factory/pkg/config"
	"github.com/v8tix/factory/pkg/factory"
	"log"
	"os"
)

const (
	carsAmount = 100
)

func main() {
	var srvCfg SrvCfg
	infoLog, errorLog := setupLoggers()
	srvCfg.Log.InfoLog = infoLog
	srvCfg.Log.ErrorLog = errorLog

	fcty := factory.New(&srvCfg)

	//Hint: change appropriately for making factory give each vehicle once assembled, even though the others have not been assembled yet,
	//each vehicle delivered to main should display testinglogs and assemblelogs with the respective vehicle id
	//receiver := make(chan vehicle.Car)
	fcty.StartAssemblingProcess(carsAmount)
	srvCfg.Wg.Wait()
}

func setupLoggers() (*log.Logger, *log.Logger) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return infoLog, errorLog
}
