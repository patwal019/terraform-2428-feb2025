package main

import "fmt"

func main() {
	//Understanding golang slice
	//golang slice works like a dynamic array

	intArray := [6] int { 10, 20, 30, 40, 50, 60 }

	//mySlice is go refer elements in the array intArray from index 2 ie 30 to 50 ]
	var mySlice []int = intArray[2:5] // { 30, 40, 50 } in this case slice is reference to the subarray {30, 40, 50}

	fmt.Println ( "Array elements are ..." )
	fmt.Println ( intArray )

	fmt.Println ( "Slice elements are ...", mySlice )

	//Modifying value
	mySlice[0] = 100  //this will modify the original array intArray[2]
	mySlice[1] = 200  //this will modify intArray[3]
	mySlice[2] = 300  //this will modify intArray[4]
	//mySlice[3] = 400 //reports index out of bounds error

	fmt.Println ( "Array elements after modifying slice elements are ..." )
	fmt.Println ( intArray )  
	//Slice uses array internally, hence it will create a new array of size 3 + 2 to store {100, 200, 300, 400, 500 }
	mySlice = append ( mySlice, 400, 500 ) //We are to append {100, 200, 300 } with {400, 500 }
	mySlice[0] = 1
	mySlice[1] = 2
	mySlice[2] = 3

	fmt.Println ( "After modifying Slice elements are ...", mySlice )
	fmt.Println ( "Array elements after modifying slice elements are ...", intArray )
}
