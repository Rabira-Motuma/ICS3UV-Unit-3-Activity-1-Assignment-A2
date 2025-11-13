/**
 * @author Rabira Motuma
 * @version 1.0.0
 * @date 2025-11-11
 * @fileoverview This file finds and displays the sidelengths of a cube with a volume of 1000 mm³.
 */

// INPUT - number data type volume
let volume : number = 1000;

// PROCESS
// calculate sideLength of cube
const sideLength : number = Math.cbrt(volume);

// OUTPUT
// display the result
console.log("The length and width and height of a cube with a volume of 1000 mm³ is " + sideLength + " mm.")

console.log("\nDone")
