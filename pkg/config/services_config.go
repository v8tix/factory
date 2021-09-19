package config

import (
	"sync"
)

type SrvCfg struct {
	Wg *sync.WaitGroup
}
