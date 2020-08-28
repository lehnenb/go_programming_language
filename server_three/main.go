// Server2 is a minimal "echo" and counter server
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request - %s: %s\n", r.URL.Path, time.Now().Format(time.RFC850))
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
