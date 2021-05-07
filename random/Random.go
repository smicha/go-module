package random

import (
	"fmt"
	"math/rand"
)

func GetRandom() string {

	return fmt.Sprintf("from random module %f ",rand.Float32())
}