---
subcategory: "LSN"
---

# Data Source: lsnsipalgcall

The lsnsipalgcall data source retrieves information about a Large Scale NAT (LSN) SIP ALG (Application Layer Gateway) call on the Citrix ADC. Unlike the resource, it only reads call state and never flushes anything, letting you inspect an active SIP ALG call from Terraform by its `callid` (optionally scoped to a cluster node).

Note: The data source reads via the NITRO get endpoint. If no SIP ALG call matches the supplied `callid` (and optional `nodeid`), the read returns an error.


## Example usage

```hcl
data "citrixadc_lsnsipalgcall" "tf_lsnsipalgcall" {
  callid = "12345-abcde"
}

output "sipalgcall_nodeid" {
  value = data.citrixadc_lsnsipalgcall.tf_lsnsipalgcall.nodeid
}
```


## Argument Reference

* `callid` - (Required) Call ID for the SIP call. Identifies the SIP ALG call to look up.
* `nodeid` - (Optional) Unique number that identifies the cluster node to scope the lookup to.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `callid` - The call ID of the matched SIP ALG call.
* `nodeid` - The cluster node that owns the matched SIP ALG call.
* `id` - A synthetic identifier for the data source read, equal to the matched `callid`.
