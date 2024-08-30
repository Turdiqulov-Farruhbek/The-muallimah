#!/bin/sh

CURRENT_DIR=$(dirname "$0")

PROTO_DIR=${1}
OUT_DIR=${CURRENT_DIR}/../genproto

mkdir -p ${OUT_DIR}\

protoc  -I=${PROTO_DIR} \
  --go_out=./ --go_opt= \
  --go-grpc_out=./ --go-grpc_opt= \
  ${PROTO_DIR}/muallimah-submodule/protos/*.proto

