Fleeting
========

This is a prototype of a tool to generate an ephemeral OpenShift installation
image.

Testing
-------

Put the desired ignition file at `data/ignition/test_ignition.ign`.

Run the tool using `go run cmd/main.go`.

The output ISO is written to `output/fleeting.iso`.
