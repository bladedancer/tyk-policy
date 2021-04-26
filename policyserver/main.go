package main

import (
	"fmt"
	"os"

	"github.com/bladedancer/tyk-policy/policyserver/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
