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

	err := http.ListenAndServe(localAddr, MainHandler{})
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

type MainHandler struct {
}

func (MainHandler) ServeHTTP(out http.ResponseWriter, request *http.Request) {
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

