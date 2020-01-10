set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -a -installsuffix cgo -ldflags "-w" -i -o nebula main.go plugin.go