file_name := go-swiftclient
out_dir := build

all: clean linux mac windows

clean:
	rm -rf $(out_dir)

linux:
	GOOS=linux GOARCH=arm64 go build -o $(out_dir)/$(file_name)-linux-arm64
	GOOS=linux GOARCH=amd64 go build -o $(out_dir)/$(file_name)-linux-amd64

mac:
	GOOS=darwin GOARCH=amd64 go build -o $(out_dir)/$(file_name)-darwin-amd64
	GOOS=darwin GOARCH=arm64 go build -o $(out_dir)/$(file_name)-darwin-arm64

windows:
	GOOS=windows GOARCH=amd64 go build -o $(out_dir)/$(file_name)-windows-amd64.exe