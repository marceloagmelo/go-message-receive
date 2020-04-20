#!/usr/bin/env bash

source setenv.sh

# Rabbitmq send
echo "Subindo o go-message-receive..."
docker run -d --name go-message-receive --network $DOCKER_NETWORK  \
-p 8282:8080 \
-e RABBITMQ_USER=${RABBITMQ_USER} \
-e RABBITMQ_PASS=${RABBITMQ_PASS} \
-e RABBITMQ_HOSTNAME=${RABBITMQ_HOSTNAME} \
-e RABBITMQ_PORT=${RABBITMQ_PORT} \
-e RABBITMQ_VHOST=${RABBITMQ_VHOST} \
-e API_SERVICE_URL=${API_SERVICE_URL} \
marceloagmelo/go-message-receive

# Listando os containers
docker ps
