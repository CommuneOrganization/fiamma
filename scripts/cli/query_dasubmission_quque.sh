#!/bin/bash

set -e


: ${CHAIN_ID:="fiamma-testnet-1"}
: ${NODE:="https://testnet-rpc.fiammachain.io"}

fiammad query zkpverify get-da-submission-queue --node $NODE
