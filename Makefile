all: bin/WebUnlaaser
.PHONY:all

bin/WebUnlaaser: main.go
	go build -compiler gc -o bin/WebUnlaaser main.go
