#!/bin/bash
cd ../framework || exit
pwd
go mod tidy -v
cd ../passport || exit
pwd
go mod tidy -v