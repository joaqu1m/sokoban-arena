GOOS=windows
GOARCH=amd64
CC=x86_64-w64-mingw32-gcc

default:
	CGO_ENABLED=1 CC=$(CC) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-s -w" -o sokoban.exe ./cmd/game

prod:
	CGO_ENABLED=1 CC=$(CC) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-H=windowsgui -s -w" -o sokoban.exe ./cmd/game
