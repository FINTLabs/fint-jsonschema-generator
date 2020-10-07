#!/usr/bin/env bash

docker build -t jsonschema-generator --build-arg VERSION=0.$(date +%y%m%d.%H%M) .
