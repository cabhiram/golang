package main

func main() {
	var Person User

	Person.name = "Abhiram"
	Person.phoneNumber = "9591071205"

	PersonPointer := &Person
	PersonPointer.Modify("Sourabh")
	Person.Print()
}
