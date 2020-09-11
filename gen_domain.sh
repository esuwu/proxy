#!/bin/bash
cd certs
openssl genrsa -out "$1".key 2048
openssl req -new -key "$1".key -out "$1".csr
#In answer to question `Common Name (e.g. server FQDN or YOUR name) []:` you should set `secure.domain.com` (your real domain name)
openssl x509 -req -in "$1".csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -days 365 -out "$1".crt
