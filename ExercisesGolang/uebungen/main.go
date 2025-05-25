// standard package is main

package main

// Import von Libarys (fmt ist die standart input output libary)

import ("fmt")


func main(){

	var name = "Marcel"
	var age = 32
	var age2 int  //Eingabe für den Scanner
	var name2 string  //Eingabe für den Scanner


	fmt.Println("Ich Lerne Go Code!")
	fmt.Println("Vairable ausgebeb: " + name)

	// 1. Das %v", muss dazu angegeben werden, wenn man einen text mit einem Zahlenwert kombiniert
	//sonst fehler
	// Printf wird dafür verwendet

	fmt.Printf("Hallo " + name + ", du bist %v Jahre alt\n", + age) //werte ansetzen

	fmt.Printf("Hallo %v, du bist %v Jahre alt\n", name,age) //Alternativ

	//Scanner bauen 

	fmt.Printf("Wie alt bist du?")

	fmt.Scan(&age2) // Eingabe um den wert für age2 zu bestimmten

	fmt.Printf("Ich bin %v Jahre alt \n", age2)

	fmt.Printf ("Wie heißt du?: ")

	fmt.Scan(&name2)

	fmt.Println("Ich Heiße " + name2)


	//if else unter GO unterscheidet sich nur gerinfügig von Java
	//Die bedingung steht hier nicht innerhalb einer Klammer

	if age2 < 18 {   // Die öffnende Klammer muss sich in der selben Zeile wie die if bedingung befinden sonst Fehler
		var diff = 18-age2
		var missing = "Jahre"

		//Variable anhand einer if bedingung unter bestimmten umständen ändern
		if diff == 1 {  
			missing = "Jahr"
		}


		if diff < 18 {  
			fmt.Printf( "%v ist Minderjährig, in %v %v bist du Volljährig \n", name2, diff,  missing)
		} 
		
	} else {
		fmt.Println(name2 + " Du bist Volljährig")
	}


 // Zahlenguesser mit for schleife

 var numb int
 var gues int

 fmt.Printf("Gib eine Zahl zum raten ein: ")

 fmt.Scan(&numb)

 fmt.Println("Welche Zahl wurde eingegeben")

 for   {
 fmt.Scan(&gues)


	if gues > numb {
		fmt.Println("Die Zahl ist zu groß")
	} 
	if  gues < numb {
		fmt.Println("Die Zahl ist zu klein")
	}

	if gues == numb{
		fmt.Println ("Die Zahl stimmt überein")
		break
	}

}




}