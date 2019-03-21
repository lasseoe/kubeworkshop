package main

import (
	"fmt"
	"time"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "-=- Kubernetes -=- 21 march 2019")
		fmt.Fprintf(w, "%s", time.Now())
	})

	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":9000", nil)
}
