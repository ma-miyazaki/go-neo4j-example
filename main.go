package main

import (
	"fmt"

	"github.com/ma-miyazaki/go-neo4j-example/internal"
)

func main() {
	fmt.Printf("Hello world.\n")

	defer internal.Neo4j.Close()

	_, err := internal.Neo4j.DeleteAll()
	if err != nil {
		fmt.Println("err", err)
		return
	}
}
