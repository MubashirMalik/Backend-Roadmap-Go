package main

import "fmt"

func main() {
    var x int = 19
    var y float32 = 21

		// var mul float32 = x * y
    var mul float32 = float32(x) * y
		//var mul int = x * int(y)

    // Displaying the result
    fmt.Printf("Multiplication = %f\n", mul)
}