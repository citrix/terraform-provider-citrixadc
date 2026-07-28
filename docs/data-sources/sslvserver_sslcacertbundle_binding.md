---
subcategory: "SSL"
---

# Data Source: sslvserver_sslcacertbundle_binding

The sslvserver_sslcacertbundle_binding data source allows you to retrieve information about a CA certificate bundle binding to an SSL virtual server.

## Example usage

```terraform
data "citrixadc_sslvserver_sslcacertbundle_binding" "tf_binding" {
  vservername      = "tf_lbvserver"
  cacertbundlename = "tf_cacertbundle"
}

output "vservername" {
  value = data.citrixadc_sslvserver_sslcacertbundle_binding.tf_binding.vservername
}

output "cacertbundlename" {
  value = data.citrixadc_sslvserver_sslcacertbundle_binding.tf_binding.cacertbundlename
}
```

## Argument Reference

The following arguments are required:

* `vservername` - (Required) Name of the SSL virtual server.
* `cacertbundlename` - (Required) CA certbundle name bound to the vserver.

The following arguments are optional:

* `skipcacertbundle` - (Optional) The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslvserver_sslcacertbundle_binding. It is the concatenation of the `vservername` and `cacertbundlename` attributes separated by a comma.
