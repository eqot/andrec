#!/bin/bash

cd src/github.com/eqot/andrec
gox -os="darwin linux windows" -arch="amd64" -output="../../../../bin/{{.Dir}}_{{.OS}}_{{.Arch}}"
cd -
