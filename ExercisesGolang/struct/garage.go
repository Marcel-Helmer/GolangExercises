package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
	Color string
}

func PrintCar(car Car) {

	fmt.Printf("Brand: %v  Model: %v  Year: %v  Color %v \n", car.Brand, car.Model, car.Year, car.Color)
}

func main() {

	// Der Slice (eine liste aus denen man objekte erstellt ähnlich wie eine ArrayList in java)
	// wir in die main erstellt und man kann einfach objekte erstellen
	garage := []Car{
		{Brand: "BMW", Model: "X5", Year: 2020, Color: "Black"},
		{Brand: "Tesla", Model: "Model S", Year: 2022, Color: "White"},
	}

	fmt.Println("Autos in der Garage:")
	for _, car := range garage {
		PrintCar(car) // Aufruf der Funktion für jedes Auto

	}
}
