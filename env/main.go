package main

import (
	"fmt"
	"os"
)

// 010 OMIT
func main() {
	fmt.Printf("My password is %q\n", os.Getenv("PASSWD"))
}

// 020 OMIT
