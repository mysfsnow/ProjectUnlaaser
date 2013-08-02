PRJROOT = ../../

all: $(PRJROOT)/bin/ChatServer
.PHONY:all

$(PRJROOT)/bin/ChatServer: main.go
	go build -compiler gc -o $@ $<
