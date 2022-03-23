---
subcategory: "Authentication"
---

# Resource: authenticationpushservice

The authenticationpushservice resource is used to create authentication pushservice resource.


## Example usage

```hcl
resource "citrixadc_authenticationpushservice" "tf_pushservice" {
  name            = "tf_pushservice"
  clientid        = "cliId"
  clientsecret    = "secret"
  customerid      = "cusID"
  refreshinterval = 50
}
```


## Argument Reference

* `name` - (Required) Name for the push service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created. 	    CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my push service" or 'my push service').
* `clientid` - (Optional) Unique identity for communicating with Citrix Push server in cloud.
* `clientsecret` - (Optional) Unique secret for communicating with Citrix Push server in cloud.
* `customerid` - (Optional) Customer id/name of the account in cloud that is used to create clientid/secret pair.
* `refreshinterval` - (Optional) Interval at which certificates or idtoken is refreshed.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationpushservice. It has the same value as the `name` attribute.


## Import

A authenticationpushservice can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationpushservice.tf_pushservice tf_pushservice
```
