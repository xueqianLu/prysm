#!/bin/bash
bazel build //cmd/beacon-chain:beacon-chain --config=minimal
bazel build //cmd/validator:validator --config=minimal
