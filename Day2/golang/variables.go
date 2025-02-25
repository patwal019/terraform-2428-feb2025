package main

import "fmt"

func main() {
	var inventor1 = "Ken Thompson"  //Declares a string implicitly and initial with some value
	inventor2  := "Rob Pile"        //Declares string variable initializing with some variable
	var inventor3 string            //Declares a string variable without initializing 
	inventor3 = "Robert Griesemer"  //initializing the string variable with a value

	fmt.Println ( "Golang inventors" )
	fmt.Println ( inventor1 )
	fmt.Println ( inventor2 )
	fmt.Println ( inventor3 )
}
