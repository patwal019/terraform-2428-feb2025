package main

import (
  "fmt"
  "time"
)

func firstTask ( ch chan string ) {
   var msg string

   for i := 0; ; i++ {
     //Writing some message onto the channel - Send
     msg = fmt.Sprintf( "From Task 1 - %d", i )
     ch <- msg //Writing to the channel
     fmt.Println ( "Task 1 sent a new message" )
     time.Sleep ( time.Millisecond * 3 )
   }
}

func secondTask ( ch chan string ) {
   for {
     //Retrieving the message from the channel - Receive
     message := <- ch  //Read from the channel
     fmt.Println ( "Task 2 - received a message ", message )
     time.Sleep ( time.Millisecond * 3 )
   }
}

func main() {
   //Create a channel for the first and second tasks to communicate with each other
   //This channel will accept string type message
   var ch chan string = make(chan string)

   go firstTask(ch)
   go secondTask(ch)

   var waitForKeyPress string
   fmt.Println( "Press any key to quit ..." )
   fmt.Scanln( &waitForKeyPress )
}
