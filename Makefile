build-linux:
	GOOS=linux GOARCH=amd64 go build cmd/ddns_server.go