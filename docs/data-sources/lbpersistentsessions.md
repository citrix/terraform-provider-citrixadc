---
subcategory: "Load Balancing"
---

# Data Source: lbpersistentsessions

The lbpersistentsessions data source retrieves information about an active load balancing persistence session on the Citrix ADC. It queries the full list of persistence sessions and returns the first one matching the optional `vserver` and `nodeid` filters, letting you inspect existing client-to-server affinity from Terraform.

Note: The data source reads via the NITRO get(all) endpoint. If there are no active persistence sessions on the ADC (or none match the supplied filters), the read returns an error.


## Example usage

```hcl
data "citrixadc_lbpersistentsessions" "tf_lbpersistentsessions" {
  vserver = "lbvserver1"
}

output "persistenceparameter" {
  value = data.citrixadc_lbpersistentsessions.tf_lbpersistentsessions.persistenceparameter
}
```


## Argument Reference

* `vserver` - (Optional) The name of the virtual server to filter persistence sessions by. If omitted, the first available persistence session is returned.
* `nodeid` - (Optional) Unique number that identifies the cluster node to filter persistence sessions by.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `persistenceparameter` - The persistence parameter of the matched session.
* `vserver` - The name of the virtual server for the matched session.
* `nodeid` - The cluster node that owns the matched session.
* `id` - A synthetic identifier with the fixed value `lbpersistentsessions-query`.
