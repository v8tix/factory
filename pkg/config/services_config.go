package config

import (
	"log"
	"sync"
)

type SrvCfg struct {
	Wg  *sync.WaitGroup
	Log struct {
		ErrorLog *log.Logger
		InfoLog  *log.Logger
	}
}
