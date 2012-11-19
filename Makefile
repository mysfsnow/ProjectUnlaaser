all: bin/WebUnlaaser
.PHONY:all

bin/WebUnlaaser: main.go
	mkdir -p bin
	go build -compiler gc -o bin/WebUnlaaser main.go
