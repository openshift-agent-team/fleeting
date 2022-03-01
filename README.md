Fleeting
========

This is a prototype of a tool to generate an ephemeral OpenShift installation
image.

Testing
-------

Put files to be populated by Ignition in the tree rooted at
`data/ignition/files/`. Put systemd units in `data/ignition/systemd/units`.

The agent.service file requires SERVICE_BASE_URL and INFRA_ENV_ID to be set.
The pull secret is also required and is written to /root/.docker/config.json in
the ISO.
For now, these are set through environment variables.

```shell
export PULL_SECRET=$(cat ~/Downloads/pull-secret.txt)
export SERVICE_BASE_URL=http://10.0.1.10:6000
export INFRA_ENV_ID=60947297-c9a1-49ac-8119-d9656a244c83
```

Run the tool using `go run cmd/main.go`.

The output ISO is written to `output/fleeting.iso`.

Boot the ISO in a VM with at least 4096MiB of RAM. No storage is required.
The assisted-service UI is available on port 8080.
