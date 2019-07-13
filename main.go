package main

import (
	"fmt"
	"net/http"
	"os"
)
var total int
func main() {
        total = 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Printf("value of environment variable %v: %v\n", "ECHO_VAR", os.Getenv("ECHO_VAR"))
                total++
		r.Write(w)
                if total > 9 {
                    os.Exit(2)
                }
	})
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}


