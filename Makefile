bindata.go: templates/index.html
	go-bindata -o bindata.go ./templates/...

gophercon: main.go bindata.go
	go build -race -o gophercon .
