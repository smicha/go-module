package random

import (
	"fmt"
	"math/rand"
)

func random() string {

	return fmt.Sprintf("from random module %f ",rand.Float32())
}