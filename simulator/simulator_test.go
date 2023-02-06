package simulator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type simulatorMock struct{
	CanCatchbool 	bool
	GetLinearF 		float64
	Error 			error
	Spy				bool
}


func (s *simulatorMock) CanCatch(distance float64, speed float64, catchSpeed float64) bool {
	return s.CanCatchbool
}

func (s *simulatorMock) GetLinear(position [2]float64) float64 {
	s.Spy = true
	return s.GetLinearF
}


func TestCanCatchAll(t *testing.T){
	t.Run("prey catched true", func(t *testing.T){
		ser := simulatorMock{
			CanCatchbool: true,
		}
		s := catchSimulator{maxTimeToCatch: 10}
		distance := 11
		speed := 12
		catchSpeed := 10
		
		timeToCatch := distance / (speed - catchSpeed)
		timeToCatchBool := timeToCatch>0
		canCatch := ser.CanCatch(float64(distance),float64(speed),float64(catchSpeed))
		
		assert.Equal(t, s.maxTimeToCatch>=float64(timeToCatch),canCatch)
		assert.Equal(t, timeToCatchBool, canCatch)
	})

	t.Run("prey catched false", func(t *testing.T){
		ser := simulatorMock{
			CanCatchbool: true,
		}
		s := catchSimulator{maxTimeToCatch: 15}
		distance := 20
		speed := 1
		catchSpeed :=0 
		
		timeToCatch := distance / (speed - catchSpeed)
		maxTimeToCatch := s.maxTimeToCatch>=float64(timeToCatch)

		timeToCatchBool := timeToCatch>0
		canCatch := ser.CanCatch(float64(distance), float64(speed), float64(catchSpeed))
		
		assert.NotEqual(t, maxTimeToCatch, canCatch)
		assert.Equal(t, timeToCatchBool, canCatch)
	})

	t.Run("error catched velocity", func(t *testing.T){
		//s := catchSimulator{maxTimeToCatch: 15}
		
		ser := simulatorMock{
			Error: errors.New("Division by zero"),
		}
		distance := 11
		speed := 12
		catchSpeed := 12

		//timeToCatch := distance / (speed - catchSpeed)
		//timeToCatchBool := timeToCatch>0
		canCatch := ser.CanCatch(float64(distance), float64(speed), float64(catchSpeed))
		dividend := speed - catchSpeed
		

		//assert.NotNil(t, timeToCatchBool, canCatch)
		assert.Error(t, ser.Error)
		assert.Error(t, ser.Error, canCatch)
		assert.Equal(t, 0, dividend)
		//assert.NotEqual(t,s.maxTimeToCatch>=float64(timeToCatch), canCatch)
		//assert.NotEqual(t, s.maxTimeToCatch>=float64(timeToCatch), canCatch)
	})
}

func TestGetLinearAll(t *testing.T){

	t.Run("getting linear", func(t *testing.T){
		ser := simulatorMock{
			GetLinearF: 4368,
		}

		var mockFloat [2]float64

		mockFloat[0] = 78
		mockFloat[1] = 56

		getline := ser.GetLinear(mockFloat)

		mockfloatexpectation := mockFloat[0] * mockFloat[1]

		assert.Equal(t,mockfloatexpectation,getline)
		assert.True(t,ser.Spy)
	})

	t.Run("not getting linear", func(t *testing.T){
		ser := simulatorMock{
			GetLinearF: 13,
			Spy: false,
		}

		var mockFloat [2]float64

		mockFloat[0] = 78
		mockFloat[1] = 56

		getline := ser.GetLinear(mockFloat)

		mockfloatexpectation := ser.GetLinearF

		assert.Equal(t,mockfloatexpectation,getline)
		assert.True(t,ser.Spy)
	})
}