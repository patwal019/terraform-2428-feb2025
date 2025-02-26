package main

import "fmt"


//user-defined data type that can have any number of variables inside and any type of variables
type Rectangle struct {
  length int
  width  int
}

//Golang Method
func (rect Rectangle) Area() int {
   area := rect.length * rect.width
   return area
}

func main() {

	rectangle := Rectangle {
	   length: 100,
	   width: 200,
	}

	fmt.Printf( "Area of rectangle: %d\n", rectangle.Area() )
}


