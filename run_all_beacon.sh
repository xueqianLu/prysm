#!/bin/bash
source env.sh
pidtmp=peerinfo

rm $pidtmp > /dev/null
pidinfo=""

currentpid=""
ip=`ip a | grep eth0 | grep "inet" | grep -Eo "[0-9\.]+" | head -n 1`

function waitcurrentpid()
{
	logfile="bnode${1}.log"
	simplepid=""
	while [ "$simplepid" == "" ]; do
		sleep 1
		simplepid=`grep "Running node with peer id of" $logfile | awk '{print $10}'`

		echo "wait node $1 peerid"
	done
	tcpport=$((33000+$1))
	currentpid="/ip4/$ip/tcp/$tcpport/p2p/$simplepid"
}

for i in $(seq 0 $MAX_NODE_IDX)
do
#	echo "run beaconnode with pidinfo $pidinfo"
	./run_beaconnode.sh $i 
	sleep 1
	#waitcurrentpid $i
	#pidinfo=`cat $pidtmp`
	#pidinfo="--peer $currentpid $pidinfo"
	#echo "$pidinfo" > $pidtmp
	#currentpid=""
done
