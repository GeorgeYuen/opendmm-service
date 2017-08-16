package main

import (
	"log"
	"net/http"
	"encoding/json"
	"time"
	"github.com/junzh0u/opendmm"
	"flag"
)

func searchDmm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uri := r.RequestURI
	w.Write([]byte(search(uri[len(*pattern):])))
}

var pattern  *string

func search(query string) (result string) {
	log.Println("start search " + query)
	match := opendmm.Search(query, nil)
	select {
	case meta, ok := <-match:
		if ok {
			metaJson, _ := json.MarshalIndent(meta, "", "  ")
			return string(metaJson)
		} else {
			log.Println("Not found")
		}
	case <-time.After(5 * time.Second):
		log.Println("Time out")
	}
	return "Not found"
}

func main() {
	pattern = flag.String("pattern", "/av/", "path to http cache")
	httpPort := flag.String("port", ":9090", "path to http cache")
	flag.Parse()
	log.Println("服务器启动")
	http.HandleFunc(*pattern, searchDmm)
	err := http.ListenAndServe(*httpPort, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
