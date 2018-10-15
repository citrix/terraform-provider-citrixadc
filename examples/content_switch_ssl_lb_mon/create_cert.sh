#!/usr/bin/bash

set -a
SSH_PORT=32770

openssl req  -newkey rsa:2048 -nodes -keyout server1.key  -x509 -days 365 -out server.crt
openssl rsa -in server1.key -out server.key
ssh -p $SSH_PORT root@localhost 'mkdir -p /var/certs'
scp -P $SSH_PORT server.* root@localhost:/var/certs/
rm server1.key
rm server.*
