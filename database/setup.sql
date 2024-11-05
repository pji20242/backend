USE pjiot;

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

CREATE USER 'connectoruser'@'%' IDENTIFIED BY 'connectorpasswrd';
GRANT ALL PRIVILEGES ON pjiot.* TO 'connectoruser'@'%';
FLUSH PRIVILEGES;

