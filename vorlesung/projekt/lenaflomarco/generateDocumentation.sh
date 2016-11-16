#!/usr/bin/env bash

cd res/doc/

godoc -http=:9090 &
wget -e robots=off -m http://localhost:9090/pkg/FileServer/