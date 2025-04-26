.PHONY: all clean windows linux

all: windows linux

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/game-of-life-windows.exe

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/game-of-life-linux

clean:
	rm -rf bin/ 