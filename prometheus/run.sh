#!/bin/bash

cd $(dirname $0)
if [ ! -x prometheus-2.43.0.linux-amd64/prometheus ]; then
    curl -OL https://github.com/prometheus/prometheus/releases/download/v2.43.0/prometheus-2.43.0.linux-amd64.tar.gz
    tar xzfv prometheus-2.43.0.linux-amd64.tar.gz
fi

./prometheus-2.43.0.linux-amd64/prometheus --config.file=prometheus.yml