package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("yo")
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://github.com",
		"http://golang.org",
		"http://groundfloor.us",
	}

	c := make(chan string)

	for _, link := range links {
		go linkChecker(link, c)
	}

	for l := range c {
		go func(anon_link string) {
			time.Sleep(5 * time.Second)
			linkChecker(anon_link, c)
		}(l)
	}
}

func linkChecker(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%v might be down\n", link)
		c <- link
		return
	}
	fmt.Printf("%v is up an running!\n", link)
	c <- link
}
