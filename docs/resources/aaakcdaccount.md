---
subcategory: "AAA"
---

# Resource: aaakcdaccount

The aaakcdaccount resource is used to create aaakcdaccount.


## Example usage

```hcl
resource "citrixadc_aaakcdaccount" "tf_aaakcdaccount" {
  kcdaccount    = "my_kcdaccount"
  delegateduser = "john"
  kcdpassword   = "my_password"
  realmstr      = "my_realm"
}
```


## Argument Reference

* `kcdaccount` - (Required) The name of the KCD account. Minimum length =  1
* `keytab` - (Optional) The path to the keytab file. If specified other parameters in this command need not be given.
* `realmstr` - (Optional) Kerberos Realm.
* `delegateduser` - (Optional) Username that can perform kerberos constrained delegation.
* `kcdpassword` - (Optional) Password for Delegated User.
* `usercert` - (Optional) SSL Cert (including private key) for Delegated User.
* `cacert` - (Optional) CA Cert for UserCert or when doing PKINIT backchannel.
* `userrealm` - (Optional) Realm of the user.
* `enterpriserealm` - (Optional) Enterprise Realm of the user. This should be given only in certain KDC deployments where KDC expects Enterprise username instead of Principal Name.
* `servicespn` - (Optional) Service SPN. When specified, this will be used to fetch kerberos tickets. If not specified, Citrix ADC will construct SPN using service fqdn.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaakcdaccount. It has the same value as the `kcdaccount` attribute.
