USE pjiot;

CREATE TABLE user (
    matricula INT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL,
    privilegio VARCHAR(50),
    ativo BOOLEAN DEFAULT TRUE
);

CREATE TABLE device (
    uuid CHAR(36) PRIMARY KEY,
    hwversion VARCHAR(50) NOT NULL,
    swversion VARCHAR(50) NOT NULL,
    latitude DECIMAL(9, 6) NOT NULL,
    longitude DECIMAL(9, 6) NOT NULL, 
    altitude DECIMAL(9, 2) NOT NULL
);

CREATE TABLE device_user (
    matricula INT,
    uuid CHAR(36),
    PRIMARY KEY (matricula, uuid),
    FOREIGN KEY (matricula) REFERENCES user(matricula) ON DELETE CASCADE,
    FOREIGN KEY (uuid) REFERENCES device(uuid) ON DELETE CASCADE
);

CREATE TABLE sensor (
    idSensor INT,
    uuid CHAR(36),
    tipo VARCHAR(50),
    unidade VARCHAR(20),
    PRIMARY KEY (idSensor, uuid),
    FOREIGN KEY (uuid) REFERENCES device(uuid) ON DELETE CASCADE
);

CREATE TABLE dados (
    timeStamp TIMESTAMP,
    uuid CHAR(36),
    idSensor INT,
    valor DECIMAL(10, 2),
    PRIMARY KEY (timeStamp, uuid, idSensor),
    FOREIGN KEY (uuid, idSensor) REFERENCES sensor(uuid, idSensor) ON DELETE CASCADE
);


CREATE USER 'connectoruser'@'%' IDENTIFIED BY 'connectorpasswrd';
GRANT ALL PRIVILEGES ON pjiot.* TO 'connectoruser'@'%';
FLUSH PRIVILEGES;

