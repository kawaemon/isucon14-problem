#!/bin/bash
set -exu

go build -o bencher .
./bencher run -s --target http://192.168.0.249:8080 --payment-url http://192.168.0.101:12345 || true
