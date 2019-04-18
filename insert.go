package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"io/ioutil"
	"log"
	"strings"
)

var (
	linkFileName = "data/2019-03-30-12_45_00.links"
)

func GetLinks(fileName string) (lines []string) {
	if contents, err := ioutil.ReadFile(fileName); err == nil {
		lines = strings.FieldsFunc(string(contents), func(r rune) bool {
			if r == '\n' {
				return true
			}
			return false
		})
	} else {
		fmt.Println(err)
	}
	return
}

func InsertNeo4j(session neo4j.Session, lines []string) (err error) {
	for _, line := range lines {
		linesplit := strings.Fields(line)
		err = InsertNode(session, linesplit[0])
		if err != nil {
			log.Fatal(err)
			return
		}
		err = InsertNode(session, linesplit[1])
		if err != nil {
			log.Fatal(err)
			return
		}
		err = InsertEdge(session, linesplit)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	return
}

func InsertTraceResult() {
	driver, session, err := GetNeo4jConnect()
	if err != nil {
		println(err)
		defer driver.Close()
		defer session.Close()
	}

	links := GetLinks(linkFileName)
	err = InsertNeo4j(session, links)

	if err != nil {
		println(err)
	}
}
