#!/bin/bash

# Loop infinito para publicar mensagens a cada 1 segundo
while true; do
    # Gera um valor de temperatura, pressao e luminosidade aleatório
    temperatura=$((RANDOM % 50 + 1))
    pressao=$((RANDOM % 200 + 1))
    luminosidade=$((RANDOM % 1500 + 1))
    
    # Publica o valor de temperatura usando o mosquitto_pub
    mosquitto_pub -h 127.0.0.1 -p 1883 -t pji3 -m "1f3cbe5b-15dd-483e-a74a-bea00227da11%T=$temperatura%P=$pressao%L=$luminosidade@"
    
    # Aguarda 1 segundo antes de repetir
    sleep 1
done
