## AgroTech API

### Requirements 

- Golang
- Go VSCode Extension
- Docker rodando com banco de dados

### Instalating:

```bash
sudo apt-get install golang-go

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
go install github.com/swaggo/swag/cmd/swag@latest

# Inicialize the project
swag init
```

go get google.golang.org/api/idtoken

### to compile it: 

This flags are used to compile the code to a linux binary file without any dependencies, it needs to be run in that way for alpine linux containers.

```bash
 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o apiconnector.out
```

## env example:

```env
CLIENT_ID=<client_id>
CLIENT_SECRET=<client_secret>
```