package arrayslicemap

import (
	"fmt"
)

// BasicArray todo
func BasicArray() ([3]int, [2]float64) {
	var a [3]int
	a[1] = 21
	b := [...]float64{1.61, 3.14}
	for k, v := range b {
		b[k] *= v
	}
	return a, b
}

// BasicSlice todo
func BasicSlice() {
	odds := [5]int{7, 9, 11, 13, 15}
	var s1 = odds[:]
	var s2 = odds[1:]
	var s3 = odds[2:]
	var s4 = odds[:3]
	fmt.Printf("%#v, %#v, %#v, %#v, %#v\n\n", odds, s1, s2, s3, s4)
}

// BasicMap todo
func BasicMap() {
	cities := map[string]int64{"fremont": 10}
	cities[`sanjose`] = 20
	towns := make(map[string]int, 10)

	fmt.Printf("%#v --- %#v\n\n", cities, towns)
}
