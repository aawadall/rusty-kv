#!/bin/bash
# Builds the protocol buffer files Locally 

protoc -I=proto/ --go_out=proto_api/ proto/*.proto