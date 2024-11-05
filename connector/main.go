package main

import (
    "fmt"
    "log"
    "strings"
    "time" // Importar o pacote time

    MQTT "github.com/eclipse/paho.mqtt.golang"
    influxdb "github.com/influxdata/influxdb1-client/v2"
)

var influxClient influxdb.Client

func main() {
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

    // Configuração do InfluxDB
    var err error
    
    influxClient, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
        Addr:     "http://influxdb:8086", // ou o endereço do seu servidor InfluxDB
        Username: "myuser",                // nome do usuário padrão
        Password: "myuser_password",        // senha do usuário padrão
    })
    if err != nil {
        log.Fatalf("Erro ao conectar ao InfluxDB: %v", err)
    }
    defer influxClient.Close()

    // Subscreve a um tópico
    topic := "pji3"
    if token := client.Subscribe(topic, 1, func(client MQTT.Client, msg MQTT.Message) {
        processMessage(influxClient, string(msg.Payload())) // Passa influxClient aqui
    }); token.Wait() && token.Error() != nil {
        log.Fatalf("Erro ao subscrever ao tópico: %v", token.Error())
    }

    // Manter o programa em execução
    fmt.Println("Aguardando mensagens...")
    select {}
}

// processMessage processa a mensagem recebida e insere os dados no InfluxDB
func processMessage(influxClient influxdb.Client, message string) {
    // Separar o UUID da mensagem dos parâmetros
    parts := strings.Split(message, "%")
    if len(parts) < 2 {
        fmt.Println("Formato da mensagem inválido.")
        return
    }

    // O primeiro elemento é o UUID
    uuid := parts[0]
    fmt.Println("UUID:", uuid)

    // Criar um mapa de campos
    fields := map[string]interface{}{}

    // Iterar sobre os parâmetros e adicionar valores ao mapa de campos
    for _, param := range parts[1:] {
        if param == "" {
            continue
        }

        // Remover o delimitador "@" do final, se presente
        if strings.HasSuffix(param, "@") {
            param = strings.TrimSuffix(param, "@")
        }

        keyValue := strings.Split(param, "=")
        if len(keyValue) != 2 {
            fmt.Println("Formato de parâmetro inválido:", param)
            continue
        }

        key := keyValue[0]
        value := keyValue[1]

        // Adicionar os valores ao mapa de campos usando switch
        switch key {
        case "T":
            fields["temperature"] = value
        case "P":
            fields["pressure"] = value
        case "L":
            fields["luminosity"] = value
        case "U":
            fields["humidity"] = value
        case "Vol":
            fields["voltage"] = value
        case "Amp":
            fields["current"] = value
        default:
            fmt.Printf("Chave desconhecida: %s\n", key)
        }
    }

    // Verificar se há campos a serem inseridos
    if len(fields) == 0 {
        fmt.Println("Nenhum campo para inserir no InfluxDB.")
        return
    }

    // Criar um ponto para o InfluxDB
    point, err := influxdb.NewPoint("sensor_data", map[string]string{"uuid": uuid}, fields, time.Now())
    if err != nil {
        log.Printf("Erro ao criar ponto: %v", err)
        return
    }

    // Criar BatchPoints
    bp, err := influxdb.NewBatchPoints(influxdb.BatchPointsConfig{
        Database:  "mydb",
        Precision: "s",
    })
    if err != nil {
        log.Printf("Erro ao criar batch points: %v", err)
        return
    }

    // Adicionar o ponto ao BatchPoints
    bp.AddPoint(point)

    // Inserir o ponto no InfluxDB
    if err := influxClient.Write(bp); err != nil {
        log.Printf("Erro ao escrever no InfluxDB: %v", err)
        return
    }

    fmt.Println("Dados inseridos no InfluxDB com sucesso!")
}