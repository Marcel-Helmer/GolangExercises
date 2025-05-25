package main

import "fmt"

func main() {

	mailadress := map[string]string{
		"Leo":    "leo@web.de",
		"Elise":  "eliseQT@gmail.com",
		"Arthur": "arthur.hero@yahoo.com",
		"Corrin": "Corrin92@gmail.com",
	}

	for {

		fmt.Println("What you wanna do?")
		fmt.Println(" --------------------- ")
		fmt.Println("1. Check a Mail adress.")
		fmt.Println("2. Add a user.")
		fmt.Println("3. Delete a user.")
		fmt.Println("0. Exit")

		menue := 0

		fmt.Scan(&menue)

		switch menue {

		case 0: 

		return

		case 1:

			fmt.Println("Which mail you wanna check?")

			seachMail := ""

			fmt.Scan(&seachMail)

			if mail, check := mailadress[seachMail]; check {

				fmt.Printf("The mail address of %s is: %s\n", seachMail, mail)
			} else {

				fmt.Println("Name not found!")
			}

		case 2:

			addName := ""

			addMail := ""

			fmt.Println("Add a name to the map.")

			fmt.Scan(&addName)

			fmt.Println("Add a mail adress to the map.")

			fmt.Scan(&addMail)

			mailadress[addName] = addMail

		case 3:

			fmt.Println("Delete a name off the map.")

			deleteName := ""

			fmt.Scan(&deleteName)

			delete(mailadress, deleteName)

			for name, mail := range mailadress {

				fmt.Printf("%s: %s \n", name, mail)

			}

		}
	}
}

//Maps bsp

/* studentGrades := map[string]int{
    "Alice": 90,
    "Bob":   85,
    "Charlie": 70,

	for name, number := range telefonnummern {

        fmt.Printf("%s: %d\n", name, number)

    }
} */
