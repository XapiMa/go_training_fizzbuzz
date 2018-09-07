package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/fizzbuzz/", func(w http.ResponseWriter, r *http.Request) {
		flag := false
		var num int
		if len(r.URL.Path) > 1 {
			param := r.URL.Path[len("/fizzbuzz/"):]
			val, err := strconv.Atoi(param)
			if err != nil {
				flag = false
			} else {
				flag = true
				num = val
			}
		}

		switch {
		case flag == false:
			http.Error(w, "Invalid request.", http.StatusBadRequest)
		case num%15 == 0:
			fmt.Fprintf(w, "Fizz Buzz\n")
		case num%3 == 0:
			fmt.Fprintf(w, "Fizz\n")
		case num%5 == 0:
			fmt.Fprintf(w, "Buzz\n")
		default:
			fmt.Fprintf(w, "%d\n", num)
		}
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
