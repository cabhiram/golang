package main

import "fmt"

type ball []string

func (b ball) Print() []string {

	s := []string{}

	for i := 0; i < len(b); i++ {
		s = append(s, b[i])
		fmt.Println(s)
	}

	return s
}
