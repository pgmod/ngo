package main

import (
	"fmt"
	"ngo/internal/config"
	"os"
)

func main() {
	if err := realMain(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func realMain() error {
	return config.Execute()
}
