#!/bin/bash
source env.sh

for i in $(seq 0 $MAX_NODE_IDX)
do
	./run_validator.sh $i
done
