// Just return some stuff to get
package main

import (
	"fmt"
	"net/http"
)

type Person struct {
	name string
}

func (p Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dude's called : %s", p.name)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("cache-control", "public")
		fmt.Printf("Method: %s\nURL: %s\nHeader: %s\nReqURI: %s \n", r.Method, r.URL, r.Header, r.RequestURI)

	})
	http.Handle("/foo", Person{"Bob"})

	http.ListenAndServe(":8000", nil)
}
