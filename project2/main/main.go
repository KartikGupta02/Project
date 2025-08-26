package main

import (
	"GoTutorial/project2/greetings"
	"fmt"
)

func main() {
	msg := greetings.Hello("Kartik")
	fmt.Println(msg)
}
