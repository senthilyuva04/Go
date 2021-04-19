package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://google.com", "http://linkedin.com", "http://facebook.com", "http://yahoo.com", "http://amazon.com", "http://flipkart.com"}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error Accessing link ", link)
		c <- link
		return
	}
	fmt.Println("Link ", link, " is Available")
	c <- link
}
