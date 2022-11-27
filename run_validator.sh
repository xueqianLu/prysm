#!/bin/bash
nodeidx=${1:-""}
basedir=$PWD
keydir="$basedir/validator_keys/v$nodeidx"
walldir="$basedir/../validator/wallet$nodeidx"
datadir="$basedir/../validator/node$nodeidx"
password="$basedir/password.txt"


#bazel run //cmd/validator:validator --config=minimal -- accounts import --keys-dir=$keydir --wallet-dir=$walldir --wallet-password-file $password

monitorport=$((8090+$nodeidx))
pprofport=$((6060+$nodeidx))
cors="*"
grpcport=$((7500+$nodeidx))
rpcport=$((7000+$nodeidx))
beaconrpcport=$((4000+$nodeidx))
beacongrpcport=$((3500+$nodeidx))

bazel run //cmd/validator:validator --config=minimal -- --wallet-dir=$walldir --chain-config-file $PWD/testnet.yaml --wallet-password-file $password --datadir=$datadir --monitoring-port $monitorport --pprofport $pprofport --grpc-gateway-corsdomain '*' --grpc-gateway-port $grpcport --rpc-port $rpcport --verbosity debug --beacon-rpc-gateway-provider "127.0.0.1:$beacongrpcport" --beacon-rpc-provider "127.0.0.1:$beaconrpcport" > vnode$nodeidx.log 2>&1 &

