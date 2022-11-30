#!/bin/bash
cd ../../framework || exit
pwd
go mod tidy -v
cd ../medium || exit
pwd
go mod tidy -v