Fleeting
========

This is a prototype of a tool to generate an ephemeral OpenShift installation
image.

Testing
-------

Put files to be populated by Ignition in the tree rooted at
`data/ignition/files/`. Put systemd units in `data/ignition/systemd/units`.

The ZTP manifests provided by the users are read from `./manifests`.
The required manifests are:
* manifests/pull-secret.yaml
* manifests/cluster-deployment.yaml
* manifests/agent-cluster-install.yaml
* manifests/infraenv.yaml

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

Node0
-------

To run the assisted service only on a pre-determined node a.k.a. node0, currently we have hardcoded a static IP 192.168.122.2. 
A systemd service named `node-zero.service` looks for the static IP and if the current node's IP matches with it then only the `assisted-service.service` systemd service is startred. The `assisted-service.service` is responsible for running the assisted service.

To set the static ip in libvirt:
1. Edit the default network
```
virsh net-edit default
```
2. Add `dns` and `host mac` for node0

```
<network>
  <name>default</name>
  <uuid>2467ce0f-aff2-4031-b0be-3e40fca96421</uuid>
  <forward mode='nat'/>
  <bridge name='virbr0' stp='on' delay='0'/>
  <mac address='52:54:00:bf:2b:d4'/>
  <dns>
    <host ip='192.168.122.2'>
      <hostname>node0</hostname>
    </host>
  </dns>
  <ip address='192.168.122.1' netmask='255.255.255.0'>
    <dhcp>
      <range start='192.168.122.2' end='192.168.122.254'/>
      <host mac='52:54:00:aa:aa:aa' ip='192.168.122.2'/>
    </dhcp>
  </ip>
</network>
```
3. Destroy and start network
```
sudo virsh net-destroy default
sudo virsh net-start default
```

