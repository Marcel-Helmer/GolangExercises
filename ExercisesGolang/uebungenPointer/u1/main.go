package main

import "fmt"

// double the ATK by use
func rage(atkPoints *int) {

	*atkPoints *= 2

}


func main() {

	atkPoints := 100

	fmt.Println("ATK vorher: ", atkPoints)

	rage(&atkPoints)

	fmt.Println("ATK danach: ", atkPoints)

}
