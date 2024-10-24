# Importa as bibliotecas mysql.connector e paho.mqtt.client
import mysql.connector
import paho.mqtt.client as mqtt
from dotenv import load_dotenv
from os import getenv
load_dotenv()

# Configura as variaveis do banco de dado
db_host = getenv("DB_HOST", default="database")
db_user = getenv("DB_USER", default="yourUser")
db_password = getenv("DB_PASSWORD", default="")
db_database = getenv("DB_DATABASE", default="currentTS")

# Configura as variaveis do MQTT
mqtt_broker = getenv("MQTT_BROKER", default="mqtt-broker")
mqtt_port = getenv("MQTT_PORT", default=1883)
mqtt_luminosity_topic = getenv("MQTT_LUMINOSITY_TOPIC", default="luminosidade")
mqtt_temperature_topic = getenv("MQTT_TEMPERATURE_TOPIC", default="temperatura")


# Função camha quanddo ua mentsagem é recebida
def on_message(client, userdata, message):
    mensagem = message.payload.decode("utf-8")
    print(f"Mensagem recebida: {mensagem}")

    try:
        #Tenta conectar ao banco de dados
        connection = mysql.connector.connect(
            host=db_host,
            user=db_user,
            password=db_password,
            database=db_database
        )

        cursor = connection.cursor()

        # Determina a tabela de destino com base no tópico MQTT
        if message.topic == mqtt_luminosity_topic:
            insert_query = "INSERT INTO luminosidade (mensagem) VALUES (%s)"
        elif message.topic == mqtt_temperature_topic:
            insert_query = "INSERT INTO temperatura (mensagem) VALUES (%s)"
        else:
            print("Tópico não reconhecido. Dados não foram salvos no banco de dados.")
            return

        #Executa a consulta de inserção no banco de dados
        cursor.execute(insert_query, (mensagem,))
        connection.commit()

        print("Mensagem salva no banco de dados.")

    except mysql.connector.Error as error:
        print(f"Erro ao conectar ao banco de dados: {error}")

    finally:
        if connection.is_connected():
            cursor.close()
            connection.close()

# Configurar o cliente MQTT
client = mqtt.Client()
client.on_message = on_message

# Conectar ao broker MQTT
client.connect(mqtt_broker, mqtt_port, 60)
client.subscribe(mqtt_luminosity_topic)
client.subscribe(mqtt_temperature_topic)

# Iniciar o loop para escutar mensagens
client.loop_forever()
