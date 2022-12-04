#!/bin/bash
ip=`ip addr | grep -v inet6 | grep inet | grep -v docker | grep -v -w lo | grep -v "br-" | grep -Eo "172.[0-9.]+*" | head -n 1`
export MAX_NODE_IDX=15
export HOSTIP=$ip
