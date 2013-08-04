package main

import (
	"time"
	"sync"
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
	p.QHead = 0;
	p.MsgList = make([]*ChatMessage, qsize)
	for i := range p.MsgList {
		p.MsgList[i] = new(ChatMessage)
	}
	p.cond = sync.NewCond(&p.mutex)

	return p
}

func (this *ChatBoard) PostMessage(msg string) {
	var head int

	this.mutex.Lock()
	{
		head = this.QHead
		this.QHead = (head + 1) % len(this.MsgList)
	}
	this.mutex.Unlock()
	// 假设消息狂多循环队扣圈了，会触发bug

	var pMsg *ChatMessage = this.MsgList[head]
	pMsg.Time = time.Now()
	pMsg.Text = msg

	this.cond.Broadcast()
}

func (this *ChatBoard) GetMessages(since int) (msgs []*ChatMessage, next int) {
	var head int

	this.mutex.Lock()
	{
		for this.QHead == since {
			this.cond.Wait()
		}
		head = this.QHead
	}
	this.mutex.Unlock()

	if (since <= head) {
		msgs = this.MsgList[since:head]
	} else {
		msgs = append(this.MsgList[since:], this.MsgList[:head]...)
	}
	next = head

	return
}
