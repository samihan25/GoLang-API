package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage")
	fmt.Println("Endpoint hit: homepage")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the about page")
	fmt.Println("Endpoint hit: about")
}

func add(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the add page")
	fmt.Println("Endpoint hit: add")
	fmt.Fprintln(w, "")

	cur1 := r.URL.Query().Get("cur1")
	cur1_int, _ := strconv.Atoi(cur1)
	fmt.Fprintln(w, "cur1= "+cur1)
	fmt.Fprintln(w, "cur1_int= %d", cur1_int)
	fmt.Fprintln(w, "")

	cur2 := r.URL.Query().Get("cur2")
	cur2_int, _ := strconv.Atoi(cur2)
	fmt.Fprintln(w, "cur2= "+cur2)
	fmt.Fprintln(w, "cur2_int= %d", cur2_int)
	fmt.Fprintln(w, "")

	addition := cur1_int + cur2_int
	fmt.Fprintln(w, "addition= %d", addition)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", about)
	http.HandleFunc("/add", add)

	fmt.Println("Listening on http://localhost:10000/")
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
