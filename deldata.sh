#!/bin/bash
for i in $(seq 0 3)
do
	rm -rf ../beaconnode/node${i}/beaconchaindata/beaconchain.db
	rm -rf ../validator/node${i}/validator.db
done

