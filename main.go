package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type fbHandler string

func (fbh fbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) > 0 {
		num, err := strconv.Atoi(r.URL.Path)
		if err == nil {
			switch {
			case num%15 == 0:
				fmt.Fprintln(w, "Fizz Buzz")
			case num%3 == 0:
				fmt.Fprintln(w, "Fizz")
			case num%5 == 0:
				fmt.Fprintln(w, "Buzz")
			default:
				fmt.Fprintf(w, "%d\n", num)
			}
		} else {
			http.Error(w, "Innumid request.", http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Innumid request.", http.StatusBadRequest)
	}
}
func main() {
	var fbh fbHandler
	http.Handle("/fizzbuzz/", http.StripPrefix("/fizzbuzz/", fbh))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
