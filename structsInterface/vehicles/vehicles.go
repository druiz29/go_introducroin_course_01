package vehicles

import (
	"fmt"
)

type vehicle interface {
	Distance() float64
}

const (
	carVehicle        = "CAR"
	motorcycleVehicle = "MOTORCYCLE"
	truckVehicle      = "TRUCK"
	airplaneVehicle   = "AIRPLANE"
	gokuVehicle       = "GOKU"
)

func New(v string, time int ) (vehicle, error) {
	switch v {
	case carVehicle:
		return &Car{Time: time}, nil
	case motorcycleVehicle:
		return &motorcycle{Time: time}, nil
	case truckVehicle:
		return &truck {Time: time}, nil
	case airplaneVehicle:
		return 	&airplane {Time: time}, nil
	case gokuVehicle:
		return &goku {Time: time}, nil
	}
	return nil, fmt.Errorf("vehicle '%s' not exist: ", v)	
}

type Car struct {
	Time int
}

func (c *Car) Distance() float64 {
	return 10 * (float64(c.Time) / 60)
}

type motorcycle struct {
	Time int
}

func (m *motorcycle) Distance() float64 {
	return 120 * (float64(m.Time) / 60)
}

type truck struct {
	Time int
}

func (t *truck) Distance() float64 {
	return 70 * (float64(t.Time) / 60)
}

type airplane struct {
	Time int
}

func (a *airplane) Distance() float64 {
	return 800 * (float64(a.Time) / 60)
}

type goku struct {
	Time int
}
func (g *goku) Distance() float64 {
	return 700 * float64(g.Time) / 60
}
