package main

import "fmt"

func main() {
	//golang array are fixed size
	var arr[5] int  //declares an array of int of capacity 5

	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50

	//arr[5] = 60 //This will report array index out of bounds

	fmt.Println ( "Array elements are ...")
	fmt.Println ( arr )
	fmt.Println ("Length of the array is ", len(arr) )

	arr1 := [5] int { 100, 200, 300, 400, 500, }  

	fmt.Println ( "Second Array elements are ...")
	fmt.Println ( arr1 )
}
