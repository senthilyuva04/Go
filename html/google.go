package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	fmt.Println(resp, err)
	fmt.Println("Getting Read Value")
	bs := make([]byte, 32*1024)
	n, e := resp.Body.Read(bs)
	fmt.Println(string(bs[0:n]))
	fmt.Printf("\n")
	fmt.Println(n, e)

}
