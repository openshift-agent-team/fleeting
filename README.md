Fleeting
========

**Development has moved to the [OpenShift Installer](https://github.com/openshift/installer/tree/agent-installer).**

Fleeting was a prototype of the agent-based installer - a tool to generate an
ephemeral OpenShift installation image. Development has moved to the
`agent-installer` branch of the
[openshift/installer](https://github.com/openshift/installer/tree/agent-installer)
repo, and will eventually merge into the main installer.

Testing
-------

Put files to be populated by Ignition in the tree rooted at
`data/ignition/files/`. Put systemd units in `data/ignition/systemd/units`.
These are all built into the binary at compilation time.

The ZTP manifests provided by the users are read from `./manifests`.
The required sample manifests are:
* [pull-secret.yaml](./manifests/samples/pull-secret.yaml)
   - Your pull secret JSON data can be [obtained from the OpenShift Console](https://console.redhat.com/openshift/install/pull-secret).
* [cluster-deployment.yaml](./manifests/samples/cluster-deployment.yaml)
* [agent-cluster-install.yaml](./manifests/samples/agent-cluster-install.yaml)
* [infraenv.yaml](./manifests/samples/infraenv.yaml)
* [nmstateconfig.yaml](./manifests/samples/nmstateconfig.yaml)
 <br>Note: Atleast one NmStateConfig should be defined which is used to identify and set node0. At the max, the cluster can have one NmStateConfigs per node.
 - To create a 3 masters compact cluster, set `spec.provisionRequirements.controlPlaneAgents` to `3` in AgentClusterInstall.
 - To create additional workers, set `spec.provisionRequirements.workerAgents` in AgentClusterInstall.
 - To create an SNO cluster, set `spec.provisionRequirements.controlPlaneAgents` to `1` in AgentClusterInstall and no worker agents. Also, do not set `spec.ingressVIP` in AgentClusterInstall.

Run the tool using `make iso`.

The output ISO is written to `output/fleeting.iso`.

Boot the ISO in a VM with at least 8192MiB of RAM. No storage is required.
The assisted-service UI is available on port 8080.

Node0
-------

To run the assisted service only on a pre-determined node a.k.a. node0, the node0 IP must be defined in nmstateconfig.yaml with a mac address that matches that node, for example:

```
apiVersion: agent-install.openshift.io/v1beta1
kind: NMStateConfig
metadata:
  name: mynmstateconfig
  namespace: spoke-cluster
  labels:
    cluster0-nmstate-label-name: cluster0-nmstate-label-value
spec:
  config:
    interfaces:
      - name: eth0
        type: ethernet
        state: up
        mac-address: 52:54:01:aa:aa:a1
        ipv4:
          enabled: true
          address:
            - ip: 192.168.122.2
              prefix-length: 24
          dhcp: false
    dns-resolver:
      config:
        server:
          - 192.168.122.1
    routes:
      config:
        - destination: 0.0.0.0/0
          next-hop-address: 192.168.122.1
          next-hop-interface: eth0
          table-id: 254
  interfaces:
    - name: "eth0"
      macAddress: "52:54:01:aa:aa:a1"
```

A systemd service named `node-zero.service` looks for the static IP and if the current node's IP matches with it then only the `assisted-service.service` systemd service is started. The `assisted-service.service` is responsible for running the assisted service. In order to supply a hostname to this node a DNS entry must be added as its currently not possible to configure the hostname in a manifest file.

For the other nodes, besides node0, the IP addresses can either be statically defined in nmstateconfig.yaml (recommended), or added as DHCP entries.

To add the IP addresses to nmstateconfig.yaml add another entry for `kind: NMStateConfig` similar to node0 separated by the YAML separator `---`. DNS entries must also be added for these additional nodes.
