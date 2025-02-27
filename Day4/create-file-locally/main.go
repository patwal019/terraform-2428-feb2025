package main

import "os" 

func main() {
	fileName := "./testfile.txt"
	myfile, err := os.Create( fileName )

	if err != nil {
		panic(err)
	}
	
	content := "This is a test file"

	myfile.WriteString( content + "\n" )
	myfile.Sync()
}
