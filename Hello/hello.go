package hello

import (
	"example.com/greetings"
	"fmt"
)

func main() {
	message := greetings.Hello("Juan")
	fmt.Println(message)
}
