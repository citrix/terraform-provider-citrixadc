---
subcategory: "Network"
---

# Data Source: rnat_nsip_binding

The rnat_nsip_binding data source allows you to retrieve information about a rnat_nsip_binding.


## Example Usage

```terraform
data "citrixadc_rnat_nsip_binding" "tf_rnat_nsip_binding" {
  name  = "my_rnat"
  natip = "10.222.74.200"
}

output "name" {
  value = data.citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding.name
}

output "natip" {
  value = data.citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding.natip
}
```


## Argument Reference

* `name` - (Required) Name of the RNAT rule to which to bind NAT IPs.
* `natip` - (Required) Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnat_nsip_binding. It is a system-generated identifier.

