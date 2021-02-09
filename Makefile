BIN=task

build:
	go build -o $(BIN)

clean:
	-rm -f $(BIN)
