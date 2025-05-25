package main

import "fmt"

// Das int nach der klammer gibt den Datentyp des wiedergabewertes (return) an

func addiere(a int, b int) (int, int, int, int) {


	addition := a + b
	subtraktion := a - b
	multiplikation := a * b
	division := a / b

	return addition, subtraktion, multiplikation, division

}

func main() {
	var a, b int
	var c int

	fmt.Println("Gib die erste Zahl ein")

	fmt.Scan(&a)

	fmt.Println("Gib die zweite Zahl ein")

	fmt.Scan(&b)



	fmt.Println("Was willst du tun?")
	fmt.Printf(" 1. %v + %v \n" , a , b)
	fmt.Printf(" 2. %v - %v \n" , a , b)
	fmt.Printf(" 3. %v * %v \n" , a , b)
	fmt.Printf(" 4. %v / %v \n" , a , b)

	fmt.Scan(&c)

	addition, subtraktion, multiplikation, division := addiere(a, b)



switch c {

case 1:

		

	fmt.Printf("Die addition von %v und %v ist: %v", a, b, addition)

	

case 2:

		fmt.Printf("Die subtraktion von %v und %v ist: %v", a, b, subtraktion)

	

case 3:

		fmt.Printf("Die multiplikation von %v und %v ist: %v", a, b, multiplikation)

	

case 4:
		fmt.Printf("Die division von %v und %v ist : %v" , a , b , division)

default:

	fmt.Println("Ungültige eingabe")


}




}
