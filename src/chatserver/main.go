package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"io/ioutil"
	"github.com/garyburd/go-websocket/websocket"
)

func getOptionLocalAddr(args []string) string {
	if len(args) == 2 {
		return args[1]
	} 
	return ":8081"
}

var board *ChatBoard

func main() {
	var localAddr string = getOptionLocalAddr(os.Args);

	http.HandleFunc("/chat", handleWs)
	http.HandleFunc("/", handleFile)

	board = NewChatBoard(4096)

	err := http.ListenAndServe(localAddr, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleFile(out http.ResponseWriter, request *http.Request) {
	http.ServeFile(out, request, "stc/chat.htm")
}

func handleWs(out http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(out, "Method not allowed", 405)
		return
	}

	ws, err := websocket.Upgrade(out, request.Header, nil, 4096, 4096)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(out, "Not a websocket handshake", 400)
	} else if err != nil {
		log.Println(err)
	} else {
		handleChat(ws)
	}
}

func handleChat(ws *websocket.Conn) {

	poster := func () int {
		for {
			since := board.QHead;
			recv_msg, next := board.GetMessages(since)
			since = next

			for _, m := range recv_msg {
				timems := m.Time.Unix() * 1000 +
						  int64(m.Time.Nanosecond() / 1000000);

				timeStr := fmt.Sprintf("%016x", timems)
				send_msg := timeStr + m.Text
				ws.WriteMessage(websocket.OpText, []byte(send_msg));
			}
		}

		return 0
	}

	fetcher := func () int {
		for {
			opcode, reader, err := ws.NextReader()
			if err != nil { break }

			switch opcode {
				case websocket.OpText: {
					msg, err := ioutil.ReadAll(reader)
					if err != nil { break }

					board.PostMessage(string(msg))
				}
				case websocket.OpBinary: {
					msg, err := ioutil.ReadAll(reader)
					if err != nil { break }

					msgStr := fmt.Sprintf("%v", msg)
					board.PostMessage(string(msgStr))
				}
			}
		}

		return 0
	}

	sync_end := make(chan int)
	
	go func () { sync_end <- poster() } ()
	fetcher()

	<- sync_end
}