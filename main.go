package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Success! Your name is: %s\n", os.Args[1])
}
