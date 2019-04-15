package main

import "fmt"

func main() {
	fmt.Println("Helllo")
	result, err := ConnectNeo4j()
	for result.Next() {
		fmt.Println(result.Record().GetByIndex(0))
	}

	if err = result.Err(); err != nil {
		fmt.Println(err)
	}
}