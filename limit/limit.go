package limit

import "fmt"

type Limit struct {
	Lower  int
	Higher int
}

func (l Limit) String() string {
	return fmt.Sprintf("Lower: %d, Higher: %d", l.Lower, l.Higher)
}
