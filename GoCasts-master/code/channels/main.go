package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for i := 0; i < len(links); i++ {
		go func() {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}()
	}

	//for l := range c {
	//	go func(link string) {
	//		time.Sleep(5 * time.Second)
	//		checkLink(link, c)
	//	}(l)
	//}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "down")
		c <- "Might be down"
		return
	}

	fmt.Println(link, "connected")
	c <- "Yep its up"
}
