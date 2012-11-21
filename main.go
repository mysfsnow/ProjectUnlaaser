package main

import (
	"fmt"
	"net/http"
	"net/http/cgi"
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
	http.HandleFunc("/cgi/", handleCgi)
	err := http.ListenAndServe(localAddr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleMain (out http.ResponseWriter, request *http.Request) {
	u := (* request).URL

	log.Print("Requesting: ", u.Path)

	switch u.Path {
	case "/":
		handleHello(out, request)
	default:
		http.NotFound(out, request)
	}
}

func handleHello(out http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(out, "<html><head><meta charset='utf-8' /><title>UnLaas</title></head><body>不能Laas</body></html>")
}

func handleCgi(out http.ResponseWriter, request *http.Request) {
	u := (* request).URL
	
	var hdl cgi.Handler
	hdl.Path = "./" + u.Path
	hdl.Root = "/cgi/"
	hdl.Dir = "."
	hdl.ServeHTTP(out, request)
}