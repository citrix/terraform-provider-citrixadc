---
subcategory: "SSL"
---

# Data Source: sslservicegroup_sslciphersuite_binding

The sslservicegroup_sslciphersuite_binding data source allows you to retrieve information about an SSL cipher suite binding to an SSL service group.

## Example Usage

```terraform
data "citrixadc_sslservicegroup_sslciphersuite_binding" "tf_binding" {
  servicegroupname = "my_gslbvservicegroup"
  ciphername       = "my_ciphersuite"
}

output "description" {
  value = data.citrixadc_sslservicegroup_sslciphersuite_binding.tf_binding.description
}
```

## Argument Reference

* `servicegroupname` - (Required) The name of the SSL service to which the SSL policy needs to be bound.
* `ciphername` - (Required) The name of the cipher group/alias/name configured for the SSL service group.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `description` - The description of the cipher.
* `id` - The id of the sslservicegroup_sslciphersuite_binding. It is the concatenation of `servicegroupname` and `ciphername` attributes separated by a comma.
