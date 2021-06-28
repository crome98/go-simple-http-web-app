package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	a := 1000000000.0
	result := 0.0
	for i:=0; i < 10; i++ {
		for j := 0; j < 600000; j++ {
			result = math.Sqrt(a)
		}
	}
	log.Printf("result %f", result)
	http.ServeFile(w, req, "./20201229_Weekly_Report_LEHOANGPHUC.docx")
	//fmt.Fprintf(w, "Hello from server %f", result)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.ListenAndServe(":8080", nil)
}

