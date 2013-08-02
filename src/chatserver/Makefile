PRJROOT = ../../

all: $(PRJROOT)/bin/ChatServer
.PHONY:all

$(PRJROOT)/bin/ChatServer: main.go
	mkdir -p bin
	go build -compiler gc -o $@ $<
