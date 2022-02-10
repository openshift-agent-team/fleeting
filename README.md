Fleeting
========

This is a prototype of a tool to generate an ephemeral OpenShift installation
image.

Testing
-------

Put files to be populated by Ignition in the tree rooted at
`data/ignition/files/`. Put systemd units in `data/ignition/systemd/units`.

Run the tool using `go run cmd/main.go`.

The output ISO is written to `output/fleeting.iso`.
