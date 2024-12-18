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