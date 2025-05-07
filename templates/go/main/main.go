
package main

import (
	"fmt"
	"os"
)

func main() {
	if err := realMain(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func realMain() error {
	fmt.Println("Hello, World!")
	return nil
}
