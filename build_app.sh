#!/bin/sh

mkdir app
cd app
go build github.com/gerow/btcreg/btcreg
cp -r ../templates ../static .
