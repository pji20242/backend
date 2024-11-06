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

## Mensageria 

A estrutura de mensagens para publicação no brokerMQTT é a seguinte: 

```
47e207dc-b841-4ddb-9b43-a93df6a73e7e%1=12%2=20%3=30
```
Onde os campos são: 

 - `47e207dc-b841-4ddb-9b43-a93df6a73e7e`: Corresponde ao UUID do device que está tentando publicar no banco (UUIDv4)

 - `1=12`: Corresponde a primeira tupla de valores id%valor do sensor que está publicando a mensagem (inteiro não nulo)

 - `2=20`: Corresponde a segunda tupla de valores id%valor do sensor que está publicando a mensagem (inteiro não nulo)

  - `3=30`: Corresponde a terceira tupla de valores id%valor do sensor que está publicando a mensagem (inteiro não nulo)

  - `%`: Caracter delimitador de parâmetros, separa os dados entre sensores. 

O mapeamento de valores do parâmetro `idsensor` no banco é o seguinte: 

```
1 - Temperatura
2 - Pressão
3 - luminosidade
4 - Umidade
5 - Corrente
6 - Tensão
```

## Arquivos sensíveis

### Banco de dados

- `database/secret-db`: senha do administrador do MySQL. Exemplo:

```ini
senha
```

- `database/setup.sql`: preparação do banco de dados. Exemplo:

```sql
CREATE TABLE usuario (
    matricula INT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL,
    privilegio VARCHAR(50),
    ativo BOOLEAN DEFAULT TRUE
);

CREATE TABLE dispositivo (
    uuid CHAR(36) PRIMARY KEY,
    hwversion VARCHAR(50),
    swversion VARCHAR(50),
    latitude DECIMAL(9, 6),
    longitude DECIMAL(9, 6),
    altitude DECIMAL(9, 2)
);

CREATE TABLE dispositivo_usuario (
    matricula INT,
    uuid CHAR(36),
    PRIMARY KEY (matricula, uuid),
    FOREIGN KEY (matricula) REFERENCES usuario(matricula) ON DELETE CASCADE,
    FOREIGN KEY (uuid) REFERENCES dispositivo(uuid) ON DELETE CASCADE
);

CREATE TABLE sensor (
    idSensor INT,
    uuid CHAR(36),
    tipo VARCHAR(50),
    unidade VARCHAR(20),
    PRIMARY KEY (idSensor, uuid),
    UNIQUE (uuid, idSensor),  -- Adicionada uma restrição UNIQUE composta para (uuid, idSensor)
    FOREIGN KEY (uuid) REFERENCES dispositivo(uuid) ON DELETE CASCADE
);


CREATE TABLE dados (
    ts TIMESTAMP,
    uuid CHAR(36),
    idSensor INT,
    valor DECIMAL(10, 2),
    PRIMARY KEY (ts, uuid, idSensor),
    FOREIGN KEY (uuid, idSensor) REFERENCES sensor(uuid, idSensor) ON DELETE CASCADE
);

```

### Broker MQTT:

- `mqtt-broker/mosquitto.cfg`: variáveis de configuração ambiente.

```ini
MQTT_BROKER=mqtt-broker
MQTT_BROKER=mosquitto
MQTT_PORT=1883
```
