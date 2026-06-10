---
subcategory: "GSLB"
---

# Data Source: gslbldnsentries

The gslbldnsentries data source allows you to retrieve information about the learned GSLB local DNS (LDNS) round-trip time (RTT) entries on the Citrix ADC. These entries record the RTT the appliance measured to the local DNS servers that resolve client queries, which drives the dynamic RTT-based GSLB load balancing method.

The data source reads the entries via the NITRO get(all) endpoint and returns the first entry that matches the supplied filter. Note: the Read errors out when the appliance has no LDNS entries (for example, immediately after the `citrixadc_gslbldnsentries` resource has cleared them, or before any have been learned).


## Example usage

```terraform
# Return the first learned LDNS entry on the appliance.
data "citrixadc_gslbldnsentries" "tf_gslbldnsentries" {
}

output "ldns_nodeid" {
  value = data.citrixadc_gslbldnsentries.tf_gslbldnsentries.nodeid
}
```

To filter the lookup to a specific cluster node:

```terraform
data "citrixadc_gslbldnsentries" "tf_gslbldnsentries" {
  nodeid = 1
}
```


## Argument Reference

* `nodeid` - (Optional) Unique number that identifies the cluster node. When set, the data source returns the first LDNS entry whose `nodeid` matches this value.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - A synthetic identifier for the query, with the constant value `gslbldnsentries-query`.
* `nodeid` - Unique number that identifies the cluster node of the matched LDNS entry.
