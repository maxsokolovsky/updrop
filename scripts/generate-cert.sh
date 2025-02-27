#!/usr/bin/env bash

CERT_FILE="cert.pem"
KEY_FILE="key.pem"
DAYS=365
SUBJ="/C=US/ST=CA/L=Local/O=MyCompany/CN=localhost"

openssl genpkey -algorithm RSA -out "$KEY_FILE"
openssl req -new -x509 -key "$KEY_FILE" -out "$CERT_FILE" -days "$DAYS" -subj "$SUBJ"

echo "Self-signed certificate generated."
echo "Certificate: $CERT_FILE"
echo "Private Key: $KEY_FILE"
