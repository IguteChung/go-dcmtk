#!/bin/sh
set -e

DIR=$(cd $(dirname ${BASH_SOURCE[0]}) && pwd)


cd $DIR/dcmtk/dcmdata/build && cmake ../apps && make
cd $DIR/dcmtk/dcmimage/build && cmake ../apps && make
cd $DIR/dcmtk/dcmjpeg/build && cmake ../apps && make

cd $DIR && go run main.go
