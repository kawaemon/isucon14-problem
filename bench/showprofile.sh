#!/bin/sh

go tool pprof -http=localhost:5483 "$@"
