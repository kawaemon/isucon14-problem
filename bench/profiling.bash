#!/bin/bash
set -exu

graph=.kprivate/flamegraph
addr2line=.kprivate/addr2line

if [ ! -d $graph ]; then
	    git clone https://github.com/brendangregg/FlameGraph.git $graph --filter=blob:none
fi
if [ ! -d $addr2line ]; then
	    git clone https://github.com/gimli-rs/addr2line.git $addr2line --filter=blob:none
fi

pushd $addr2line
cargo b --features=bin --release --bin=addr2line
popd

# wget -O profile 'http://localhost:6060/debug/pprof/profile?seconds=30'

go tool pprof \
    -raw -output=perf.data.scripted \
    'http://localhost:6060/debug/pprof/profile?seconds=60'

cat ./perf.data.scripted | $graph/stackcollapse-go.pl > perf.data.collapsed
cat ./perf.data.collapsed | $graph/flamegraph.pl --width 1920 > out.svg
chromium out.svg
