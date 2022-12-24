#!/usr/bin/env bash
#
# Copyright (c) 2022 honeysense.com All rights reserved.
# Author: sunrui
# Date: 2022-12-01 02:26:39
#
swag="$PWD/swag"

# shellcheck disable=SC2164
cd ../../../api/admin
$swag fmt
$swag init --parseDependency --parseInternal

# shellcheck disable=SC2164
cd ../public
$swag fmt
$swag init --parseDependency --parseInternal

# shellcheck disable=SC2164
cd ../user
$swag fmt
$swag init --parseDependency --parseInternal