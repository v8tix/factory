package assemblyspot

import (
	"errors"
	"fmt"
	. "github.com/v8tix/factory/pkg/config"
	. "github.com/v8tix/factory/pkg/models/vehicle"
	"time"
)

type AssemblySpot struct {
	vehicleToAssemble *Car
	assemblyLog       string
	SrvCfg            *SrvCfg
}

func (s *AssemblySpot) SetVehicle(v *Car) {
	s.vehicleToAssemble = v
}

func (s *AssemblySpot) GetAssembledVehicle() *Car {
	return s.vehicleToAssemble
}

func (s *AssemblySpot) GetAssembledLogs() string {
	return s.assemblyLog
}

func (s *AssemblySpot) AssembleVehicle() (*Car, error) {
	if s.vehicleToAssemble == nil {
		return nil, errors.New("no set to start assembling")
	}

	s.assembleChassis()
	s.assembleTires()
	s.assembleEngine()
	s.assembleElectronics()
	s.assembleDash()
	s.assembleSeats()
	s.assembleWindows()

	return s.vehicleToAssemble, nil
}

func (s *AssemblySpot) assembleChassis() {
	s.vehicleToAssemble.Chassis = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Chassis at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) assembleTires() {
	s.vehicleToAssemble.Tires = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Tires at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) assembleEngine() {
	s.vehicleToAssemble.Engine = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Engine at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) assembleElectronics() {
	s.vehicleToAssemble.Electronics = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Electronics at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) assembleDash() {
	s.vehicleToAssemble.Dash = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Dash at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) assembleSeats() {
	s.vehicleToAssemble.Sits = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Sits at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) assembleWindows() {
	s.vehicleToAssemble.Windows = "Assembled"
	time.Sleep(1 * time.Second)
	s.assemblyLog += fmt.Sprintf("Windows at [%s], ", time.Now().Format("2006-01-02 15:04:05.000"))
}

func (s *AssemblySpot) TestCar(car *Car) string {
	logs := ""

	result, err := car.StartEngine()
	if err == nil {
		logs += result + ", "
	} else {
		logs += err.Error() + ", "
	}

	result, err = car.MoveForwards(10)
	if err == nil {
		logs += result + ", "
	} else {
		logs += err.Error() + ", "
	}

	result, err = car.MoveForwards(10)
	if err == nil {
		logs += result + ", "
	} else {
		logs += err.Error() + ", "
	}

	result, err = car.TurnLeft()
	if err == nil {
		logs += result + ", "
	} else {
		logs += err.Error() + ", "
	}

	result, err = car.TurnRight()
	if err == nil {
		logs += result + ", "
	} else {
		logs += err.Error() + ", "
	}

	result, err = car.StopEngine()
	if err == nil {
		logs += result + ", "
	} else {
		logs += err.Error() + ", "
	}

	return logs
}
