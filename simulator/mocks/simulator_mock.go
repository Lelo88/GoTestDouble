package mocks

import "github.com/stretchr/testify/mock"

//aca voy a crear una implementacion ficticia para que no se utilice la verdadera

//constructor = representara la interfaz simulator
func NewCatchSimulatorMock() *catchSimulatorMock {
	return &catchSimulatorMock{}
}

//controller
type catchSimulatorMock struct {
	mock.Mock //vamos a utilizar la libreria mock testify
}

func (c *catchSimulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	args := c.Called(distance, speed, catchSpeed) 
	return args.Bool(0)
}

func (c *catchSimulatorMock) GetLinearDistance(position [2]float64) float64 {
	args := c.Called(position)
	return args.Get(0).(float64)
}

//reset
func (c *catchSimulatorMock) Reset() {
	c.Mock = mock.Mock{}
}

//estas implementaciones vana a devolver los resultados de la misma forma que lo hace simulator.go