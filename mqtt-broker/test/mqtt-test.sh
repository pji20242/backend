#!/bin/bash

# Loop infinito para publicar mensagens a cada 1 segundo
while true; do
    # Gera um valor de temperatura aleatório entre 1 e 50
    temperatura=$((RANDOM % 50 + 1))
    
    # Publica o valor de temperatura usando o mosquitto_pub
    mosquitto_pub -h 127.0.0.1 -p 1883 -t pji3 -m "temperatura $temperatura"
    
    # Aguarda 1 segundo antes de repetir
    sleep 1
done
