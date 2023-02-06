	package shark

import (
	"fmt"
	"math/rand"

	"time"

	"github.com/Lelo88/GoTestDouble/prey"
	"github.com/Lelo88/GoTestDouble/simulator"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type whiteShark struct {
	// speed in m/s
	speed float64
	// position of the shark in the map of 500 * 500 meters
	position [2]float64
	// simulator
	simulator simulator.CatchSimulator
}

//verifica si un tiburon puede cazar a una presa de acuerdo a su velocidad y posicion 
func (w *whiteShark) Hunt(prey prey.Prey) error {
	if w.simulator.CanCatch(w.simulator.GetLinearDistance(w.position), w.speed, prey.GetSpeed()) {
		fmt.Println("ñam ñam")
		return nil
	}
	return fmt.Errorf("could not hunt the prey")
}

func CreateWhiteShark(simulator simulator.CatchSimulator) Shark {
	return &whiteShark{
		speed:     144,
		position:  getInitialPosition(),
		simulator: simulator,
	}
}

func getInitialPosition() [2]float64 {
	x := rand.Float64() * 500
	rand.Seed(time.Now().UnixNano())
	y := rand.Float64() * 500
	return [2]float64{x, y}
}
