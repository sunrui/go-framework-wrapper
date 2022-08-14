#!/bin/bash
cd ../api/public || exit
pwd
go mod tidy -v
cd ../admin || exit
pwd
go mod tidy -v
cd ../user || exit
pwd
go mod tidy -v

cd ../../framework || exit
pwd
go mod tidy -v
cd ../generate || exit
pwd
go mod tidy -v
cd ../middleware || exit
pwd
go mod tidy -v
cd ../service || exit
pwd
go mod tidy -v
