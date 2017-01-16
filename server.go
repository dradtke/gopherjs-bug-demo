package main

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		panic(err)
	}
}
