package main

import "fmt"

func main() {
	var inventor1 = "Ken Thompson"  //Declares a string implicitly and initial with some value
	inventor2  := "Rob Pile"        //Declares string variable initializing with some variable
	var inventor3 string            //Declares a string variable without initializing 
	inventor3 = "Robert Griesemer"  //initializing the string variable with a value


	var exists = true   // boolean
	isSuccess := false

	fmt.Println ( exists, isSuccess )

	var x uint8 // -127 - 127
	//var y int8 //int16 uint8 uint16 uint32 uint64
	//var z int32

	x = 127
	fmt.Println( "Value of x is ", x )
	x = 255 
	fmt.Println( "Value of x is ", x )

	a := 1.2

	fmt.Printf( "%T\n", x )
	fmt.Printf( "%T\n", a )
	fmt.Printf( "%T\n", a )


	fmt.Println ( "Value of a is ", a )
	fmt.Println ( "Golang inventors" )
	fmt.Println ( inventor1 )
	fmt.Println ( inventor2 )
	fmt.Println ( inventor3 )
}
