package main 

import "fmt"

func main() {

	//var count int //Declares count var
	//count = 5 //initialize

	//var count = 5

	count := 5  //Declare count int variable and assigning value 5

	for count > 0 {
		fmt.Println ( count )
		count--  // equivalent to count = count - 1
		// --count ++count not supported in golang
	}

	fmt.Println ("Value of count is ", count, " after for loop ")

	for count = 1; count < 10; count++ {
	    fmt.Printf("%d\t", count )
	}

	str := ""
	for i := 0; i<5; i++ {
	    str = fmt.Sprintf("some message %d-%s", 100, "abc" )
	    fmt.Println( str )
	}
/*
	for { //never ending for loop
		fmt.Println ("This will keep running, press Ctrl + C to exit the application" )
	}
*/

	//do while loop 
	
	count = 10

	for {
		if count < 10 {
		   continue  //will force the for loop to move on to the next iteration
		} else if count > 10 {
		   count++ 
		} else {
		   break
	        }	
		  
		fmt.Println ("Inside for loop ...")
	}

	fmt.Println() // leave a blank line 
}
