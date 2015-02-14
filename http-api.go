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
		fmt.Fprintf(w, "Hello world")
	})
	http.Handle("/foo", Person{"Bob"})

	http.ListenAndServe(":8000", nil)
}
