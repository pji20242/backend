# backend
Código fonte da infraestrutura e aplicação do servidor da solução

## Estrutura

```mermaid
flowchart TD
    DB([Banco de dados])
        DB --- backend

    mqtt([Broker MQTT])
        mqtt --- backend
    backend([App Backend GO])
        
    frontend([Frontend HTTP Server / API Restful])
        frontend --- backend
```

## Arquivos sensíveis

### Banco de dados

- `database/secret-db`: senha do administrador do MySQL. Exemplo:

```ini
senha
```

- `database/setup.sql`: preparação do banco de dados. Exemplo:

```sql
USE currentTS;

CREATE TABLE temperatura (
  id INT AUTO_INCREMENT PRIMARY KEY,
  mensagem TEXT,
  data_insercao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE luminosidade (
  id INT AUTO_INCREMENT PRIMARY KEY,
  mensagem TEXT,
  data_insercao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE corrente (
  id INT AUTO_INCREMENT PRIMARY KEY,
  mensagem TEXT,
  data_insercao TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE releac (
  id INT AUTO_INCREMENT PRIMARY KEY,
  mensagem TEXT
);

CREATE TABLE relelamp (
  id INT AUTO_INCREMENT PRIMARY KEY,
  mensagem TEXT
);

CREATE USER 'yourUser'@'%' IDENTIFIED BY 'senha';
GRANT ALL PRIVILEGES ON currentTS.* TO 'yourUser'@'%';
FLUSH PRIVILEGES;
```

### Conector

- `connector/.env`: variáveis de ambiente. Exemplo:

```ini
DB_HOST=database
DB_USER=yourUser
DB_PASSWORD=senha
DB_DATABASE=currentTS
MQTT_BROKER=mqtt-broker
MQTT_BROKER=mosquitto
MQTT_PORT=1883
MQTT_LUMINOSITY_TOPIC=luminosidade
MQTT_TEMPERATURE_TOPIC=temperatura
```
