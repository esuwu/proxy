#!/bin/sh

ID=$(id -u) # saves your user id in the ID variable

PROM_DATA_DIR="~/Projects/prometheus/prom_db"

mkdir -p $PROM_DATA_DIR

sudo docker run \
        -d \
        --user $ID \
        -p 9090:9090 \
        --name prometheus \
        --network bridge \
        -v ~/Projects/prometheus/prometheus.yml:/etc/prometheus/prometheus.yml \
        -v ~/Projects/prometheus/prom_db/:/prometheus/ \
        prom/prometheus
