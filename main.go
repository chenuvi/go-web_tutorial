package main

import (
	// "log"
	"fmt"
	"net/http"
)

func main() {
	// 4. 
	server := http.Server{
        Addr: ":8080",
	}
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024)
		fmt.Println(w, r.FormValue("first_name"))
		fmt.Println(w, r.PostFormValue("first_name"))
	})

	server.ListenAndServe()

	// 2. 请求参数
	// http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
    //     url := r.URL
	// 	query := url.Query()

	// 	id := query["id"]
	// 	log.Println(id)
	// 	name := query.Get("name")
	// 	log.Println(name)
    // })

	// 1. 文件服务器

	// http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot")))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "wwwroot" + r.URL.Path) 
	// })


	// http.ListenAndServe(":8080", nil)
}
