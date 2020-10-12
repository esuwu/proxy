#!/bin/sh

ID=$(id -u) # saves your user id in the ID variable

DOCKER_HOST_NET=$(ip -4 addr show docker0 | grep -Po 'inet \K[\d.]+')

GRAFANA_DIR="~/Projects/grafana"

GRAFANA_PROVISIONING="$GRAFANA_DIR/grafana_provisioning"
GRAFANA_STORAGE="$GRAFANA_DIR/grafana_storage"

mkdir -p "$GRAFANA_PROVISIONING"
mkdir -p "$GRAFANA_STORAGE"

sudo docker run \
        -d \
        --user $ID \
        --add-host host.docker.internal:$DOCKER_HOST_NET \
        -p 3000:3000 \
        --name grafana \
        --network bridge \
        -v $GRAFANA_PROVISIONING:/etc/grafana/provisioning \
        -v $GRAFANA_STORAGE:/var/lib/grafana \
        grafana/grafana
