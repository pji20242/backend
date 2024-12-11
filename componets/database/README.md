
# Descrição Relacional Simples

```
usuario(matricula, nome, email, senha, privilegio, ativo)

dispositivo(uuid, hwversion, swversion, latitude, longitude, altitude)

dispositivo_usuario(matricula, uuid)
    matricula -> usuario(matricula)
    uuid -> dispositivo(uuid)        

sensor(idSensor, uuid, tipo, unidade)
    uuid -> dispositivo(uuid)

dados(timeStamp, uuid, idSensor, valor)
    idSensor, uuid -> sensor(idSensor, uuid)
```
