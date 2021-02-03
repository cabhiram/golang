package main

import (
	"encoding/json"
	"fmt"
)

type AdventureList struct {
	Adventures []Adventure `json:"adv"`
}
type Adventure struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	var s = `{"adv":[{"id":1, "name":"title_1"},{"id":2, "name":"title_2"}]}`
	a := AdventureList{}
	json.Unmarshal([]byte(s), &a)
	for i := 0; i < len(a.Adventures); i++ {
		fmt.Println(a.Adventures[i])
	}
}
