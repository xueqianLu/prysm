#!/bin/bash
pidtmp=peerinfo

rm $pidtmp > /dev/null
pidinfo=""

currentpid=""

function waitcurrentpid()
{
	logfile="bnode${1}.log"
	while [ "$currentpid" == "" ]; do
		sleep 1
		currentpid=`grep "Node started p2p server" $logfile | grep -Eo "/ip[0-9a-zA-Z\/\.]*"`
		echo "wait node $1 peerid"
	done
#	echo "got current peer info $currentpid"
}

for i in $(seq 0 3)
do
#	echo "run beaconnode with pidinfo $pidinfo"
	./run_beaconnode.sh $i 
	waitcurrentpid $i
	pidinfo=`cat $pidtmp`
	pidinfo="--peer $currentpid $pidinfo"
	echo "$pidinfo" > $pidtmp
	currentpid=""
done
