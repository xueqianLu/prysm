#!/bin/bash
MAX_NODE_NUM_IDX=15
./run_bootnode.sh
sleep 1
for i in $(seq 0 $MAX_NODE_NUM_IDX)
do
	./run_beaconnode.sh $i
done


for i in $(seq 0 $MAX_NODE_NUM_IDX)
do
	./run_validator.sh $i
done
