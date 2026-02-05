---
subcategory: "Authentication"
---

# Data Source `authenticationpushservice`

The authenticationpushservice data source allows you to retrieve information about authentication push services.


## Example usage

```terraform
data "citrixadc_authenticationpushservice" "tf_pushservice" {
  name = "example_pushservice"
}

output "customerid" {
  value = data.citrixadc_authenticationpushservice.tf_pushservice.customerid
}

output "refreshinterval" {
  value = data.citrixadc_authenticationpushservice.tf_pushservice.refreshinterval
}
```


## Argument Reference

* `name` - (Required) Name for the push service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `clientid` - Unique identity for communicating with Citrix Push server in cloud.
* `clientsecret` - Unique secret for communicating with Citrix Push server in cloud.
* `customerid` - Customer id/name of the account in cloud that is used to create clientid/secret pair.
* `refreshinterval` - Interval at which certificates or idtoken is refreshed.
* `id` - The id of the authenticationpushservice. It has the same value as the `name` attribute.


## Import

A authenticationpushservice can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationpushservice.tf_pushservice example_pushservice
```
