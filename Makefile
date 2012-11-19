all: bin/WebUnlaaser
.PHONY:all

bin/WebUnlaasern: main.go
	mkdir -p bin
	go build -compiler gc -o bin/WebUnlaaser main.go
