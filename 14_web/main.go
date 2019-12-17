package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>About</h1>")
}

// Go is very good to build rest api, micro services
// You can also use template language, template engine to build them.
func main()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Starting...")
	http.ListenAndServe(":3000", nil)
}
