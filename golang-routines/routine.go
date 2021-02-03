package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{"http://google.com", "http://amazon.com", "http://linkedin.com"}
	c := make(chan string)
	go getStatus(websites, c)
	for i := 0; i < len(websites); i++ {
		fmt.Println(<-c)
	}
}

func getStatus(w []string, c chan string) {
	for i := 0; i < len(w); i++ {
		res, _ := http.Get(w[i])
		if res.Status == "" {
			fmt.Println("Error while fetching data")
		}
		c <- res.Status + " " + w[i]

	}

}
