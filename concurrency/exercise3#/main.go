package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(50)

	for i := 100.0; i < 150.; i++ {
		go func(val float64, wg *sync.WaitGroup) {
			x := math.Sqrt(val)
			fmt.Printf("Square root of %.2f is %.6f\n", val, x)

			wg.Done()
		}(i, &wg)

	}

	wg.Wait()

}
