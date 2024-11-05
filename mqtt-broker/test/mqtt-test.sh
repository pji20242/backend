#!/bin/bash

# Loop infinito para publicar mensagens a cada 1 segundo
while true; do
    # Gera um valor de temperatura, pressao e luminosidade aleatório
    temperatura=$((RANDOM % 50 + 1))
    pressao=$((RANDOM % 200 + 1))
    luminosidade=$((RANDOM % 1500 + 1))
    umidade=$((RANDOM % 100 + 1))
    tensao=$((RANDOM % 100 + 1))
    corrente=$((RANDOM % 100 + 1))

    
    # Publica os sensores de temperatura, pressao e luminosidade do device 1 
    mosquitto_pub -h 127.0.0.1 -p 1883 -t pji3 -m "1f3cbe5b-15dd-483e-a74a-bea00227da11%T=$temperatura%P=$pressao%L=$luminosidade"

    # Publica os sensores de umidade do device 2
    mosquitto_pub -h 127.0.0.1 -p 1883 -t pji3 -m "7f7b30cd-3a52-46b7-8615-feff437503e5%U=$umidade"

    # Publica os sensores de tensao e corrente do device 3
    mosquitto_pub -h 127.0.0.1 -p 1883 -t pji3 -m "375311ba-6e97-4c19-8c9d-45c5c479a520%Vol=$tensao%Amp=$corrente"
    
    # Aguarda 1 segundo antes de repetir
    sleep 1
done
