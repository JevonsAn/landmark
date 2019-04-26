package main

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
	"strings"
)

var (
	neo4jHost     = "10.10.11.141"
	neo4jPort     = "7687"
	neo4jUsername = "neo4j"
	neo4jPassword = "1q2w3e4r"
	nodeTableName = "node"
	edgeTableName = "edge"
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

func neo4jExec(session neo4j.Session, cypher string, params map[string]interface{}) (result neo4j.Result, err error) {

	result, err = session.Run(cypher, params)
	if err != nil {
		return // handle error
	}
	return
}

func getShortestPath(session neo4j.Session, ip string, ipList []string, n int) ([]neo4j.Path, error) {
	var result neo4j.Result
	var err error

	paths := make([]neo4j.Path, 0)

	cypher := fmt.Sprintf("MATCH (n:%s {ip: $ip}),(m:%s), p=shortestpath((n)-[e:%s *..20]-(m)) ", nodeTableName, nodeTableName, edgeTableName)
	cypher += "where m.ip in ['" + strings.Join(ipList, "', '") + "'] return p "
	cypher += "order by length(p) limit $n"

	params := map[string]interface{}{"ip": ip, "n": n}
	fmt.Println(cypher)

	_, err = session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		if result, err = tx.Run(cypher, params); err != nil {
			return nil, err
		}

		for result.Next() {
			path := result.Record().GetByIndex(0)
			if value, ok := path.(neo4j.Path); ok {
				paths = append(paths, value)
			} else {
				fmt.Println("onepath isn`t neo4j.path")
			}
		}
		return result.Consume()
	})
	return paths, err
}

func getNearestPathLength(session neo4j.Session, ip string, ipList []string) map[string]int {
	var result neo4j.Result
	var err error

	ipLengths := make(map[string]int, 0)

	cypher := fmt.Sprintf("MATCH (n:%s {ip: $ip}),(m:%s), p=shortestpath((n)-[e:%s *..20]-(m)) ", nodeTableName, nodeTableName, edgeTableName)
	cypher += "where m.ip in ['" + strings.Join(ipList, "', '") + "'] and length(p) > 0 return m.ip, length(p) "
	//cypher += "order by length(p)"

	params := map[string]interface{}{"ip": ip}
	//fmt.Println(cypher)

	_, err = session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		if result, err = tx.Run(cypher, params); err != nil {
			return nil, err
		}

		for result.Next() {
			mip := result.Record().GetByIndex(0)
			length := result.Record().GetByIndex(1)
			value1, ok1 := mip.(string)
			value2, ok2 := length.(int64)
			if ok1 && ok2 {
				if value2 > 0 {
					ipLengths[value1] = int(value2)
				}
			} else {
				fmt.Println("get length wrong")
			}
		}
		return result.Consume()
	})
	if err != nil {
		log.Fatal(err)
	}
	return ipLengths
}

func InsertNode(session neo4j.Session, ip string) error {
	var result neo4j.Result
	var err error

	cypher := fmt.Sprintf("MERGE (n:%s { ip: $ip }) RETURN count(n)", nodeTableName)
	params := map[string]interface{}{"ip": ip}

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		if result, err = tx.Run(cypher, params); err != nil {
			return nil, err
		}
		return result.Consume()
	})

	return err
}

func InsertEdge(session neo4j.Session, link []string) error {
	// link： 1.in 2.out 3.is_dest 4.star 5.delay 6.freq 7.ttl 8.monitor 9.time1 10. time2
	//{
	// "firstseen": "1534545761",
	// "lastseen": "1534545761",
	// "is_dest": "Y",
	// "delay": "0.000",
	// "star": "0",
	// "freq": "1",
	// "monitor": "198.86.53.6",
	// "ttl": "4"
	//}
	var result neo4j.Result
	var err error

	cypher := fmt.Sprintf("match (a:%s {ip: $src_ip}), (b:%s {ip: $dst_ip}) ", nodeTableName, nodeTableName)
	cypher += fmt.Sprintf("merge (a)-[e:%s]->(b) ", edgeTableName)
	cypher += "on match set e.firstseen = $firstseen, e.lastseen=$lastseen, e.is_dest=$is_dest, e.delay=$delay, e.star=$star, e.freq=$freq, e.monitor=$monitor, e.ttl=$ttl "
	cypher += "on create set e.firstseen = $firstseen, e.lastseen=$lastseen, e.is_dest=$is_dest, e.delay=$delay, e.star=$star, e.freq=$freq, e.monitor=$monitor, e.ttl=$ttl "
	cypher += "return count(e)"

	params := map[string]interface{}{
		"src_ip":    link[0],
		"dst_ip":    link[1],
		"is_dest":   link[2],
		"star":      link[3],
		"delay":     link[4],
		"freq":      link[5],
		"ttl":       link[6],
		"monitor":   link[7],
		"firstseen": link[8],
		"lastseen":  link[9],
	}

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		if result, err = tx.Run(cypher, params); err != nil {
			return nil, err
		}
		return result.Consume()
	})

	return err
}

// merge关系 CQL
//MATCH (charlie:Person { name: 'Charlie Sheen' }),(wallStreet:Movie { title: 'Wall Street' })
//MERGE (charlie)-[r:ACTED_IN]->(wallStreet)
//RETURN charlie.name, type(r), wallStreet.title

//match (a:ayh_node {ip:'1.1.1.1'}), (b:ayh_node {ip:'2.2.2.2'})
//merge (a)-[e:ayh_edge]->(b)
//on match set e.time = time
//on create set e.time = time
//return e

// 创建关系CQL
//MATCH (a:Person),(b:Movie)
//WHERE a.name = 'Tom Hanks' AND b.title = 'Forrest Gump'
//CREATE (a)-[r:ACTED_IN { roles:['Forrest'] }]->(b)
//RETURN r;

//match (a:ayh_node {ip:'1.1.1.1'}), (b:ayh_node {ip:'2.2.2.2'})
//create (a)-[e:ayh_edge {time:'001'}]->(b)
//return e

//最短路径
// MATCH (n:node {ip:'8.8.8.8'}),(m:node), p=shortestpath((n)-[e:edge *..20]-(m))
// where m.ip in ['114.114.114.114', '98.182.1.86'] return p
// order by length(p) limit 3
