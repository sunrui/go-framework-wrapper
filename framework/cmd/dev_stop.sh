#!/bin/bash
redis-cli shutdown
mysql.server stop
pkill redis
pkill mysql