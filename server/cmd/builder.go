package main

import (
	"os"
)

func main() {
	err := os.Rename("ui/build", "server/public")
	if err != nil {
		panic(err)
	}
}
