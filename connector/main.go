package main

import (
    "fmt"
    "log"
    "strings"

    MQTT "github.com/eclipse/paho.mqtt.golang"
)

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

// processMessage processa a mensagem recebida e imprime os dados formatados
func processMessage(message string) {
    // Separar o UUID da mensagem dos parâmetros
    parts := strings.Split(message, "%")
    if len(parts) < 2 {
        fmt.Println("Formato da mensagem inválido.")
        return
    }

    // O primeiro elemento é o UUID
    uuid := parts[0]
    fmt.Println("UUID:", uuid)

    // Iterar sobre os parâmetros
    for _, param := range parts[1:] {
        // Se o parâmetro não estiver vazio
        if param == "" {
            continue
        }

        // Separar chave e valor
        keyValue := strings.Split(param, "=")
        if len(keyValue) != 2 {
            fmt.Println("Formato de parâmetro inválido:", param)
            continue
        }

        key := keyValue[0]
        value := keyValue[1]
        fmt.Printf("parametro: \"%s\" - Valor: \"%s\"\n", key, value)
    }
}
