#!/bin/bash

VAULT_HOST="http://localhost:1234"
TOKEN="myroot"
NAMESPACE=""

curl -X 'POST' \
  $VAULT_HOST/v1/sys/auth/approle \
  -H "accept: */*" \
  -H "Content-Type: application/json" \
  -H "X-Vault-Token: $TOKEN" \
  -H "X-Vault-Namespace: $NAMESPACE" \
  -d '{
    "path": "approle",
    "type": "approle",
    "config": {}
}'


curl -X 'POST' \
  $VAULT_HOST/v1/sys/policy/test_policy \
  -H "accept: */*" \
  -H "Content-Type: application/json" \
  -H "X-Vault-Token: $TOKEN" \
  -H "X-Vault-Namespace: $NAMESPACE" \
  -d '{
  "policy": "path \"kv/test/data/demo\" {capabilities =  [ \"read\" ]}"
}'

create_role=$(curl -X 'POST' \
  $VAULT_HOST/v1/auth/approle/role/test_role \
  -H "accept: */*" \
  -H "Content-Type: application/json" \
  -H "X-Vault-Token: $TOKEN" \
  -H "X-Vault-Namespace: $NAMESPACE" \
  -d '{  "token_policies": ["test_policy"],"token_ttl": "4h","token_max_ttl":"4h"
}')

echo "response for create role : $create_role"

read_role=$(curl -X 'GET' \
  $VAULT_HOST/v1/auth/approle/role/test_role/role-id \
  -H "accept: */*" \
  -H "X-Vault-Token: $TOKEN" \
  -H "X-Vault-Namespace: $NAMESPACE")

  echo "response of read role: $read_role"

create_secret_id=$(curl -X 'POST' \
  $VAULT_HOST/v1/auth/approle/role/test_role/secret-id \
  -H "accept: */*" \
  -H "Content-Type: application/json" \
  -H "X-Vault-Token: $TOKEN" \
  -H "X-Vault-Namespace: $NAMESPACE" \
  -d '{}')

  echo $create_secret_id