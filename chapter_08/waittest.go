// Reverb1 is a TCP server that simulates an echo.
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {

	done := make(chan bool)
	log.Println("Hey")
	go func() {
		fmt.Println("\t", "Hello")
		time.Sleep(1 * time.Second)
		fmt.Println("\t", "It's goroutin")
		time.Sleep(2 * time.Second)
		fmt.Println("\t", "Bye")
		done <- true
	}()
	<-done
	log.Println("Bye")
}
