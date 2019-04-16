package main

import "fmt"

func main() {
	driver, session, err := GetNeo4jConnect()
	defer driver.Close()
	defer session.Close()

	result, err := Neo4jExec(session, "", map[string]interface{}{})
	for result.Next() {
		fmt.Println(result.Record().GetByIndex(0))
	}

	if err = result.Err(); err != nil {
		fmt.Println(err)
	}
}
