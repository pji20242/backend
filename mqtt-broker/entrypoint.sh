#!/bin/sh

apk update && apk add --no-cache postgresql-dev

cd /mosquitto-auth-plug/

make clean
make

# Start the MQTT broker
/usr/sbin/mosquitto -v -c /mosquitto/config/mosquitto.conf 
