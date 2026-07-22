---
subcategory: "SSL"
---

# Data Source: sslservice_sslcacertbundle_binding

The sslservice_sslcacertbundle_binding data source allows you to retrieve information about the binding between an SSL service and a CA certificate bundle.

## Example usage

```terraform
data "citrixadc_sslservice_sslcacertbundle_binding" "example" {
  servicename      = "tf_service"
  cacertbundlename = "tf_cacertbundle"
}

output "skipcacertbundle" {
  value = data.citrixadc_sslservice_sslcacertbundle_binding.example.skipcacertbundle
}
```

## Argument Reference

* `servicename` - (Required) Name of the SSL service for which to set advanced configuration.
* `cacertbundlename` - (Required) CA certbundle name bound to the service.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslservice_sslcacertbundle_binding.
* `skipcacertbundle` - The flag is used to indicate whether all CA_names in this particular CA certificate bundle needs to be sent to the SSL client while requesting for client certificate in a SSL handshake.
