build_for_windows:
	GOOS=windows GOARCH=amd64 go build -o ./build/ir-remocon-win.exe

build_for_mac:
	GOOS=darwin GOARCH=amd64 go build -o ./build/ir-remocon-mac

build_for_linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/ir-remocon-linux
