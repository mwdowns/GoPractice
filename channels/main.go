package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("yo")
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://github.com",
		"http://golang.org",
	}

	for _, link := range links {
		linkChecker(link)
	}
}

func linkChecker(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}
	fmt.Println(link, "is up!")
}
