package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	// 返回路径，并对访问次数计数
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(rw, "URL.Path = %q\n", r.URL.Path)
	})

	// 展示访问次数
	http.HandleFunc("/count", func(rw http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(rw, "Count %d\n", count)
		mu.Unlock()
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
