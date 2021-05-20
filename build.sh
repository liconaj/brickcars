#!/bin/sh

current=$PWD
cd assets
./go-bindata data/...
sed -i 's/main/assets/g' bindata.go
cd $current

[ ! -d  bin/i386 ] && mkdir -p bin/i386
[ ! -d  bin/x86_64 ] && mkdir -p bin/x86_64

GOOS=windows ARCH=x86_64 go build 
GOOS=linux ARCH=x86_64 go build
mv brickcars* bin/x86_64

GOOS=windows ARCH=386 go build
GOOS=linux ARCH=386 go build
mv brickcars* bin/i386
