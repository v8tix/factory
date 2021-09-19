package utils

import (
	"fmt"
	. "github.com/v8tix/factory/pkg/config"
	"math/rand"
	"time"
)

func BackgroundTask(fn func(), srvCfg *SrvCfg) {

	srvCfg.Wg.Add(1)

	go func() {

		defer srvCfg.Wg.Done()

		defer func() {
			if err := recover(); err != nil {
				srvCfg.Log.ErrorLog.Print(fmt.Errorf("%s", err), nil)
			}
		}()

		fn()
	}()
}

func computeBinaryInt() int {
	t := time.Now().UnixNano()
	rand.Seed(t)
	target := rand.Intn(2) % 2
	return target
}

func SleepRandomTime() {
	time.Sleep(time.Duration(computeBinaryInt()+1) * time.Nanosecond)
}
