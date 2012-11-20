package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

func main() {
	args := os.Args;
	var localAddr string
	if len(args) == 2 {
		localAddr = args[1]
	} else {
		localAddr = ":80"
	}

	http.HandleFunc("/", handleMain)
	err := http.ListenAndServe(localAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleMain(out http.ResponseWriter, request *http.Request) {
	u := (* request).URL;

	log.Print("Requesting: ", u.Path)

	if u.Path == "/" {
		handleHello(out, request)
	} else {
		http.NotFound(out, request)
	}
}

func handleHello(out http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(out, "<html><head><meta charset='utf-8' /><title>UnLaas</title></head><body>不能Laas</body></html>")
}

