all: build

install_dep_tool:
	go get github.com/tools/godep

initialize_deps:
	go get -d -v ./...
	go get -d -v github.com/onsi/ginkgo
	go get -d -v github.com/onsi/gomega
	go get -d -v github.com/golang/lint/golint
	godep save

build:
	go build -a -installsuffix cgo -o fake_go_main .

.PHONY: build install_dep_tool initialize_deps
