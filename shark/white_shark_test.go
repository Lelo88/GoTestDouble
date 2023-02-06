package shark

import (
	"testing"

	TunaMock "github.com/Lelo88/GoTestDouble/prey/mocks"
	SimulatorMock "github.com/Lelo88/GoTestDouble/simulator/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHunt(t *testing.T){
	//arrante8
	ms := SimulatorMock.NewCatchSimulatorMock()
	sh := CreateWhiteShark(ms) //acepta los metodos de la nueva interface



	//act 
	t.Run("ok", func(t *testing.T){
		//arrange
		
		//ahora vamos a inicializar una presa mockeada y no harcodeada
		prey := TunaMock.NewTunaMock()
		prey.On("GetSpeed").Return(0.0)
		ms.On("CanCatch", mock.Anything, mock.Anything, mock.Anything).Return(true)
		//on le indica al mock como tiene que reaccionar ante ese resultado
		

		//hay que chequear el metodo getlinear
		ms.On("GetLinearDistance", mock.Anything).Return(0.0)
		//act
		//aca llamamos al metodo para que el tiburon realice su caceria
		//el metodo me pide un prey, por lo que voy a tener que realizar un mockeado, ya que lo implementa tuna, para ver que pasa al pasar valores mas reales
		err := sh.Hunt(prey) 


		//assert
		assert.NoError(t, err)
		ms.AssertExpectations(t)
		prey.AssertExpectations(t)
	})

	//en el caso de que no haya llamado a todos los metodos (mas en ese orden) el test me iba a paniquear

	t.Run("error", func(t *testing.T){
		prey := TunaMock.NewTunaMock()
		prey.On("GetSpeed").Return(0.0)
		ms.On("CanCatch", mock.Anything, mock.Anything, mock.Anything).Return(false)
		//on le indica al mock como tiene que reaccionar ante ese resultado
		

		//hay que chequear el metodo getlinear
		ms.On("GetLinearDistance", mock.Anything).Return(0.0)
		//act
		//aca llamamos al metodo para que el tiburon realice su caceria
		//el metodo me pide un prey, por lo que voy a tener que realizar un mockeado, ya que lo implementa tuna, para ver que pasa al pasar valores mas reales
		err := sh.Hunt(prey) 


		//assert
		assert.Error(t, err)
		ms.AssertExpectations(t)
		prey.AssertExpectations(t)
	})
}