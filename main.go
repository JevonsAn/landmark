package main

import (
	"encoding/json"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
	"net/http"
	"strconv"
)

var (
	driver   neo4j.Driver
	session  neo4j.Session
	idToType map[int]string
	LMinfo   map[string]LandMartInfo
	typeToIp map[int][]string
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	//fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	//fmt.Println("path", r.URL.Path)
	//fmt.Println("scheme", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的 }
}

func nearestLMHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "[]")
		return
	}
	if len(r.Form["ip"]) <= 0 || len(r.Form["n"]) <= 0 {
		fmt.Fprintf(w, "[]")
		return
	}
	ip := r.Form["ip"][0]
	n, _ := strconv.Atoi(r.Form["n"][0])
	//fmt.Println("ip, n:", ip, n)
	LMTable := GetLMTable(ip, n)
	//fmt.Println(LMTable)
	jsonResult, err := json.Marshal(LMTable)
	//fmt.Println(jsonResult)
	if err != nil {
		fmt.Fprintf(w, "[]")
		log.Println(err)
		return
	}
	fmt.Fprintf(w, string(jsonResult))

	//fmt.Println(GetLMTable("8.8.8.8",3))

	//ipRoad, relations, err := getShortestLM(session, "8.8.8.8")
}

func initDatabase() (err error) {
	driver, session, err = GetNeo4jConnect()
	if err != nil {
		log.Fatal(err)
		return
	}
	idToType, err = GetLandmarkType()
	if err != nil {
		log.Fatal(err)
		return
	}
	LMinfo, typeToIp, err = GetLandmarkData()
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func shutdown() (err error) {
	if session != nil {
		err = session.Close()
	}
	if driver != nil {
		err = driver.Close()
	}
	return
}

func main() {
	var err error
	err = initDatabase()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("初始化成功")
	defer shutdown()

	// 设置访问的路由
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/nearLM", nearestLMHandler)

	err = http.ListenAndServe(":8888", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
