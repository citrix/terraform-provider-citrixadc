#!/bin/bash

# Fail early
set -e

# Just making sure configuration packages don't have obvious errors
for configpack in $(find ../config -mindepth 1 -maxdepth 1 -type d)
do
echo "Will build $configpack"
go build $configpack
done
