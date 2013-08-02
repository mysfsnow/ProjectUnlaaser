PRJROOT = ../../

all: $(PRJROOT)/bin/WebUnlaaser
.PHONY:all

$(PRJROOT)/bin/WebUnlaaser: main.go
	go build -compiler gc -o $@ $<
