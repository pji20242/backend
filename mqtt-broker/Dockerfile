FROM arthurcadore/mqtt-broker-auth:2.0

RUN apk update && apk add --no-cache postgresql-dev

CMD ['tail', '-f', '/dev/null']
# Comando para rodar o Mosquitto
# CMD ["/usr/sbin/mosquitto", "-c", "/mosquitto/config/mosquitto.conf"]
