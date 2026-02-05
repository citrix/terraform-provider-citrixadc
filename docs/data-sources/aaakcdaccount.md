---
subcategory: "AAA"
---

# Data Source `aaakcdaccount`

The aaakcdaccount data source allows you to retrieve information about AAA KCD (Kerberos Constrained Delegation) accounts.


## Example usage

```terraform
data "citrixadc_aaakcdaccount" "tf_aaakcdaccount" {
  kcdaccount = "my_kcdaccount"
}

output "delegateduser" {
  value = data.citrixadc_aaakcdaccount.tf_aaakcdaccount.delegateduser
}

output "realmstr" {
  value = data.citrixadc_aaakcdaccount.tf_aaakcdaccount.realmstr
}
```


## Argument Reference

* `kcdaccount` - (Required) The name of the KCD account.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `delegateduser` - Username that can perform kerberos constrained delegation.
* `kcdpassword` - Password for the delegated user.
* `realmstr` - Kerberos realm of the delegated user.
* `servicespn` - Service SPN (Service Principal Name) of the application.
* `usercert` - SSL certificate to use for the KCD account.
* `userrealm` - Realm of the user.
* `enterpriserealm` - Enterprise realm of the user. This value is used only when the realm based KCD is configured.
* `keytab` - Keytab file to use for the KCD account.
* `cacert` - CA certificate to verify the KDC.

## Attribute Reference

* `id` - The id of the aaakcdaccount. It has the same value as the `kcdaccount` attribute.


## Import

A aaakcdaccount can be imported using its kcdaccount, e.g.

```shell
terraform import citrixadc_aaakcdaccount.tf_aaakcdaccount my_kcdaccount
```
