#!/bin/sh
# Make sure we're in the right directory
cd "$(dirname "$0")"

echo "downloading go dependencies..."
go get
echo "done"
echo "creating application directory..."
mkdir app
echo "done"
cd app
echo "executing go build..."
go build github.com/gerow/btcreg/btcreg
echo "done"
echo "copying necessary files..."
cp -r ../templates ../static .
echo "done"
