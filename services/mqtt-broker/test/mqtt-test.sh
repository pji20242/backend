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
    mosquitto_pub -h 127.0.0.1 -p 1883 -t pji3 -m "3170dd2b-f944-4835-8a4a-e2ab5dee3b25%1=$temperatura%2=$pressao%3=$luminosidade"

    # # Publica os sensores de umidade do device 2
    mosquitto_pub -h 127.0.0.1 -p 1883 -t pji3 -m "887010d3-a456-4572-b3b6-c5edfc7d765c%4=$umidade"


    # # Publica os sensores de tensao e corrente do device 3
    # mosquitto_pub -h 127.0.0.1 -p 1883 -t pji3 -m "375311ba-6e97-4c19-8c9d-45c5c479a520%5=$tensao%6=$corrente"
    
    # Aguarda 1 minuto antes de repetir
    sleep 1
done
