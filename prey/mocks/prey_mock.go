package mocks

import "github.com/stretchr/testify/mock"

//controller
func NewTunaMock() *tunaMock{
	return &tunaMock{}
}

//implementamos el mock de una estructura que ya esta inventada, en este caso, tuna
type tunaMock struct {
	mock.Mock
} 

//no recibe un argumento, pero que devuelva lo que va a pasar en el test
func (t *tunaMock) GetSpeed() float64{
	args := t.Called()
	return args.Get(0).(float64)
}

func (t *tunaMock) Reset() {
	t.Mock = mock.Mock{}
}