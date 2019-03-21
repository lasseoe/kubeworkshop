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
		fmt.Fprintf(w, "-=- Kubernetes -=- 21 march 2019")
		fmt.Fprintf(w, "%s\n", time.Now())
		fmt.Fprintf(w, " -- hostname: %s", name)
	})

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":9000", nil)
}
