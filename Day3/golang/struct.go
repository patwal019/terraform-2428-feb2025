package main

import "fmt"


//user-defined data type that can have any number of variables inside and any type of variables
type Rectangle struct {
  length int
  width  int
}

//Golang Method - this is member of function of Rectangle
//Hence can only be invoked using a rectangle instance/object
func (rect Rectangle) Area() int {
   area := rect.length * rect.width
   return area
}

func (rect Rectangle) ComputeArea(l int, w int) int {
   rect.length = l
   rect.width  = w
   area := rect.length * rect.width
   return area
}

func main() {

	rectangle := Rectangle {
	   length: 100,
	   width: 200,
	}

	fmt.Printf( "Area of rectangle: %d\n", rectangle.Area() )
	fmt.Printf( "Area of rectangle: %d\n", rectangle.ComputeArea(10, 20) )

	//Area() // Will not work - can't be invoked outside the scope of rectangle
	//ComputeArea(10,20) //Will not work - can't be invoked without rectangle instance/object
}


