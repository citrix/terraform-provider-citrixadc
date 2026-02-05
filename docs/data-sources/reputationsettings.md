---
subcategory: "Reputation"
---

# Data Source `reputationsettings`

The reputationsettings data source allows you to retrieve information about Reputation settings configuration.


## Example usage

```terraform
data "citrixadc_reputationsettings" "tf_reputationsettings" {
}

output "proxyserver" {
  value = data.citrixadc_reputationsettings.tf_reputationsettings.proxyserver
}

output "proxyport" {
  value = data.citrixadc_reputationsettings.tf_reputationsettings.proxyport
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `proxypassword` - Password with which user logs on.
* `proxyport` - Proxy server port.
* `proxyserver` - Proxy server IP to get Reputation data.
* `proxyusername` - Proxy Username
