package main

import (
	"github.com/aabdullahgungor/product-api/server"
)

func main() {
	

	s := server.NewServer()
	s.Run()
}