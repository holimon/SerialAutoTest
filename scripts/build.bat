cd %~dp0
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-s -w" -i -o ../bin/SerialTestServer.exe ../cmd/main.go
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-s -w" -i -o ../bin/SerialTestServer ../cmd/main.go