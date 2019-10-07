#!/bin/bash

# Check for schema-generate binary
which schema-generate
if [[ $? -ne 0 ]]; then
echo Check https://github.com/giorgos-nikolopoulos/generate
echo '`go get -u github.com/giorgos-nikolopoulos/generate/...` will probably work'
exit 1
fi

outputdir=${1:-../config}
mkdir -p $outputdir
for jsonf in $(find ../jsonconfig -name \*.json)
do
    folder=$(basename $(dirname $jsonf))
    name=$(basename $jsonf .json)
    mkdir -p $outputdir/$folder
    echo $outputdir/$folder/$name.go
    schema-generate -i $jsonf -o $outputdir/$folder/$name.go -p $folder -idPrefix http://citrix.com
    gofmt -w $outputdir/$folder/$name.go
done
