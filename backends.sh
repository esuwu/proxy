#!/bin/sh

#OUT_PATH="~/server/dummy.out"

OUT_PATH="~"
TMP_ENV_DIR="/Users/adolgavin/Projects/proxy"
TMP_ENV="$TMP_ENV_DIR/.env"
DUMMY_PATH="/Users/adolgavin/Projects/proxy/main"
KEY_PATH="~/Downloads/SSH.pem"

echo "BACKEND_ID=1" > "$TMP_ENV"
scp -i "$KEY_PATH" "$DUMMY_PATH" "$TMP_ENV" ubuntu@ec2-18-221-25-126.us-east-2.compute.amazonaws.com:$OUT_PATH

echo "BACKEND_ID=2" > "$TMP_ENV"
scp -i "$KEY_PATH" "$DUMMY_PATH" "$TMP_ENV" ubuntu@ec2-18-220-52-248.us-east-2.compute.amazonaws.com:$OUT_PATH

echo "BACKEND_ID=3" > "$TMP_ENV"
scp -i "$KEY_PATH" "$DUMMY_PATH" "$TMP_ENV" ubuntu@ec2-18-221-104-251.us-east-2.compute.amazonaws.com:$OUT_PATH

