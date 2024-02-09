package main

import "fmt"

func main() {
	/*
		%d : Integer
		%c : Character
		%f : Float
		%t : Boolean
		%s : String
		%o : Base 8
		%x : Base 16
		%v : Guesses based on data type
		%T : Type of supplied value
	*/
	fmt.Println("%9f\n", 3.14)
	fmt.Println("%.2f\n", 3.141592)
	fmt.Println("%9.f\n", 3.141592)

}
