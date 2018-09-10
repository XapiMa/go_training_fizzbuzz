package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fbf := func(w http.ResponseWriter, r *http.Request) {
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
	http.Handle("/fizzbuzz/", http.StripPrefix("/fizzbuzz/", http.HandlerFunc(fbf)))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
