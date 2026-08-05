#!/bin/bash

set -e

BASE_URL="http://localhost:8080"

echo "===================================="
echo "Config Service Smoke Test"
echo "===================================="

echo ""
echo "Checking Health Endpoint..."

curl --fail -s ${BASE_URL}/ping

echo ""
echo "Health Check Passed"

echo ""
echo "Creating Configuration..."

curl --fail -s -X POST ${BASE_URL}/configs \
-H "Content-Type: application/json" \
-d '{
"id":"cfg-test",
"host":"localhost",
"port":8080,
"app_name":"config-service",
"log_level":"INFO"
}'

echo ""
echo "POST Passed"

echo ""
echo "Retrieving Configuration..."

curl --fail -s ${BASE_URL}/configs/cfg-test

echo ""
echo "GET Passed"

echo ""
echo "Smoke Test Passed ✅"