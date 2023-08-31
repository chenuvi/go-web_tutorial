package main

import "net/http"

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

type aboutHandler struct{}

func (h *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome !"))
}

func main() {

	helloH := helloHandler{}
	aboutH := aboutHandler{}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil, // Handler 为 nil 时，为  DefaultHTTPServeMux （路由)
	}
	http.Handle("/hello", &helloH)
	http.Handle("/about", &aboutH)

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home"))
	})
	http.HandleFunc("/welcome", welcome)
	server.ListenAndServe()

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, world!"))
	// })

	// http.ListenAndServe("localhost:8080", nil) // DefaultHTTPServeMux
}
