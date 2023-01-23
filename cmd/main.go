package main

import (
	"fmt"
	"groupie-tracker/server"
)

func main() {
	fmt.Println("http://localhost:8080/")
	server.Run()
}
