PRJROOT = ../../

all: $(PRJROOT)/bin/WebUnlaaser
.PHONY:all

$(PRJROOT)/bin/WebUnlaaser: main.go
	mkdir -p bin
	go build -compiler gc -o $@ $<
