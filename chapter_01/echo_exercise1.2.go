// Echo2 prints its command-line arguments.

package main

import (
	"fmt"
	"os"
    "strconv"
)

func main() {
	var s string
    for index, arg := range os.Args[1:] {
        s += "Index: " + strconv.Itoa(index) + " Value: " + arg + "\n"
	}
	fmt.Println(s)
}
