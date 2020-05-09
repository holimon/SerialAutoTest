cd %~dp0
go build -ldflags "-s -w" -i -o ../bin/SerialTestServer.exe ../cmd/main.go