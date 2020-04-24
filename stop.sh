#!/usr/bin/env bash

source setenv.sh

echo "Finalizando ${APP_NAME}..."
docker rm -f ${APP_NAME}

