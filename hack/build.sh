#!/bin/sh

set -ex

version() { IFS="."; printf "%03d%03d%03d\\n" $@; unset IFS;}

minimum_go_version="1.17"
readonly minimum_go_version
current_go_version=$(go version | cut -d " " -f 3)
readonly current_go_version

if [ "$(version "${current_go_version#go}")" -lt "$(version "$minimum_go_version")" ]; then
     echo "Go version should be greater or equal to $minimum_go_version"
     exit 1
fi

go build -o bin/fleeting