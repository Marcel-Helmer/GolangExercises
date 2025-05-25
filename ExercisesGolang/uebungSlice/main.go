//xml datei


package main

import (
	"fmt"
)

func menue (){

	fmt.Println("What you like to do?")

	fmt.Println("1. See the average of all grades")
	fmt.Println("2. See the highest grade only")
	fmt.Println("3. See the lowest grade only")
	fmt.Println("4. Add a Grade")
	fmt.Println("5. Delete the last Grade")
	fmt.Println("0. Close")

}





func main() {

	studentGrade := []int{74, 49, 88, 94, 64}

	for {

		fmt.Println("Grade of all Students = ", studentGrade)

		totalGrade := 0
		highestGrade := studentGrade[0]
		lowestGrade := studentGrade[0]

		for i := 0; i < len(studentGrade); i++ {

			totalGrade += studentGrade[i]

			if highestGrade < studentGrade[i] {

				highestGrade = studentGrade[i]

			}

			if lowestGrade > studentGrade[i] {

				lowestGrade = studentGrade[i]

			}

		}

		menue()

		

		ch := 0

		fmt.Scan(&ch)

		switch ch {

		case 0:
			fmt.Println("Close.")

			return

		case 1:

			fmt.Println("The average of all grades is: ", totalGrade/len(studentGrade))

		case 2:

			fmt.Println("The highest grade is: ", highestGrade)

		case 3:

			fmt.Println("The lowest grade is: ", lowestGrade)

		case 4:

			fmt.Printf("Type a Grade to add: ")

			add := 0

			fmt.Scan(&add)

			studentGrade = append(studentGrade, add)

			fmt.Println("Add grade to slice: ", studentGrade)

		case 5:

			fmt.Println("Choose a grade to delete: ")

			for j := 0; j < len(studentGrade); j++ {

				fmt.Println((j + 1), ": ", studentGrade[j])
			}

			gradeNo := 0

			fmt.Scan(&gradeNo)

			if gradeNo > 0 && gradeNo < len(studentGrade)+1 {

				deleteNo := studentGrade[gradeNo-1]

				for m := 0; m < len(studentGrade); m++ {

					if deleteNo == studentGrade[m] {
						studentGrade[m] = studentGrade[len(studentGrade)-1]

						break
					}

				}

				studentGrade = studentGrade[:len(studentGrade)-1]

				fmt.Println("Delete grade to slice: ", studentGrade)

			} else {

				fmt.Println("Falsche eingabe")

			}

		default:

			fmt.Println("Wrong Input, again.")

		}

	}

}


// Slices bsp

/* smurf := []int{1, 2, 3}
smurf = append(smurf, 4) // Dynamisch hinzufügen
fmt.Println(smurf) // [1 2 3 4] */


