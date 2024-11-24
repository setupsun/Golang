// package main

// import (
// 	"fmt"
// 	"math"
// )

// func calcAge(yearFirst float64, yearSecond float64) string{
// 	calcWithMod := math.Mod(yearFirst, yearSecond)
// 	return fmt.Sprintf("Разница между годами: ", calcWithMod)
// }

// func main() {
// 	message := calcAge(1000, 2000)
// 	fmt.Println(message)
// }

package main

import (
	"fmt"
	// "math"
)

// func calcAge(yearFirst float64, yearSecond float64) string{
// 	calcWithMod := math.Mod(yearFirst, yearSecond)
// 	return fmt.Sprintf("Разница между годами: ", calcWithMod)
// }

func calcAge(yearFirst int, yearSecond int) string{
	calc := yearFirst - yearSecond
	if calc < 0 {
		calc = calc * (-1)
	}
	return fmt.Sprintf("Разница в годах: %d", calc)
}

func main() {
	message := calcAge(1000, 2000)
	fmt.Println(message)
}