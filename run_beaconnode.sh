#!/bin/bash
nodeidx=${1:-""}
peerinfo=`cat peerinfo`

gethrpcport=$((9545+$nodeidx))
rpcport=$((4000+$nodeidx))
grpcport=$((3500+$nodeidx))
monitorport=$((8080+$nodeidx))
tcpp2pport=$((33000+$nodeidx))
udpp2pport=$((32000+$nodeidx))

bootnode_tmp=`head -n 4 $PWD/bootnode.log | grep -Eo "enr:*.*\""`
bootnode=`echo ${bootnode_tmp%?}`

#echo "bootnode=$bootnode, peerinfo is $peerinfo"


#./beaconnode --execution-endpoint=$PWD/../node${nodeidx}/data/geth.ipc --suggested-fee-recipient=0x905D5E8F7db76bCA91fdcA0990be7263dfD23335 --chain-id 1024 --contract-deployment-block 3 --network-id 1024 --datadir ${PWD}/../beaconnode/node${nodeidx} --http-web3provider http://127.0.0.1:${gethrpcport} --no-discovery --chain-config-file ${PWD}/testnet.yaml --enable-debug-rpc-endpoints --force-clear-db --min-sync-peers 0 --p2p-local-ip=0.0.0.0 --verbosity=debug --monitoring-port $monitorport --rpc-port $rpcport --grpc-gateway-port $grpcport --p2p-tcp-port ${tcpp2pport} --p2p-udp-port ${udpp2pport} --bootstrap-node "$bootnode" > bnode${nodeidx}.log 2>&1 &
bazel run //cmd/beacon-chain:beacon-chain --config=minimal -- --execution-endpoint=$PWD/../node${nodeidx}/data/geth.ipc --suggested-fee-recipient=0x905D5E8F7db76bCA91fdcA0990be7263dfD23335 --chain-id 1024 --contract-deployment-block 3 --network-id 1024 --datadir ${PWD}/../beaconnode/node${nodeidx} --http-web3provider http://127.0.0.1:${gethrpcport} --no-discovery --chain-config-file ${PWD}/testnet.yaml --enable-debug-rpc-endpoints --force-clear-db --min-sync-peers 0 --p2p-host-ip=172.26.14.44 --verbosity=debug --monitoring-port $monitorport --rpc-port $rpcport --grpc-gateway-port $grpcport --p2p-tcp-port ${tcpp2pport} --p2p-udp-port ${udpp2pport} --bootstrap-node "$bootnode" $peerinfo > bnode${nodeidx}.log 2>&1 &
