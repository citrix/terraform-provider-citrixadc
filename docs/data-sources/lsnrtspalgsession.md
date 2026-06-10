---
subcategory: "LSN"
---

# Data Source: lsnrtspalgsession

The lsnrtspalgsession data source retrieves information about a Large Scale NAT (LSN) RTSP ALG (Application Layer Gateway) session on the Citrix ADC. Unlike the resource, it only reads session state and never flushes anything, letting you inspect an active RTSP ALG session from Terraform by its `sessionid` (optionally scoped to a cluster node).

Note: The data source reads via the NITRO get endpoint. If no RTSP ALG session matches the supplied `sessionid` (and optional `nodeid`), the read returns an error.


## Example usage

```hcl
data "citrixadc_lsnrtspalgsession" "tf_lsnrtspalgsession" {
  sessionid = "10.102.43.13:6789"
}

output "rtspalgsession_nodeid" {
  value = data.citrixadc_lsnrtspalgsession.tf_lsnrtspalgsession.nodeid
}
```


## Argument Reference

* `sessionid` - (Required) Session ID for the RTSP call. Identifies the RTSP ALG session to look up.
* `nodeid` - (Optional) Unique number that identifies the cluster node to scope the lookup to.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `sessionid` - The session ID of the matched RTSP ALG session.
* `nodeid` - The cluster node that owns the matched RTSP ALG session.
* `id` - A synthetic identifier for the data source read.
