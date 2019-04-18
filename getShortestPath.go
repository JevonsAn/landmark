package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"io/ioutil"
	"log"
	"strings"
)

var landmarkFileName = "data/LG_task_IPs_50.ip_list"

func GetLandMarkIPs(landmarkFileName string) (ipList []string) {
	if contents, err := ioutil.ReadFile(landmarkFileName); err == nil {
		ipList = strings.Fields(string(contents))
		// fmt.Println(len(ipList))
	} else {
		fmt.Println(err)
	}
	return
}

func GetShortestLM(session neo4j.Session, ip string) (ipRoad [][]string, relationships [][]map[string]string, err error) {
	ipList := GetLandMarkIPs(landmarkFileName)

	results, err := GetShortestPath(session, ip, ipList)
	if err != nil {
		log.Fatal(err)
		return
	}

	idToIp := make(map[int64]string)

	for _, onepath := range results {
		ipPath := make([]string, 0)
		for _, node := range onepath.Nodes() {
			ip := node.Props()["ip"]
			if value, ok := ip.(string); ok {
				ipPath = append(ipPath, value)
				idToIp[node.Id()] = value
			} else {
				fmt.Println(node)
			}
		}
		ipRoad = append(ipRoad, ipPath)

		relationship := make([]map[string]string, 0)
		for _, rel := range onepath.Relationships() {
			rela := make(map[string]string)
			for key, value := range rel.Props() {
				if v, ok := value.(string); ok {
					rela[key] = v
				}
			}
			rela["in"] = idToIp[rel.StartId()]
			rela["out"] = idToIp[rel.EndId()]
			relationship = append(relationship, rela)
		}
		relationships = append(relationships, relationship)
	}
	return
}
