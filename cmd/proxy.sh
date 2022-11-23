#!/usr/bin/env bash
echo "export GO111MODULE=on" >> ~/.profile
echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
source ~/.profile
echo "export GO111MODULE=on" >> ~/.zprofile
echo "export GOPROXY=https://goproxy.cn" >> ~/.zprofile
source ~/.zprofile
