package main

import "fmt"

func main() {
//                        key    value
// key and value data type can be different or same, in this case they are same
	//Balanced binary search tree(BST) - e.g red black bst
	//Map can retrieve values very fast worst case scenario runtime efficiency log n
	//1024 contacts are stored in your mobile
	//What is the maximum number of comparison contacts app has to make to pick a number
	//10 search is enough to say a particular contact is there in your contacts or not
	//2048 contacts are stored in your mobile, what is the extra search required to find a number
	//11 search is enough to report a number is present or not present
	toolsPath := map[string]string {
		"java_home": "/usr/lib/jdk-11",
		"mvn_home" : "/usr/share/maven",
	}

	//Given a key, map will be able to retrieve the value
	fmt.Println ( "Java home directory :", toolsPath["java_home"] )

	//add a new key,value pair into an existing map
	toolsPath["golang_home"] = "/usr/bin/go"

	//iterating a map and printing its value
	for key,value := range toolsPath {
		fmt.Println ( key, value )
	}

	//delete a key-value pair
	delete ( toolsPath, "golang_home" )
}
