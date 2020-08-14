#!/bin/bash

protoc -I ./proto/ ./proto/gobonsai.proto --go_out=plugins=grpc:proto
