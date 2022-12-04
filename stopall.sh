#!/bin/bash
killall bootnode

# stop beaconnode
ps -ef | grep beaconnode | grep -v grep| awk '{print $2}' | xargs kill -9
rm bnode*.log vnode*.log

# stop validator
ps -ef | grep validator | grep -v grep| awk '{print $2}' | xargs kill -9
