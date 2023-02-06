package shark

import (
	"testing"

	"github.com/Lelo88/GoTestDouble/simulator/mocks"
)

func TestHunt(t *testing.T){
	//arrante8
	ms := mocks.NewCatchSimulatorMock()
	sh := CreateWhiteShark(ms) //acepta los metodos de la nueva interface



	//act 
	t.Run("ok", func(t *testing.T){

	})

	t.Run("error", func(t *testing.T){

	})
}