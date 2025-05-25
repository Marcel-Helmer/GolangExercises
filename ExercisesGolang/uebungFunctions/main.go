package main

import "fmt"

//  Menue

func menue() {

	fmt.Println("Choose and option")
	fmt.Println("--------------------")
	fmt.Println("1. Add item to inventory.")
	fmt.Println("2. Delete item off inventory.")
	fmt.Println("3. Show inventory.")
	fmt.Println("0. Exit")
}

func addItem(inventory *[]string, item string) {

	*inventory = append(*inventory, item)

}

func deleteItem(inventory *[]string) {

	for i := 0; i < len(*inventory); i++ {

		fmt.Printf("%v: %s \n", i+1, (*inventory)[i])

	}

	for {
		fmt.Println("Choose a number to delete an Item")

		s := 0

		fmt.Scan(&s)

		if s > 0 && s <= len(*inventory) {

			*inventory = append((*inventory)[:s-1], (*inventory)[s:]...)

			return

		} else {

			fmt.Println("Wrong input. Try again.")
		}

	}
}

func main() {

	inventory := []string{}

	for {

		menue()

		choose := 0

		fmt.Scan(&choose)

		switch choose {

		case 0:

			return

		case 1:

			fmt.Printf("Add an Item:")

			items := ""

			fmt.Scan(&items)

			addItem(&inventory, items)

		case 2:

			deleteItem(&inventory)

		case 3:

			for i := 0; i < len(inventory); i++ {

				fmt.Printf("%v: %s \n", i+1, inventory[i])

			}

		}
	}

}
