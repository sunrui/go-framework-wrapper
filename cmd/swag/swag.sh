#!/usr/bin/env bash
#
# Copyright (c) 2022 honeysense.com All rights reserved.
# Author: sunrui
# Date: 2022-12-01 02:26:39
#
swag="$PWD/swag"

# shellcheck disable=SC2164
cd ../../passport/api/public

$swag fmt
$swag init --parseDependency