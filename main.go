package main

import (
	"fmt"
	"os"

	"github.com/mducoli/cronupper/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
