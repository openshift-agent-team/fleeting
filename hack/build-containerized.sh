#!/bin/sh

set -ex

podman run --rm \
--volume "${PWD}:/go/src/github.com/openshift-agent-team:z"  \
--workdir /go/src/github.com/openshift-agent-team \
registry.ci.openshift.org/openshift/release:golang-1.17 \
sh hack/build.sh