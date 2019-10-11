#!/bin/bash -x

# Create a full CA chain with self signed CA root
# Intermediate CA
# and 3 client certificates
set -e

for C in `echo root-ca intermediate out`; do

  if [[ -d $C ]]; then
    echo "Deleting directory $C"
    rm -rf $C
  fi
done


for C in `echo root-ca intermediate`; do

  mkdir $C
  cd $C
  mkdir certs crl newcerts private
  cd ..

  echo 1000 > $C/serial
  touch $C/index.txt $C/index.txt.attr

  echo '
[ ca ]
default_ca = CA_default
[ CA_default ]
dir            = '$C'                     # Where everything is kept
certs          = $dir/certs               # Where the issued certs are kept
crl_dir        = $dir/crl                 # Where the issued crl are kept
database       = $dir/index.txt           # database index file.
new_certs_dir  = $dir/newcerts            # default place for new certs.
certificate    = $dir/cacert.pem          # The CA certificate
serial         = $dir/serial              # The current serial number
crl            = $dir/crl.pem             # The current CRL
private_key    = $dir/private/ca.key.pem  # The private key
RANDFILE       = $dir/.rnd                # private random number file
nameopt        = default_ca
certopt        = default_ca
policy         = policy_match
default_days   = 365
default_md     = sha256

[ policy_match ]
countryName            = optional
stateOrProvinceName    = optional
organizationName       = optional
organizationalUnitName = optional
commonName             = supplied
emailAddress           = optional

[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name

[req_distinguished_name]

[v3_req]
basicConstraints = CA:TRUE
' > $C/openssl.conf
done

openssl genrsa -out root-ca/private/ca.key 2048
openssl req -config root-ca/openssl.conf -new -x509 -days 3650 -key root-ca/private/ca.key -sha256 -extensions v3_req -out root-ca/certs/ca.crt -subj '/CN=Root-ca'

# Copy root ca certificate to top level testdata directory
cp root-ca/certs/ca.crt ../../

openssl genrsa -out intermediate/private/intermediate.key 2048
openssl req -config intermediate/openssl.conf -sha256 -new -key intermediate/private/intermediate.key -out intermediate/certs/intermediate.csr -subj '/CN=Interm.'
openssl ca -batch -config root-ca/openssl.conf -keyfile root-ca/private/ca.key -cert root-ca/certs/ca.crt -extensions v3_req -notext -md sha256 -in intermediate/certs/intermediate.csr -out intermediate/certs/intermediate.crt

# Copy intermediate certificate to top level testdata directory
cp intermediate/certs/intermediate.crt ../../

mkdir out

for I in `seq 1 3` ; do
  openssl genrsa -out out/key${I}.pem 2048
  openssl req -new -out out/$I.request -days 365 -key out/key${I}.pem -nodes -subj "/CN=$I.example.com"
  openssl ca -batch -config root-ca/openssl.conf -keyfile intermediate/private/intermediate.key -cert intermediate/certs/intermediate.crt -out out/certificate${I}.crt -infiles out/$I.request

  # Copy to root of testdata
  cp out/key${I}.pem ../../
  cp out/certificate${I}.crt ../../
done

echo All Good!
