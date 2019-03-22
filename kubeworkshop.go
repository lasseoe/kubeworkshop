package main

import (
	"fmt"
	"time"
	"net/http"
	"os"
)

func main() {

	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}


	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "-=- Kubernetes workshop @ Google DK -=-\n")
		fmt.Fprintf(w, " -- time: %s\n", time.Now())
		fmt.Fprintf(w, " -- hostname: %s", name)
	})

	fmt.Printf("Listening on %s:9000\n", name)
	http.ListenAndServe(":9000", nil)
}
