// This functioın shows your cars fuel consumption in 2 ways
package main

import "fmt"

func main() {
	var k, l float32
	var depo, katYol, fulDepoFyt float32

	fmt.Println("Your car tank = ")
	fmt.Scan(&depo)
	fmt.Println("Road traveled until fuel runs out = ")
	fmt.Scan(&katYol)
	k = yakit(depo, katYol)
	fmt.Println(k, " = avarage fuel consumption per 100km")
	fmt.Println("------------------***-------------------")
	fmt.Println("Cost of fuel(whole tank) = ")
	fmt.Scan(&fulDepoFyt)
	fmt.Println("Road traveled until fuel runs out = ")
	fmt.Scan(&katYol)
	l = yakit2(fulDepoFyt, katYol)
	fmt.Println(l, "= Turkish liras Per Kilometre")

}
func yakit(x float32, y float32) float32 {
	var tuketim float32

	tuketim = (x / y) * 100 // x--> cars tank / y-->road traveled until fuel runs out

	return tuketim

}
func yakit2(a float32, b float32) float32 {

	var kmbasina float32
	kmbasina = a / b // a --> cost of fuel(whole tank)/ b--> road traveled until fuel runs out
	return kmbasina

}
