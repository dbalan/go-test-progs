// A simple math api using gorilla

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type op func(int, int) int

var opMap = map[string]op{
	"add": add,
	"sub": sub,
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler)

	pr := router.Methods("POST").Subrouter()
	pr.HandleFunc("/{operation}/", opHandler)

	http.Handle("/", router)
	http.ListenAndServe(":8000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there!\n")
}

func opHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	opr := vars["operation"]
	f := opMap[opr]
	// Be a cowboy! :-P
	// undefined = 0
	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))
	fmt.Fprintf(w, "Here ya go: %d\n", f(x, y))
}

func add(x int, y int) int {
	z := x + y
	return z
}

func sub(x int, y int) int {
	z := x - y
	return z
}
