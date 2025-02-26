# Day 3

## Lab - Understanding Golang if else control statement
```
cd ~/terraform-2428-feb2025
git pull
cd Day3/golang
cat if-else.go
go run ./if-else.go
```

Expected output
![image](https://github.com/user-attachments/assets/394b73d8-bf59-4b1a-940a-598b1d93aead)

## Lab - Understanding golang loops
```
cd ~/terraform-2428-feb2025
git pull
cd Day3/golang
cat loops.go
go run ./loops.go
```

Expected output
![image](https://github.com/user-attachments/assets/d356af06-a918-4ac0-8008-d895468b0423)

![image](https://github.com/user-attachments/assets/19e15146-9aad-47a1-b1b8-bb1177aae778)

![image](https://github.com/user-attachments/assets/b45e1b45-74c6-4d31-8178-4dc4c7caad03)

## Lab - Understanding golang arrays
Golang array are fixed in size, they won't grow dynamically at runtime.

```
cd ~/terraform-2428-feb2025
git pull
cd Day3/golang
cat arrays.go
go run ./arrays.go
```

Expected output
![image](https://github.com/user-attachments/assets/b24ef271-96ae-4eea-a74a-14ddc8a51ffa)

## Lab - Understanding golang slice
```
cd ~/terraform-2428-feb2025
git pull
cd Day3/golang
cat slice.go
go run ./slice.go
```

Expected output
![image](https://github.com/user-attachments/assets/c38f2624-7ec7-4444-a96d-d3e759e14a9a)

![image](https://github.com/user-attachments/assets/d3f6e608-e58c-4a65-afec-0ec57ac9a510)


## Lab - Understanding golang map
```
cd ~/terraform-2428-feb2025
git pull
cd Day3/golang
cat map.go
go run ./map.go
```

Expected output
![image](https://github.com/user-attachments/assets/93998a44-1804-4214-993f-6a3e36768aea)
![image](https://github.com/user-attachments/assets/f925c7bc-c6ae-4eee-a73f-421755e2b793)


## Lab - Understanding golang struct
```
cd ~/terraform-2428-feb2025
git pull
cd Day3/golang
cat struct.go
go run ./struct.go
```

Expected output
![image](https://github.com/user-attachments/assets/109650a7-9835-4579-8c8a-198555d4e39b)

## Lab - Understanding golang concurrency
```
cd ~/terraform-2428-feb2025
git pull
cd Day3/golang
cat concurrency.go
go run ./concurrency.go
```

Expected output
![image](https://github.com/user-attachments/assets/5301b537-553c-4f1b-813b-97fa2719c150)
![image](https://github.com/user-attachments/assets/ec9af359-ab23-4e5c-b8fb-14d9ef4fb5c3)



## Lab - Understanding golang channels and concurrency
```
cd ~/terraform-2428-feb2025
git pull
cd Day3/golang
cat channel.go
go run ./channel.go
```

Expected output
![image](https://github.com/user-attachments/assets/82ae699d-040e-4476-a93c-247d6a82573f)
![image](https://github.com/user-attachments/assets/b73b00e8-8e13-468f-b3d2-512e41b96722)
![image](https://github.com/user-attachments/assets/6265b58e-eabd-4e1b-a005-1a21263fe124)

## Info - Golang Package
<pre>
- is a set of reusable related functions
- example
  - fmt package supports many functions for input/output operations
</pre>  

## Info - Golang Module
<pre>
- is a collection many golang packages 
- reusable code
</pre>

## Lab - Creating a custom go modules

In this exercise, we will be creating two custom go modules, namely hello and tektutor.  Hence in your home directory, create two folders namely hello and tektutor as shown below

```
cd ~
mkdir hello tektutor
cd hello
go mod init tektutor.org/hello
cat go.mod
go mod tidy
```

Now let's create a file named hello.go under hello folder as shown below
<pre>
package hello

import "fmt"

func SayHello( name string ) string {
  message := fmt.Sprintf( "Hello, %v !", name )
  return message
}
</pre>

Now, step out of hello folder and navigate to tektutor folder
```
cd ../tektutor
go mod init tektutor.org/tektutor
cat go.mod
go mod tidy
```

Let's create a main.go file with the below content under tektutor folder
<pre>
package main

import (
   "fmt"
   "tektutor.org/hello"
)

func main() {
  msg := hello.SayHello ( "Golang" )
  fmt.Println(msg)
}
</pre>

Now, let's run the below command under tektutor folder
```
cd ~/tektutor
go mod tidy
go mod edit --replace tektutor.org/hello=../hello
go mod tidy
go run ./main.go
```

Expected output
![image](https://github.com/user-attachments/assets/4de40444-89f3-4b9e-b72d-b7dee4d280d0)
![image](https://github.com/user-attachments/assets/ded526e0-8b07-4511-bd60-cd747a6d4fbd)
