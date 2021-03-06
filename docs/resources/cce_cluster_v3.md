---
subcategory: "Cloud Container Engine (CCE)"
---

# opentelekomcloud_cce_cluster_v3

Provides a cluster resource management.

## Example Usage

```hcl
variable "flavor_id" { }
variable "vpc_id" { }
variable "subnet_id" { }

resource "opentelekomcloud_cce_cluster_v3" "cluster_1" {
  name        = "cluster"
  description = "Create cluster"

  cluster_type           = "VirtualMachine"
  flavor_id              = var.flavor_id
  vpc_id                 = var.vpc_id
  subnet_id              = var.subnet_id
  container_network_type = "overlay_l2"
  authentication_mode    = "rbac"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Cluster name. Changing this parameter will create a new cluster resource.

* `labels` - (Optional) Cluster tag, key/value pair format. Changing this parameter will create a new cluster resource.

* `annotations` - (Optional) Cluster annotation, key/value pair format. Changing this parameter will create a new cluster resource.

* `flavor_id` - (Required) Cluster specifications. Changing this parameter will create a new cluster resource.
  * `cce.s1.small` - small-scale single cluster (up to 50 nodes).
  * `cce.s1.medium` - medium-scale single cluster (up to 200 nodes).
  * `cce.s1.large` - large-scale single cluster (up to 1000 nodes).
  * `cce.s2.small` - small-scale HA cluster (up to 50 nodes).
  * `cce.s2.medium` - medium-scale HA cluster (up to 200 nodes).
  * `cce.s2.large` - large-scale HA cluster (up to 1000 nodes).
  * `cce.t1.small` - small-scale single physical machine cluster (up to 10 nodes).
  * `cce.t1.medium` - medium-scale single physical machine cluster (up to 100 nodes).
  * `cce.t1.large` - large-scale single physical machine cluster (up to 500 nodes).
  * `cce.t2.small` - small-scale HA physical machine cluster (up to 10 nodes).
  * `cce.t2.medium` - medium-scale HA physical machine cluster (up to 100 nodes).
  * `cce.t2.large` - large-scale HA physical machine cluster (up to 500 nodes).

* `cluster_version` - (Optional) For the cluster version, possible values are `v1.13.10-r0`, `v1.15.6-r1`.
  Changing this parameter will create a new cluster resource. [OTC-API](https://docs.otc.t-systems.com/en-us/api2/cce/cce_02_0236.html)

* `cluster_type` - (Required) Cluster Type, possible values are `VirtualMachine` and `BareMetal`. Changing this parameter will create a new cluster resource.

* `description` - (Optional) Cluster description.

* `billing_mode` - (Optional) Charging mode of the cluster, which is 0 (on demand). Changing this parameter will create a new cluster resource.

* `extend_param` - (Optional) Extended parameter. Changing this parameter will create a new cluster resource.

* `vpc_id` - (Required) The ID of the VPC used to create the node. Changing this parameter will create a new cluster resource.

* `subnet_id` - (Required) The Network ID of the subnet used to create the node. Changing this parameter will create a new cluster resource.

* `highway_subnet_id` - (Optional) The ID of the high speed network used to create bare metal nodes. Changing this parameter will create a new cluster resource.

* `container_network_type` - (Required) Container network type.
  * `overlay_l2` - An overlay_l2 network built for containers by using Open vSwitch(OVS)
  * `underlay_ipvlan` - An underlay_ipvlan network built for bare metal servers by using ipvlan.
  * `vpc-router` - An vpc-router network built for containers by using ipvlan and custom VPC routes.

* `container_network_cidr` - (Optional) Container network segment. Changing this parameter will create a new cluster resource.

* `authentication_mode` - (Optional) Authentication mode of the cluster, possible values are `rbac` and `authenticating_proxy`.
  Defaults to `rbac`. Changing this parameter will create a new cluster resource.

* `multi_az` - (Optional) Enable multiple AZs for the cluster, only when using HA flavors. Changing this parameter will create a new cluster resource.

* `eip` - (Optional) EIP address of the cluster.

* `kubernetes_svc_ip_range` - (Optional) Service CIDR block, or the IP address range which the kubernetes
  clusterIp must fall within. This parameter is available only for clusters of v1.11.7 and later.

## Attributes Reference

All above argument parameters can be exported as attribute parameters along with attribute reference.

* `id` - ID of the cluster resource.

* `status` - Cluster status information.

* `internal` - The internal network address.

* `external` - The external network address.

* `external_otc` - The endpoint of the cluster to be accessed through API Gateway.

* `certificate_clusters/name` - The cluster name.

* `certificate_clusters/server` - The server IP address.

* `certificate_clusters/certificate_authority_data` - The certificate data.

* `certificate_users/name` - The user name.

* `certificate_users/client_certificate_data` - The client certificate data.

* `certificate_users/client_key_data` - The client key data.

* `kube_proxy_mode` - (Optional) Service forwarding mode. Two modes are available:
  * `iptables`: Traditional kube-proxy uses iptables rules to implement service load balancing.
  In this mode, too many iptables rules will be generated when many services are deployed.
  In addition, non-incremental updates will cause a latency and even obvious performance issues
  in the case of heavy service traffic.
  * `ipvs`: Optimized kube-proxy mode with higher throughput and faster speed.
  This mode supports incremental updates and can keep connections uninterrupted during service updates.
  It is suitable for large-sized clusters.

## Import

Cluster can be imported using the cluster id, e.g.

```sh
terraform import opentelekomcloud_cce_cluster_v3.cluster_1 4779ab1c-7c1a-44b1-a02e-93dfc361b32d
```
