package main

import (
    "fmt"
    "log"

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
        fmt.Printf("Mensagem recebida no tópico %s: %s\n", msg.Topic(), msg.Payload())
    }); token.Wait() && token.Error() != nil {
        log.Fatalf("Erro ao subscrever ao tópico: %v", token.Error())
    }

    // Manter o programa em execução
    fmt.Println("Aguardando mensagens...")
    select {}
}
