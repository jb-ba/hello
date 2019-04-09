package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", homeHandler)
	http.HandleFunc("/hello/hello", helloHandler)
	http.HandleFunc("/hello/info", infoHandler)
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprint(w, "This is your path\n")
		fmt.Fprintln(w, r.URL.Path)
		return
	}
	fmt.Fprint(w, "welcome home\n")
}
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "v3: welcome to hello\n")
}
func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "v3\n")
}
