#!/bin/bash
outputdir=${1:-../config}
mkdir -p $outputdir
for jsonf in $(find ../jsonconfig -name \*.json)
do
    folder=$(basename $(dirname $jsonf))
    name=$(basename $jsonf .json)
    mkdir -p $outputdir/$folder
    echo $outputdir/$folder/$name.go
    $PWD/schema-generate -i $jsonf -o $outputdir/$folder/$name.go -p $folder
    gofmt -w $outputdir/$folder/$name.go
done
