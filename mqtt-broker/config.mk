# Select your backends from this list
BACKEND_CDB ?= no
BACKEND_MYSQL ?= no
BACKEND_SQLITE ?= no
BACKEND_REDIS ?= no
BACKEND_POSTGRES ?= yes
BACKEND_LDAP ?= no
BACKEND_HTTP ?= no
BACKEND_JWT ?= no
BACKEND_GOLANG ?= no
BACKEND_MONGO ?= no
BACKEND_FILES ?= no

# Specify the path to the Mosquitto sources here
MOSQUITTO_SRC = /mosquitto/src

# Specify the path the OpenSSL here
OPENSSLDIR = /usr/include/openssl

# Add support for django hashers algorithm name
SUPPORT_DJANGO_HASHERS ?= no

# Substituindo a configuração MySQL por PostgreSQL
CFG_LDFLAGS = -lcares -lpq   # Adicionando a biblioteca do PostgreSQL (libpq)
CFG_CFLAGS = -I/mosquitto/src/include -I/mosquitto/src/lib/ -fPIC -Wall -Werror -DBE_POSTGRES -I/usr/include/openssl/include -I/usr/include/postgresql   # Incluindo os diretórios necessários para o PostgreSQL
