---
subcategory: "GSLB"
---

# Resource: gslbldnsentry_delete

This resource is used to delete a learned LDNS entry from the Citrix ADC GSLB subsystem.


## Example usage

```hcl
resource "citrixadc_gslbldnsentry_delete" "tf_gslbldnsentry_delete" {
  ipaddress = "192.0.2.10"
}
```


## Argument Reference

* `ipaddress` - (Required) IP address of the LDNS server. Applying the resource removes the GSLB LDNS entry with this IP address. Changing this attribute forces a new resource to be created, which removes the entry for the new IP address.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the gslbldnsentry_delete resource. It is set to `gslbldnsentry_delete`.
