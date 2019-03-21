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

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))


	fmt.Printf("Starting on %s\n", name)
	fmt.Println("Listening on port 9000 :-)\n")
	http.ListenAndServe(":9000", nil)
}
