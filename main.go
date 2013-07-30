package main

import (
	"fmt" // 按格式输出字符串的库
	"net/http"
	"net/http/cgi"
	"log"
	"os"
	"io"
	"io/ioutil"
	"github.com/garyburd/go-websocket/websocket"
)

/** 主函数 */
func main() {
	args := os.Args;
	var localAddr string
	if len(args) == 2 {
	/* 如果使用 一个参数 */
		localAddr = args[1]
	} else {
	/* 不使用参数运行，默认使用 80 端口 */
		localAddr = ":80"
	}

	// 对于cgi路径下的网址，调用相应名称的cgi程序
	http.HandleFunc("/cgi/", handleCgi)
	// 对于stc路径下的网址，调用相应的静态文件
	http.HandleFunc("/stc/", handleFile)
	// WebSocket
	http.HandleFunc("/ws/echo", handleWs);
	// 对于其他路径下的网址
	http.HandleFunc("/", handleMain)

	// 服务器启动
	err := http.ListenAndServe(localAddr, nil)

	// 当服务器出错退出时
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleMain (out http.ResponseWriter, request *http.Request) {
	u := (*request).URL

	log.Print("Requesting: ", u.Path)

	// 检查所访问的网址
	switch u.Path {
	// 根目录
	case "/":
		handleHello(out, request)
	default:
		http.NotFound(out, request)
	}
}

/** 输出Hello的Handler */
func handleHello(out http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(out, "<html><head><meta charset='utf-8' /><title>UnLaas</title></head><body>不能Laas</body></html>")
}

/** 调用CGI程序的Handler */
func handleCgi(out http.ResponseWriter, request *http.Request) {
	u := (* request).URL

	var hdl cgi.Handler
	hdl.Path = "./" + u.Path
	hdl.Root = "/cgi/"
	hdl.Dir = "."
	hdl.ServeHTTP(out, request)
}

/** 静态文件的Handler */
func handleFile(out http.ResponseWriter, request *http.Request) {
	u := (* request).URL
	
	log.Print("Requesting Static File: ", u.Path)
	if u.Path[:5] == "/stc/" {
		http.ServeFile(out, request, u.Path[1:])
	} else {
		http.NotFound(out, request)
	}
}

func handleWs(out http.ResponseWriter, request *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}

	ws, err := websocket.Upgrade(out, request.Header, nil, 40966, 4096)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
		return
	}

	handleWebSocket(ws)
}
func handleWebSocket(ws *websocket.Conn) {
	for {
		opcode, reader, err := ws.NextReader()
		if err != nil { break }

		switch op {
		case websocket.OpText:
			msg, err := ioutil.ReadAll(reader)
			if err != nil { break }

			msgText := string(msg)
			replyText := "你说: {" + msgText + "}."

			ws.WriteMessage(OpText, byte[](replyText))
		}
	}

}