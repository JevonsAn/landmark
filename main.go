package main

import "fmt"

func main() {
	driver, session, err := GetNeo4jConnect()
	if err != nil {
		println(err)
		defer driver.Close()
		defer session.Close()
	}

	ipRoad, relations, err := GetShortestLM(session, "8.8.8.8")

	if err != nil {
		println(err)
	} else {
		fmt.Println(ipRoad, relations)
	}
}
