---
subcategory: "SSL"
---

# Data Source: sslservicegroup_sslcacertbundle_binding

The sslservicegroup_sslcacertbundle_binding data source allows you to retrieve information about a CA certificate bundle binding on an SSL service group.

## Example usage

```terraform
data "citrixadc_sslservicegroup_sslcacertbundle_binding" "tf_binding" {
  servicegroupname = "tf_servicegroup"
  cacertbundlename = "tf_cacertbundle"
}

output "cacertbundlename" {
  value = data.citrixadc_sslservicegroup_sslcacertbundle_binding.tf_binding.cacertbundlename
}
```

## Argument Reference

* `servicegroupname` - (Required) The name of the SSL service group to which the CA certificate bundle is bound.
* `cacertbundlename` - (Required) The name of the CA certificate bundle bound to the service group.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservicegroup_sslcacertbundle_binding. It is the concatenation of the `servicegroupname` and `cacertbundlename` attributes separated by a comma.
