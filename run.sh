#!/bin/bash
./run_bootnode.sh
sleep 5
for i in $(seq 0 3)
do
	./run_beaconnode.sh $i
done


for i in $(seq 0 3)
do
	./run_validator.sh $i
done
