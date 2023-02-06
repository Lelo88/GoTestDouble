package prey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpeedTotal(t *testing.T){

	

	t.Run("happy path", func(t *testing.T){
		var newTuna = tuna{maxSpeed: 100.00}
		speedactual := newTuna.GetSpeed()
		assert.NotEqual(t,newTuna.maxSpeed,speedactual)
	})

	t.Run("zero velocity", func(t *testing.T){
		var newTuna = tuna{maxSpeed: 0.0}
		speedactual := newTuna.GetSpeed()
		assert.Zero(t, newTuna.maxSpeed, speedactual)
	})

	t.Run("empty tuna", func(t *testing.T){
		var newTuna = tuna{}
		speedactual := newTuna.GetSpeed()
		assert.Empty(t, newTuna.maxSpeed, speedactual)
	})

	t.Run("error value negative int", func(t *testing.T){
		var newTuna = tuna{maxSpeed: -1}
		speedactual := newTuna.GetSpeed()
		assert.NotEqualValues(t, newTuna.maxSpeed, speedactual)
	})

	t.Run("error value positive int", func(t *testing.T){
		var newTuna = tuna{maxSpeed: 10}
		speedactual := newTuna.GetSpeed()
		assert.NotEqualValues(t, newTuna.maxSpeed, speedactual)
	})

	t.Run("error value string", func(t *testing.T){
		var newTuna = tuna{maxSpeed: 67}
		speedactual := newTuna.GetSpeed()
		assert.NotSame(t, "a", speedactual)
	})

}