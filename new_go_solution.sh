#!/bin/bash

readonly GO_STARTER_CODE="iolib/iolib2.go"

mkdir problems/$1
cp $GO_STARTER_CODE problems/$1/$1.go

exit 0