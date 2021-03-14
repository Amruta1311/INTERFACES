package main

import "fmt"

//Interfaces helps in solving the issue of reusing one function between two different data types

type bot interface { //Helps condense down to a one single implementation
	//New custom type bot says that if you are a type in this program
	//with a function getGreeting and you return a string then
	//you are now an honorary member of type "bot"
	getGreeting() string //string is the return type
}

type englishBot struct{}

type spanishBot struct{}

func main() {

	eb := englishBot{}

	sb := spanishBot{}

	printGreeting(eb)

	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//VERY custom logic for generating an english greeting

	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

//Rule Of Interfaces
// type bot interface {
// 	getGreeting(string, int) (string, error)
// 	//string, int denote the list of arguement types
// 	//string, error denotes the list of return types
// }

//Concrete types are the data types from which we can create a value using each of them
//Interface type are the data types from which we cannot directly create a value

//Interfaces are not generic types

// Interfaces are implicit: We don't manually have to say that our custom type satisfies some interface

//Interfaces are a contract to help us manage types

//Understand how to read interfaces in the standard lib.
