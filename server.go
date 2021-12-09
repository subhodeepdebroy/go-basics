package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	pairs "./go-basics"
)

type Response struct {
	Success bool
	Error   string
	Result  [2]int
}

func getPair(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["number"]
	var sumOfTwo int
	if ok {
		if i, err := strconv.Atoi(keys[0]); err == nil {
			sumOfTwo = i
		}
	} else {
		http.Error(w, "Somthing went wrong", http.StatusBadRequest)
	}
	inputArray := []int{7, 8, 7, 9, 1, 7, 1, 4, 1, 3}
	pairArray, err := pairs.FindPair(inputArray, sumOfTwo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := Response{true, "non", pairArray}
	fmt.Print(resp)
	js, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func handleRequests() {
	http.HandleFunc("/pair", getPair)
	log.Fatal(http.ListenAndServe(":3005", nil))
}

func main() {
	handleRequests()
}
