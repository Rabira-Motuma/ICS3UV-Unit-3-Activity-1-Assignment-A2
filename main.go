/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date 2025-11-14
 * This file finds and displays the sidelengths of a cube with a volume of 1000 mm³.
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	// INPUT - int data type volume
	var volume int = 1000

	// PROCESS
	// calculate side lengths
	var sideLength int = int(math.Cbrt(float64(volume)))

	// OUTPUT
	// display side lengths in statement
	fmt.Println("The length and width and height of a cube with a volume of 1000 mm³ is", sideLength, "mm." )

	fmt.Println("\nDone.")
	}
