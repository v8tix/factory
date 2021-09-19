package vehicle

import "fmt"

type Car struct {
	Id            int
	Chassis       string
	Tires         string
	Engine        string
	Electronics   string
	Dash          string
	Sits          string
	Windows       string
	EngineStarted bool
	TestingLog    string
	AssembleLog   string
}

func NewCar(id int) *Car {
	return &Car{
		Id:            id,
		Chassis:       "NotSet",
		Tires:         "NotSet",
		Engine:        "NotSet",
		Electronics:   "NotSet",
		Dash:          "NotSet",
		Sits:          "NotSet",
		Windows:       "NotSet",
		EngineStarted: false,
	}
}

func (c *Car) StartEngine() (string, error) {
	if c.EngineStarted {
		return "", fmt.Errorf("cannot start engine already started")
	}

	return "engine started!", nil
}

func (c *Car) StopEngine() (string, error) {
	if !c.EngineStarted {
		return "", fmt.Errorf("cannot stop engine already stopped")
	}

	return "engine stopped!", nil
}

func (c *Car) MoveForwards(distance int) (string, error) {
	if !c.EngineStarted {
		return "", fmt.Errorf("cannot move with stopped engine")
	}

	return fmt.Sprintf("moved forward %d meters!", distance), nil
}

func (c *Car) MoveBackwards(distance int) (string, error) {
	if !c.EngineStarted {
		return "", fmt.Errorf("cannot move with stopped engine")
	}

	return fmt.Sprintf("moved backwards %d meters!", distance), nil
}

func (c *Car) TurnRight() (string, error) {
	if !c.EngineStarted {
		return "", fmt.Errorf("cannot turn right with stopped engine")
	}

	return fmt.Sprintf("turned Right!"), nil
}

func (c *Car) TurnLeft() (string, error) {
	if !c.EngineStarted {
		return "", fmt.Errorf("cannot turn left with stopped engine")
	}

	return fmt.Sprintf("turned Right"), nil
}
