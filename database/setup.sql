CREATE DATABASE IF NOT EXISTS pjiot;
USE pjiot;

-- Tabela cooperativa
CREATE TABLE cooperativa (
    cnpj CHAR(14) PRIMARY KEY,
    endereco VARCHAR(255),
    email VARCHAR(100),
    nome VARCHAR(100) NOT NULL
);

-- Tabela privilégio
CREATE TABLE privilegio (
    id INT AUTO_INCREMENT PRIMARY KEY,
    descricao VARCHAR(50) NOT NULL
);

-- Tabela usuario
CREATE TABLE usuario (
    matricula INT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL,
    user VARCHAR(50) NOT NULL,
    ativo BOOLEAN DEFAULT TRUE,
    licencas VARCHAR(50)
);

-- Tabela dispositivo
CREATE TABLE dispositivo (
    uuid CHAR(36) PRIMARY KEY,
    hw_version VARCHAR(50),
    fw_version VARCHAR(50),
    latitude DECIMAL(9, 6),
    longitude DECIMAL(9, 6),
    peso DECIMAL(10, 2),
    alt DECIMAL(9, 6)
);

-- Tabela usuario_cooperativa
CREATE TABLE usuario_cooperativa (
    matricula INT,
    cnpj CHAR(14),
    idPrivilegio INT,
    PRIMARY KEY (matricula, cnpj),
    FOREIGN KEY (matricula) REFERENCES usuario(matricula) ON DELETE CASCADE,
    FOREIGN KEY (cnpj) REFERENCES cooperativa(cnpj) ON DELETE CASCADE,
    FOREIGN KEY (idPrivilegio) REFERENCES privilegio(id) ON DELETE CASCADE
);

-- Relacionamento dispositivo-usuario (licenciado)
CREATE TABLE dispositivo_usuario (
    matricula INT,
    uuid CHAR(36),
    licenciado BOOLEAN DEFAULT TRUE,
    PRIMARY KEY (uuid),
    FOREIGN KEY (matricula) REFERENCES usuario(matricula) ON DELETE CASCADE,
    FOREIGN KEY (uuid) REFERENCES dispositivo(uuid) ON DELETE CASCADE
);

-- Tabela sensor
CREATE TABLE sensor (
    idSensor INT AUTO_INCREMENT,
    uuid CHAR(36),
    tipo VARCHAR(50),
    unidade VARCHAR(20),
    PRIMARY KEY (idSensor, uuid),
    UNIQUE (uuid, idSensor),
    FOREIGN KEY (uuid) REFERENCES dispositivo(uuid) ON DELETE CASCADE
);

-- Tabela dados
CREATE TABLE dados (
    ts TIMESTAMP,
    uuid CHAR(36),
    idSensor INT,
    valor DECIMAL(10, 2),
    PRIMARY KEY (ts, uuid, idSensor),
    FOREIGN KEY (uuid, idSensor) REFERENCES sensor(uuid, idSensor) ON DELETE CASCADE
);

-- Usuário para conexão
CREATE USER 'connectoruser'@'%' IDENTIFIED BY 'connectorpasswrd';
GRANT ALL PRIVILEGES ON pjiot.* TO 'connectoruser'@'%';
FLUSH PRIVILEGES;
