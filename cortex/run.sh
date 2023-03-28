#!/bin/bash

cd $(dirname $0)
if [ ! -f cortex ]; then
    curl -Lo cortex https://github.com/cortexproject/cortex/releases/download/v1.14.1/cortex-linux-amd64
    chmod +x cortex
fi

./cortex -config.file=single-process-config-blocks-local.yaml

