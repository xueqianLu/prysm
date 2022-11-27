#!/bin/bash

for i in $(seq 0 3)
do
	./run_validator.sh $i
done
