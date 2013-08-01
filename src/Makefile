all: bin/WebUnlaaser
.PHONY:all

PRJROOT = ../

bin/WebUnlaaser: main.go
	mkdir -p bin
	go build -compiler gc -o $(PRJROOT)/bin/WebUnlaaser main.go
