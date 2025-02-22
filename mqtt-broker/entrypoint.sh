#!/bin/sh

# sleep 5 seconds to wait for the network to be ready
# echo "Sleeping for 10 seconds to wait for the network to be ready..."
# sleep 10
apk update && apk add --no-cache postgresql-dev

# Start the MQTT broker
/usr/sbin/mosquitto -v -c /mosquitto/config/mosquitto.conf 

tail -f /dev/null
# tail the logs
# tail -f /var/log/mosquitto/mosquitto.log

