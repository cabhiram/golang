package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name     string
	PhoneNum string
}

func main() {
	p := Person{}
	p.Name = "Pooja"
	p.PhoneNum = "9591071205"
	WriteToFile(p)
	ReadFile()

}

func WriteToFile(p Person) {
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile("temp", b, 0666)
}

func ReadFile() {
	b, err := ioutil.ReadFile("temp")
	if err != nil {
		fmt.Println(err)
		return
	}
	q := Person{}
	json.Unmarshal(b, &q)
	fmt.Println(q)
}
