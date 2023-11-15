// main.go
package main

import (
	"log"
	"github.com/spf13/hello/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
