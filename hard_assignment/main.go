package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error Occured while trying to open file: ", filename)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
