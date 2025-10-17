package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("init")
	port := os.Getenv("PORT")
	if port == "" {
		port = string("8080")
	}
	fmt.Println("server started at port", port)
}
