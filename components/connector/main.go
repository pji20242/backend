package main

import (
    "database/sql"
    "fmt"
    "log"
    "strings"
    "time"
    "strconv" // Adicione esta linha

    MQTT "github.com/eclipse/paho.mqtt.golang"
    _ "github.com/go-sql-driver/mysql"
)

// Configuração de conexão com o banco de dados MySQL
var db *sql.DB

func initDB() {
    var err error
    db, err = sql.Open("mysql", "connectoruser:connectorpasswrd@tcp(database:3306)/pjiot")
    if err != nil {
        log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
    }

    // Testa a conexão com o banco
    if err = db.Ping(); err != nil {
        log.Fatalf("Erro ao pingar o banco de dados: %v", err)
    }
    fmt.Println("Conectado ao banco de dados MySQL!")
}

func main() {
    // Inicializar a conexão com o banco de dados
    initDB()
    defer db.Close()

    // Configuração do cliente MQTT
    opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt-broker:1883")
    opts.SetClientID("go_mqtt_client")
    opts.OnConnect = func(c MQTT.Client) {
        fmt.Println("Conectado ao broker MQTT!")
    }
    opts.OnConnectionLost = func(c MQTT.Client, err error) {
        fmt.Printf("Conexão perdida: %v\n", err)
    }

    // Criação do cliente
    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatalf("Erro ao conectar ao broker: %v", token.Error())
    }

    // Subscreve a um tópico
    topic := "pji3"
    if token := client.Subscribe(topic, 1, func(client MQTT.Client, msg MQTT.Message) {
        processMessage(string(msg.Payload()))
    }); token.Wait() && token.Error() != nil {
        log.Fatalf("Erro ao subscrever ao tópico: %v", token.Error())
    }

    // Manter o programa em execução
    fmt.Println("Aguardando mensagens...")
    select {}
}

// processMessage processa a mensagem recebida e insere os dados no banco de dados
func processMessage(message string) {
    parts := strings.Split(message, "%")
    if len(parts) < 2 {
        fmt.Println("Formato da mensagem inválido.")
        return
    }

    uuid := parts[0]
    fmt.Println("UUID:", uuid)

    // Iterar sobre os parâmetros e extrair os dados
    for _, param := range parts[1:] {
        if param == "" {
            continue
        }

        if strings.HasSuffix(param, "@") {
            param = strings.TrimSuffix(param, "@")
        }

        keyValue := strings.Split(param, "=")
        if len(keyValue) != 2 {
            fmt.Println("Formato de parâmetro inválido:", param)
            continue
        }

        // Obter o tipo de sensor e o valor
        tipoSensor := keyValue[0]
        valor := keyValue[1]

        // Determina o ID do sensor com base no tipo
        idSensor := getSensorID(tipoSensor, uuid)

        // Insere os dados na tabela `dados`
        inserirDados(uuid, idSensor, valor)
    }
}

// getSensorID obtém o ID do sensor com base no tipo de sensor e no UUID do dispositivo
func getSensorID(tipoSensor, uuid string) int {
    var idSensor int
    err := db.QueryRow("SELECT idSensor FROM sensor WHERE tipo = ? AND uuid = ?", tipoSensor, uuid).Scan(&idSensor)
    if err != nil {
        log.Printf("Erro ao obter o idSensor para tipo %s e uuid %s: %v", tipoSensor, uuid, err)
    }
    return idSensor
}

// inserirDados insere os dados na tabela `dados`
func inserirDados(uuid string, idSensor int, valor string) {
    // Converte o valor para float
    valorFloat, err := strconv.ParseFloat(valor, 64)
    if err != nil {
        log.Printf("Erro ao converter o valor '%s' para float: %v", valor, err)
        return
    }

    // Executa a inserção no banco de dados
    _, err = db.Exec("INSERT INTO dados (ts, uuid, idSensor, valor) VALUES (?, ?, ?, ?)",
        time.Now(), uuid, idSensor, valorFloat)
    if err != nil {
        log.Printf("Erro ao inserir dados no banco de dados: %v", err)
    } else {
        fmt.Printf("Dados inseridos: UUID=%s, idSensor=%d, Valor=%.2f\n", uuid, idSensor, valorFloat)
    }
}
