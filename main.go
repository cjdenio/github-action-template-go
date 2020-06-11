package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Success! Your name is: %s\n", os.Args[1])
	input, exists := os.LookupEnv("INPUT_NAME")
	if exists == true {
		fmt.Printf("Env variable is %s\n", input)
	} else {
		fmt.Println("Env variable doesn't exist")
	}
}
