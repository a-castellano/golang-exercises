package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func main() {
	filename := os.Args[1]
	var read []byte
	read, _ = ReadFile(filename)
	fmt.Println(read[0])
}
