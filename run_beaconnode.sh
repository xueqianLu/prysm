#!/bin/bash
nodeidx=${1:-""}
gethrpcport=$((9545+$nodeidx))
rpcport=$((4000+$nodeidx))
grpcport=$((3500+$nodeidx))
monitorport=$((8080+$nodeidx))
p2pport=$((13000+$nodeidx))

bazel run //cmd/beacon-chain:beacon-chain --config=minimal -- --execution-endpoint=$PWD/../node${nodeidx}/data/geth.ipc --suggested-fee-recipient=0x905D5E8F7db76bCA91fdcA0990be7263dfD23335 --chain-id 1024 --contract-deployment-block 3 --network-id 1024 --datadir ${PWD}/../beaconnode/node${nodeidx} --http-web3provider http://127.0.0.1:${gethrpcport} --no-discovery --chain-config-file ${PWD}/testnet.yaml --enable-debug-rpc-endpoints --force-clear-db --min-sync-peers 0 --p2p-local-ip=127.0.0.1 --verbosity=debug --monitoring-port $monitorport --rpc-port $rpcport --grpc-gateway-port $grpcport --p2p-tcp-port ${p2pport} --bootstrap-node "enr:-LG4QITO8KU8ewpVxPdoSs5PIZNjIej5ok2pGMVT0Qgskls2KpCezlg1srExr2bCpFJHWwV8h4YluPDya0eNKqmHiOeGAYS5TcSxh2F0dG5ldHOIAAAAAAAAAACEZXRoMpAYrkzLAAAAAf__________gmlkgnY0gmlwhKwaDiyJc2VjcDI1NmsxoQMIaOlMFePWmvIMGNZ8Y8ZdX1IK_aY9tfy_v496kabProN1ZHCCIyg" > bnode${nodeidx}.log 2>&1 &
