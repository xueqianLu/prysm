#!/bin/bash
killall bootnode

# stop beaconnode
ps -ef | grep beaconnode | grep -v grep| awk '{print $2}' | xargs kill -9
rm bnode*.log vnode*.log

# stop validator
ps -ef | grep validator | grep -v grep| awk '{print $2}' | xargs kill -9

./deldata.sh


curdir=$PWD

# start beacon bootnode
./run_bootnode.sh 
sleep 1

# start beaconnode manual
./run_all_beacon.sh

# start validator manual
./run_all_validator.sh 
