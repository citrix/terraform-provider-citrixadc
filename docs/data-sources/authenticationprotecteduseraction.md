---
subcategory: "Authentication"
---

# Data Source: authenticationprotecteduseraction

The authenticationprotecteduseraction data source allows you to retrieve information about an authentication protected-user action configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_authenticationprotecteduseraction" "tf_protecteduseraction" {
  name = "my_protecteduseraction"
}

output "realmstr" {
  value = data.citrixadc_authenticationprotecteduseraction.tf_protecteduseraction.realmstr
}
```


## Argument Reference

* `name` - (Required) Name of the action to configure.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `realmstr` - Kerberos Realm.
* `maxconcurrentusers` - Max number of concurrent users allowed.
* `id` - The id of the authenticationprotecteduseraction. It has the same value as the `name` attribute.
