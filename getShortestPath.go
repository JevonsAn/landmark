package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"io/ioutil"
	"log"
	"strings"
)

//var landmarkFileName = "data/LG_task_IPs_50.ip_list"

func GetLandMarkIPs(landmarkFileName string) (ipList []string) {
	if contents, err := ioutil.ReadFile(landmarkFileName); err == nil {
		ipList = strings.Fields(string(contents))
		// fmt.Println(len(ipList))
	} else {
		fmt.Println(err)
	}
	return
}

type shortestPath struct {
	ipRoad        []string
	relationships []map[string]string
}
type ipAndjump struct {
	ip   string
	jump int
}

func getShortestLM(session neo4j.Session, ip string, ipList []string, n int, ch chan shortestPath) {
	//	(ipRoad [][]string, relationships [][]map[string]string, err error)

	results, err := getShortestPath(session, ip, ipList, n)
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
		//ipRoad = append(ipRoad, ipPath)

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
		//relationships = append(relationships, relationship)
		fmt.Println(ipPath, relationship)
		ch <- shortestPath{ipPath, relationship}
	}
	return
}

func getTopnLM(ip string, n int) []ipAndjump {
	allIpList := make([]string, 0)
	for k := range LMinfo {
		allIpList = append(allIpList, k)
	}
	topIps := make([]ipAndjump, 0)
	ipToLength := getNearestPathLength(session, ip, allIpList)
	//fmt.Println(ipToLength)
	for _, iplist := range typeToIp {
		//fmt.Println(typenum, iplist)
		length := len(iplist)
		//fmt.Println(iplist, length)
		t := make([]ipAndjump, length)
		for i, ip := range iplist {
			jump, ok := ipToLength[ip]
			if ok == false {
				jump = 21
			}
			t[i] = ipAndjump{ip, jump}
		}
		//fmt.Println(t)
		for i := 0; i < length-1; i++ {
			mnIndex := i
			for j := i + 1; j < length; j++ {
				if t[j].jump < t[mnIndex].jump {
					mnIndex = j
				}
			}
			temp := t[i]
			t[i] = t[mnIndex]
			t[mnIndex] = temp
		}
		for i := 0; i < n && i < length; i++ {
			topIps = append(topIps, t[i])
		}
		//fmt.Println(topIps)
	}
	return topIps
}

type LMTableLine struct {
	Ip       string
	JumpNum  int
	Typename string
	Name     string
	Country  string
	Province string
	City     string
}

func GetLMTable(ip string, n int) (table []LMTableLine) {
	ipJumps := getTopnLM(ip, n)
	for _, v := range ipJumps {
		LMIP := v.ip
		table = append(table, LMTableLine{
			LMIP,
			v.jump,
			idToType[LMinfo[LMIP].typenum],
			LMinfo[LMIP].name,
			LMinfo[LMIP].country,
			LMinfo[LMIP].province,
			LMinfo[LMIP].city,
		})
	}
	return
}
