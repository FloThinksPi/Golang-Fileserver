#!/usr/bin/env bash

godoc -http=:9090 &
sleep 2
wget -r -np -N -E -p -k -e robots=off http://localhost:9090/pkg/FileServer
killall godoc
rm -r res/doc/
mv localhost:9090 res/doc