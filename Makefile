all: build

build:
	go build -a -installsuffix cgo -o fake_go_main .

.PHONY: build
