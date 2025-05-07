package main

import (
	"fmt"

	"os"

	"github.com/pgmod/ngo/internal/config"
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
