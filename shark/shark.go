package shark

import "github.com/Lelo88/GoTestDouble/prey"


type Shark interface {
	Hunt(prey prey.Prey) error
}
