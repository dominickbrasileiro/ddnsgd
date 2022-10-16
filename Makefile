default: build

install_deps:
	go get -v ./...

build: install_deps
	go build -o ./bin/ddnsgd .

clean:
	rm -rf ./bin
