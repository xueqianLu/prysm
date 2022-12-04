#!/bin/bash
source env.sh

for i in $(seq 0 $MAX_NODE_IDX)
do
	rm -rf ../beaconnode/node${i}/beaconchaindata/beaconchain.db
	rm -rf ../validator/node${i}/validator.db
done
rm -f bnode*.log vnode*.log
