package main

import (
	"time"
	"net/http"
	"log"
	"os"
	"io/ioutil"
	"sync"
	"github.com/garyburd/go-websocket/websocket"
)


type ChatMessage struct {
	Time time.Time
	Text string
}

type ChatBoard struct {
	mutex	sync.Mutex
	cond    *sync.Cond
	MsgList []*ChatMessage
	QHead   int
}

func NewChatBoard(qsize int) *ChatBoard {
	p := new(ChatBoard)
	p.QHead = 1;
	p.MsgList = make([]*ChatMessage, qsize)
	for i := range p.MsgList {
		p.MsgList[i] = new(ChatMessage)
	}
	p.cond = sync.NewCond(&p.mutex)

	return p
}

func (this *ChatBoard) postMessage(msg string) {
	var head int

	this.mutex.Lock()
	{
		head = this.QHead
		this.QHead = head + 1
	}
	this.mutex.Unlock();
	// 假设消息狂多循环队扣圈了，会触发bug

	var pMsg *ChatMessage = this.MsgList[head]
	pMsg.Time = time.Now()
	pMsg.Text = msg

	this.cond.Broadcast();
}

func (this *ChatBoard) getMessages(since int) (msgs []*ChatMessage, next int) {
	var head int

	this.mutex.Lock();
	{
		for this.QHead == since {
			this.cond.Wait();
		}
		head = this.QHead;
	}
	this.mutex.Unlock();

	if (since <= head) {
		msgs = this.MsgList[since:head]
	} else {
		msgs = append(this.MsgList[since:], this.MsgList[:head]...)
	}
	next = head

	return
}

func getOptionLocalAddr(args []string) string {
	if len(args) == 2 {
		return args[1]
	} else {
		return ":8081"
	}
}

var board *ChatBoard

func main() {
	var localAddr string = getOptionLocalAddr(os.Args);

	http.HandleFunc("/chat", handleWs)
	http.HandleFunc("/", http.NotFound)

	board = NewChatBoard(4096)

	err := http.ListenAndServe(localAddr, nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
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
		handleWebSocket(ws)
	}
}

func handleWebSocket(ws *websocket.Conn) {

	go func () {
		for {
			since := board.QHead;
			recv_msg, next := board.getMessages(since)
			since = next

			for _, m := range recv_msg {
				send_msg := "[" + m.Time.Format("15:04:05") + "] " + m.Text
				ws.WriteMessage(websocket.OpText, []byte(send_msg));
			}
		}
	} ()

	for {
		opcode, reader, err := ws.NextReader()
		if err != nil { break }

		switch opcode {
		case websocket.OpText:
			msg, err1 := ioutil.ReadAll(reader)
			if err1 != nil { break }

			board.postMessage(string(msg))
		}
	}
}