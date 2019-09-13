#!/usr/bin/bash

set -a
SSH_PORT=32769

openssl req  -newkey rsa:2048 -nodes -keyout server1.key  -x509 -days 365 -out server.crt
openssl rsa -in server1.key -out server.key
ssh -p $SSH_PORT nsroot@localhost 'mkdir -p /var/certs'
scp -P $SSH_PORT server.* nsroot@localhost:/var/certs/
rm server1.key
rm server.*
