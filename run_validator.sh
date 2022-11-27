#!/bin/bash
nodeidx=${1:-""}
basedir=$PWD
keydir="$basedir/validator_keys/v$nodeidx"
walldir="$basedir/../validator/wallet$nodeidx"
datadir="$basedir/../validator/node$nodeidx"
password="$basedir/password.txt"


bazel run //cmd/validator:validator --config=minimal -- accounts import --keys-dir=$keydir --wallet-dir=$walldir --wallet-password-file $password

monitorport=$((8090+$nodeidx))
pprofport=$((6060+$nodeidx))
cors="*"
grpcport=$((7500+$nodeidx))
rpcport=$((7000+$nodeidx))

bazel run //cmd/validator:validator --config=minimal -- --wallet-dir=$walldir --chain-config-file $PWD/testnet.yaml --wallet-password-file $password --datadir=$datadir --monitoring-port $monitorport --pprofport $pprofport --grpc-gateway-corsdomain '*' --grpc-gateway-port $grpcport --rpc-port $rpcport --verbosity debug > vnode$nodeidx.log 2>&1 &

#./prysm.sh validator accounts import --keys-dir=$dir  --wallet-dir=$PWD/wallets
#./prysm.sh validator --wallet-dir=$PWD/wallets --chain-config-file $PWD/testnet.yaml

#bazel run //cmd/validator:validator --config=minimal -- --wallet-password-file password.txt --datadir

