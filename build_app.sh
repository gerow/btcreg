#!/bin/sh

mkdir app
cd app
echo "executing go build..."
go build github.com/gerow/btcreg/btcreg
echo "done"
echo "copying necessary files..."
cp -r ../templates ../static .
echo "done"
