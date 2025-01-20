#!/usr/bin/env zsh
set -exu

# https://meetup-jp.nhncloud.com/1509
sudo sysctl -w net.ipv4.tcp_max_syn_backlog=8192
sudo sysctl -w net.core.netdev_max_backlog="30000"
sudo sysctl -w net.ipv4.tcp_fin_timeout=1
sudo sysctl -w net.ipv4.tcp_max_tw_buckets="18000000"
# sudo sysctl -w ipv4.tcp_timestamps="1"
sudo sysctl -w net.ipv4.tcp_tw_reuse="1"

ulimit -n 512000
go build -o bencher .
./bencher run -s --target http://192.168.0.249:8080 --payment-url http://192.168.0.101:12345 || true
