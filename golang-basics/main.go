package main

import (
	"fmt"
	"strconv"
)

func main() {

	var s ball

	for i := 0; i < 5; i++ {
		s = append(s, "abhiram"+strconv.Itoa(i))
	}

	l := s.Print()
	fmt.Println(l)
}
