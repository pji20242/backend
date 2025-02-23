package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	_ "github.com/go-sql-driver/mysql"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("mysql", "root:rootpass@tcp(database:3306)/pjiot")
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Erro ao pingar o banco de dados: %v", err)
	}
	fmt.Println("Conectado ao banco de dados MySQL!")
}

func main() {
	initDB()
	defer db.Close()

	mqttUsername := os.Getenv("MQTT_USERNAME")
	mqttPassword := os.Getenv("MQTT_PASSWORD")

	if mqttUsername == "" || mqttPassword == "" {
		log.Fatal("As variáveis de ambiente MQTT_USERNAME ou MQTT_PASSWORD não estão definidas.")
	}

	opts := MQTT.NewClientOptions().AddBroker("tcp://mqtt-broker:1883")
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(mqttUsername)
	opts.SetPassword(mqttPassword)

	opts.OnConnect = func(c MQTT.Client) {
		fmt.Println("Conectado ao broker MQTT!")
	}
	opts.OnConnectionLost = func(c MQTT.Client, err error) {
		fmt.Printf("Conexão perdida: %v\n", err)
	}

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Erro ao conectar ao broker: %v", token.Error())
	}

	topic := "pji3"
	if token := client.Subscribe(topic, 1, func(client MQTT.Client, msg MQTT.Message) {
		processMessage(string(msg.Payload()))
	}); token.Wait() && token.Error() != nil {
		log.Fatalf("Erro ao subscrever ao tópico: %v", token.Error())
	}

	fmt.Println("Aguardando mensagens...")
	select {}
}

func processMessage(message string) {
	parts := strings.Split(message, "%")
	if len(parts) < 2 {
		fmt.Println("Formato da mensagem inválido.")
		return
	}

	uuid := parts[0]
	fmt.Println("UUID:", uuid)

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

		tipoSensor := keyValue[0]
		valor := keyValue[1]

		idSensor := getSensorID(tipoSensor, uuid)
		inserirDados(uuid, idSensor, valor)
		inserirDadosInflux(uuid, tipoSensor, valor)
	}
}

func getSensorID(tipoSensor, uuid string) int {
	var idSensor int
	err := db.QueryRow("SELECT idSensor FROM sensor WHERE tipo = ? AND uuid = ?", tipoSensor, uuid).Scan(&idSensor)
	if err != nil {
		log.Printf("Erro ao obter o idSensor para tipo %s e uuid %s: %v", tipoSensor, uuid, err)
	}
	return idSensor
}

func inserirDados(uuid string, idSensor int, valor string) {
	valorFloat, err := strconv.ParseFloat(valor, 64)
	if err != nil {
		log.Printf("Erro ao converter o valor '%s' para float: %v", valor, err)
		return
	}

	_, err = db.Exec("INSERT INTO dados (ts, uuid, idSensor, valor) VALUES (?, ?, ?, ?)",
		time.Now(), uuid, idSensor, valorFloat)
	if err != nil {
		log.Printf("Erro ao inserir dados no banco de dados: %v", err)
	} else {
		fmt.Printf("Dados inseridos: UUID=%s, idSensor=%d, Valor=%.2f\n", uuid, idSensor, valorFloat)
	}
}

// inserirDadosInflux envia os dados para o InfluxDB
func inserirDadosInflux(uuid string, idSensor string, valor string) {
	// Converte o valor para float
	valorFloat, err := strconv.ParseFloat(valor, 64)
	if err != nil {
		log.Printf("Erro ao converter o valor '%s' para float: %v", valor, err)
		return
	}

	// Recupera as variáveis de ambiente
	influxURL := os.Getenv("INFLUXDB_URL")
	influxToken := os.Getenv("INFLUXDB_TOKEN")
	influxOrg := os.Getenv("INFLUXDB_ORG")
	influxBucket := os.Getenv("INFLUXDB_BUCKET")

	if influxURL == "" || influxToken == "" || influxOrg == "" || influxBucket == "" {
		log.Println("Erro: Variáveis de ambiente do InfluxDB não estão definidas corretamente.")
		return
	}

	// Criar cliente do InfluxDB
	client := influxdb2.NewClient(influxURL, influxToken)
	defer client.Close()

	// Criar escritor para enviar os dados
	writeAPI := client.WriteAPIBlocking(influxOrg, influxBucket)

	// Criar ponto de medição para inserir no banco
	p := influxdb2.NewPoint(
		"dados_sensores",
		map[string]string{"uuid": uuid, "sensor_id": idSensor},
		map[string]interface{}{"valor": valorFloat},
		time.Now(),
	)

	// Escrever no InfluxDB
	err = writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		log.Printf("Erro ao inserir dados no InfluxDB: %v", err)
	} else {
		fmt.Printf("Dados inseridos no InfluxDB: UUID=%s, idSensor=%s, Valor=%.2f\n", uuid, idSensor, valorFloat)
	}
}
