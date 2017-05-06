// reffer to https://gist.github.com/zyxar/2317744
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	//return fmt.Sprintf("cannot Sqrt negativ number: %g", float64(e))
	return fmt.Sprintf("cannot Sqrt negativ number: %g", e)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 1, ErrNegativeSqrt(f)
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
