package main

import "fmt"

type character struct {
	Name  string
	Class string
	ATK   int
	DEF   int
	STR   int
	AGI   int
	INTE  int
}

type weapon struct {
	Type  string
	ATKb  int
	DEFb  int
	STRb  int
	AGIb  int
	INTEb int
}

func createCharacterList() []character {
	return []character{
		{Name: "Corrin", Class: "Princess", ATK: 10, DEF: 7, STR: 9, AGI: 6, INTE: 11},
		{Name: "Arthur", Class: "Berserk", ATK: 13, DEF: 8, STR: 11, AGI: 4, INTE: 2},
		{Name: "Charlotte", Class: "Berserk", ATK: 14, DEF: 6, STR: 12, AGI: 3, INTE: 3},
		{Name: "Laslowe", Class: "Hero", ATK: 10, DEF: 9, STR: 11, AGI: 8, INTE: 7},
	}
}

func createWeaponList() []weapon {

	return []weapon{
		{Type: "Sword", ATKb: 2, DEFb: 1, STRb: 1, AGIb: 2, INTEb: 0},
		{Type: "Shield", ATKb: 0, DEFb: 5, STRb: 2, AGIb: 0, INTEb: 0},
		{Type: "Axe", ATKb: 4, DEFb: -1, STRb: 3, AGIb: 0, INTEb: 0},
		{Type: "Spear", ATKb: 2, DEFb: 3, STRb: 0, AGIb: 4, INTEb: 0},
		{Type: "Wand", ATKb: 0, DEFb: -2, STRb: 0, AGIb: 2, INTEb: 6},
	}

}

func statIncrease(char []character, choose1 *int) {

	*choose1 = *choose1 - 1

	fmt.Println("Which stat of ", char[*choose1].Name, " do you wanna increase by 1?")

	fmt.Printf("1. STR: %v\n2. AGI: %v\n3. INTE: %v\n ", char[*choose1].STR, char[*choose1].AGI, char[*choose1].INTE)

	statChoose := 0

	fmt.Scan(&statChoose)

	switch statChoose {

	case 1:

		char[*choose1].STR += 1
		fmt.Println("STR succussfully increased!")

	case 2:

		char[*choose1].AGI += 1
		fmt.Println("Agi succussfully increased!")

	case 3:

		char[*choose1].INTE += 1
		fmt.Println("INTE succussfully increased!")

	}

}

func weaponEQ(chars []character, choose1 *int, weap []weapon) {

	*choose1 = *choose1 - 1

	fmt.Println("Which Weapon you wanna Equip for ", chars[*choose1].Name, "?")

	for j := 0; j < len(weap); j++ {

		fmt.Printf("%v: %s\n", j+1, weap[j].Type)

	}

	chooseW := 0

	fmt.Scan(&chooseW)

	chars[*choose1].ATK += weap[chooseW-1].ATKb
	chars[*choose1].DEF += weap[chooseW-1].DEFb
	chars[*choose1].STR += weap[chooseW-1].STRb
	chars[*choose1].AGI += weap[chooseW-1].AGIb
	chars[*choose1].INTE += weap[chooseW-1].INTEb

}

func showStats(chars []character, choose1 int) {

	choose1 = choose1 - 1

	fmt.Println("The class of ", chars[choose1].Name, " is '", chars[choose1].Class, "' and the Stats are:")

	fmt.Printf("Without Classboost: ATK: %v | DEF: %v | STR: %v | AGI: %v | INTE: %v\n",
		chars[choose1].ATK, chars[choose1].DEF, chars[choose1].STR, chars[choose1].AGI, chars[choose1].INTE)

	fmt.Printf("With Classboost: ")

	if chars[choose1].Class == "Princess" {

		fmt.Printf("ATK: %v | DEF: %v | STR: %v | AGI: %v | INTE: %v\n",
			chars[choose1].ATK+1, chars[choose1].DEF+0, chars[choose1].STR+1, chars[choose1].AGI+2, chars[choose1].INTE+4)

	} else if chars[choose1].Class == "Berserk" {

		fmt.Printf("ATK: %v | DEF: %v | STR: %v | AGI: %v | INTE: %v\n",
			chars[choose1].ATK+5, chars[choose1].DEF+2, chars[choose1].STR+3, chars[choose1].AGI-1, chars[choose1].INTE+0)

	} else if chars[choose1].Class == "Hero" {

		fmt.Printf("ATK: %v | DEF: %v | STR: %v | AGI: %v | INTE: %v\n",
			chars[choose1].ATK+2, chars[choose1].DEF+4, chars[choose1].STR+1, chars[choose1].AGI+3, chars[choose1].INTE+2)

	}

}

func menue2() {

	fmt.Println("What you wish to do with your character?")
	fmt.Println("----------------------------------------")
	fmt.Println("1. Check Stats.")
	fmt.Println("2. Equip Weapon.")
	fmt.Println("3. Increase Stats.")
	fmt.Println("0. Back to character Select")
	fmt.Printf("Pick a number:")

}

func menue1(chars []character) {

	fmt.Println("Choose a Character")
	fmt.Println("-----------------")

	for i := 0; i < len(chars); i++ {

		fmt.Printf("%v: %s \n", i+1, chars[i].Name)

	}

	fmt.Println("0. End Programm")

	fmt.Printf("Pick a number: ")
}

func main() {

	charList := createCharacterList()
	weaponList := createWeaponList()

	for {

		menue1(charList)

		var choose1 int = 0

		for {

			fmt.Scan(&choose1)

			if choose1 > 0 && choose1 <= len(charList) {

				fmt.Println("You choose: ", charList[choose1-1].Name)

				break

			} else if choose1 < 0 && choose1 > len(charList) {
				fmt.Println("Invalid Input. Again!")
			} else {
				break
			}

		}

		if choose1 == 0 {
			fmt.Println("End Programm!")
			break
		}

		menue2()

		var choose2 int = 0

		fmt.Scan(&choose2)

		switch choose2 {

		case 1:

			showStats(charList, choose1)

		case 2:

			weaponEQ(charList, &choose1, weaponList)

			fmt.Printf("New base Stats without Class boost\n ATK: %v\n DEF: %v\n STR: %v\n AGI: %v\n INTE %v\n",
				charList[choose1].ATK, charList[choose1].DEF, charList[choose1].STR, charList[choose1].AGI, charList[choose1].INTE)

		case 3:

			statIncrease(charList, &choose1)

		}

	}

}
