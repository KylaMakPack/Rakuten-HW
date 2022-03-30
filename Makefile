BINARY_NAME=main.out
 
all: build test
 
build:
	go build -o ${BINARY_NAME} main.go
 
run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}