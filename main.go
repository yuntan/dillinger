package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	port = 8070
)

func main() {
	fmt.Println("===== dillinger (powerd by go) =====")

	startServer()
}

func startServer() {
	// files under ./public
	http.Handle("/css/", http.FileServer(http.Dir("./public")))
	http.Handle("/files/", http.FileServer(http.Dir("./public")))
	http.Handle("/ico/", http.FileServer(http.Dir("./public")))
	http.Handle("/img/", http.FileServer(http.Dir("./public")))
	http.Handle("/js/", http.FileServer(http.Dir("./public")))

	// editor
	http.HandleFunc("/", handler)

	log.Println("Server Started")
	log.Printf("Waiting connection on localhost:%d\n", port)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func handler(res http.ResponseWriter, req *http.Request) {
	editor, _ := NewEditor("")
	err := editor.Render(res)
	if err != nil {
		log.Fatalln(err)
	}
}
