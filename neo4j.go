package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var (
	neo4jHost     = "10.10.11.141"
	neo4jPort     = "7687"
	neo4jUsername = "neo4j"
	neo4jPassword = "1q2w3e4r"
)

func GetNeo4jConnect() (driver neo4j.Driver, session neo4j.Session, err error) {
	if driver, err = neo4j.NewDriver(fmt.Sprintf("bolt://%s:%s", neo4jHost, neo4jPort), neo4j.BasicAuth(neo4jUsername, neo4jPassword, "")); err != nil {
		return
	}

	if session, err = driver.Session(neo4j.AccessModeWrite); err != nil {
		return
	}
	return
}

func Neo4jExec(session neo4j.Session, cypher string, params map[string]interface{}) (result neo4j.Result, err error) {

	result, err = session.Run(cypher, params)
	if err != nil {
		return // handle error
	}
	return
}
