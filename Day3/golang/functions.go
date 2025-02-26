package main

import "fmt"

func sayHello( name string ) string {
	return "Hello, " + name + " !"
}
//pass by value
func anotherFunction( x int ) string {
	x++
	fmt.Println ( "Value of x is ", x )

	return "success"
}

//assume some complex computation happens here
//pass by pointer
//helps in improving application performance
//as the value of result is not copied instead the functions works on the original variable itself
func compute( x *int) {
   *x++
   fmt.Println ( "Value of x is ", *x )
}

func main() {
	fmt.Println ( sayHello( "Golang" ) )

	result := 0

	fmt.Println ( "Value of a before calling function is ", result )

	fmt.Println ( anotherFunction ( a ) )

	fmt.Println ( "Value of a after calling function is ", result )

	compute( &result )
	fmt.Println ( "Value of a after calling function is ", result )

	var ptr *int
	ptr = &result

	fmt.Println ( "Value of a is ", result )
	fmt.Println ( "Value of a as seen by ptr ", *ptr )

	result = 200
	fmt.Println ( "Value of a is ", result )
	fmt.Println ( "Value of a as seen by ptr ", *ptr )

	result++  // result = result + 1 it will result in value 201

	ptr++ // this will increment the value stored at address, instead it moves the pointer to next address value
	(*ptr)++ //dereference a pointer,the value stored at the address pointed by ptr has to be incremented

}
