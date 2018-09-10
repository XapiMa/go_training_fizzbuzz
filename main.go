package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fbf := func(w http.ResponseWriter, r *http.Request) {
		num, err := strconv.Atoi(r.URL.Path)
		if err != nil {
			http.Error(w, "Innumid request.", http.StatusBadRequest)
			return
		}
		switch {
		case num%15 == 0:
			fmt.Fprintln(w, "Fizz Buzz")
		case num%3 == 0:
			fmt.Fprintln(w, "Fizz")
		case num%5 == 0:
			fmt.Fprintln(w, "Buzz")
		default:
			fmt.Fprintln(w, num)
		}
	}
	http.Handle("/fizzbuzz/", http.StripPrefix("/fizzbuzz/", http.HandlerFunc(fbf)))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
