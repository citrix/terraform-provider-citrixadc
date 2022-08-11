---
subcategory: "Network"
---

# Resource: rnat_nsip_binding

The rnat_nsip_binding resource is used to create rnat_nsip_binding.


## Example usage

```hcl
resource "citrixadc_rnat_nsip_binding" "tf_rnat_nsip_binding" {
  name  = "my_rnat"
  natip = "10.222.74.200"
}

```


## Argument Reference

* `natip` - (Required) Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range. Minimum length =  1
* `name` - (Required) Name of the RNAT rule to which to bind NAT IPs. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rnat_nsip_binding. It has the same value as the `name` and `natip` attributes separated by a comma.


## Import

A rnat_nsip_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_rnat_nsip_binding.tf_rnat_nsip_binding my_rnat,10.222.74.200
```
