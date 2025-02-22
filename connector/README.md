## AgroTech API

### Requirements 

- Golang
- Go VSCode Extension
- Docker rodando com banco de dados

### Instalating:

```bash
sudo apt-get install golang-go
```

### to compile it: 

This flags are used to compile the code to a linux binary file without any dependencies, it needs to be run in that way for alpine linux containers.

```bash
 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o connector.out
```