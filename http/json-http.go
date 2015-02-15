// Just return some stuff to get
package main

import (
	"encoding/json"
	"net/http"
)

func JsonHomeHandler(w http.ResponseWriter, r *http.Request) {
	t := map[string]int{"Apples": 5, "Oranges": 6}
	json.NewEncoder(w).Encode(t)
}
func main() {

	http.HandleFunc("/", JsonHomeHandler)

	http.ListenAndServe(":8000", nil)
}
