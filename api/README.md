## AgroTech API

### Requistos 

- Golang
- Go VSCode Extension
- Docker rodando com banco de dados

### Instalação

```bash
# Instalar Golang
sudo apt-get install golang-go

# Certifique-se de que o GOPATH está configurado corretamente
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# Inicializar o SWAG
swag init

# Executar o projeto
go run main.go
```
# to download it: 

```
sudo apt install golang-go
```

# to compile it: 
```
 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o apiconnector.out
```