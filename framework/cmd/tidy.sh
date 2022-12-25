#!/bin/bash
cd ../../framework || exit
pwd
go mod tidy -v
cd ../medium || exit
pwd
go mod tidy -v
cd ../api/admin || exit
pwd
go mod tidy -v
cd ../user || exit
pwd
go mod tidy -v
cd ../public || exit
pwd
go mod tidy -v