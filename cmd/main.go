package main

import (
	"fmt"

	"github.com/arvindpunk/word-proximity-service/internal/handlers"
)

func main() {
	fmt.Println("Ligma")
	r := handlers.NewRouter()
	r.Run()
}
