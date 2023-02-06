package shark

import (
	"testing"

	"github.com/Lelo88/GoTestDouble/prey"
	"github.com/stretchr/testify/assert"
)

type mockWhiteShark struct{
	Err error
	WhiteShark whiteShark
}

func NewWhiteSharkService() *mockWhiteShark {
	return &mockWhiteShark{}
}

func (m *mockWhiteShark) Hunt(prey prey.Prey ) error{
	return m.Err
}

func TestHuntOrNot(t *testing.T){
	
	t.Run("hunted", func(t *testing.T){
		ser := mockWhiteShark{
			Err: nil,
			WhiteShark: whiteShark{
				speed: 14,
				position: [2]float64{12,12},
				//simulator: NewWhiteSharkService().WhiteShark.simulator.CanCatch(),
			},
		}

		newTuna := prey.CreateTuna()

		hunted :=ser.Hunt(newTuna)

		assert.NoError(t, hunted)
		assert.Equal(t, nil, hunted)
	})

	t.Run("not hunted", func(t *testing.T){

	})
}