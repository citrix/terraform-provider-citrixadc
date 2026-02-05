---
subcategory: "Application"
---

# Data Source `appalgparam`

The appalgparam data source allows you to retrieve information about Application ALG (Application Level Gateway) parameters configuration.


## Example usage

```terraform
data "citrixadc_appalgparam" "tf_appalgparam" {
}

output "pptpgreidletimeout" {
  value = data.citrixadc_appalgparam.tf_appalgparam.pptpgreidletimeout
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `pptpgreidletimeout` - Interval in sec, after which data sessions of PPTP GRE is cleared.

## Attribute Reference

* `id` - The id of the appalgparam. It is a system-generated identifier.
