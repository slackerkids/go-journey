package main

import (
	"fmt"
	"github.com/slackerkids/go-journey/greetings"
)

func main()  {
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}