#!/bin/bash

if [[ $1 == "" ]]; then
    go run ./src/main.go ./sample-config.yml

fi


if [[ $1 == "build" ]]; then
    go build ./src/main.go

fi
