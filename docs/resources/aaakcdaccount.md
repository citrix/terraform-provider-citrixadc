---
subcategory: "AAA"
---

# Resource: aaakcdaccount

The aaakcdaccount resource is used to create and manage AAA KCD (Kerberos Constrained Delegation) accounts on Citrix ADC.


## Example Usage

### Using kcdpassword (sensitive attribute - persisted in state)

```hcl
variable "aaakcdaccount_kcdpassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaakcdaccount" "tf_aaakcdaccount" {
  kcdaccount    = "my_kcdaccount"
  delegateduser = "john"
  kcdpassword   = var.aaakcdaccount_kcdpassword
  realmstr      = "EXAMPLE.COM"
}
```

### Using kcdpassword_wo (write-only/ephemeral - NOT persisted in state)

The `kcdpassword_wo` attribute provides an ephemeral path for the delegated user password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `kcdpassword_wo_version`.

```hcl
variable "aaakcdaccount_kcdpassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaakcdaccount" "tf_aaakcdaccount" {
  kcdaccount            = "my_kcdaccount"
  delegateduser         = "john"
  kcdpassword_wo        = var.aaakcdaccount_kcdpassword
  kcdpassword_wo_version = 1
  realmstr              = "EXAMPLE.COM"
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_aaakcdaccount" "tf_aaakcdaccount" {
  kcdaccount            = "my_kcdaccount"
  delegateduser         = "john"
  kcdpassword_wo        = var.aaakcdaccount_kcdpassword
  kcdpassword_wo_version = 2  # Bumped to trigger update
  realmstr              = "EXAMPLE.COM"
}
```


## Argument Reference

* `kcdaccount` - (Required) The name of the KCD account.
* `cacert` - (Optional) CA Cert for UserCert or when doing PKINIT backchannel.
* `delegateduser` - (Optional) Username that can perform kerberos constrained delegation.
* `enterpriserealm` - (Optional) Enterprise Realm of the user. This should be given only in certain KDC deployments where KDC expects Enterprise username instead of Principal Name.
* `kcdpassword` - (Optional, Sensitive) Password for Delegated User. The value is persisted in Terraform state (encrypted). See also `kcdpassword_wo` for an ephemeral alternative.
* `kcdpassword_wo` - (Optional, Sensitive, WriteOnly) Same as `kcdpassword`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `kcdpassword_wo_version`. If both `kcdpassword` and `kcdpassword_wo` are set, `kcdpassword_wo` takes precedence.
* `kcdpassword_wo_version` - (Optional) An integer version tracker for `kcdpassword_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `keytab` - (Optional) The path to the keytab file. If specified other parameters in this command need not be given.
* `realmstr` - (Optional) Kerberos Realm.
* `servicespn` - (Optional) Service SPN. When specified, this will be used to fetch kerberos tickets. If not specified, Citrix ADC will construct SPN using service fqdn.
* `usercert` - (Optional) SSL Cert (including private key) for Delegated User.
* `userrealm` - (Optional) Realm of the user.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaakcdaccount. It has the same value as the `kcdaccount` attribute.


## Import

A aaakcdaccount can be imported using its kcdaccount, e.g.

```shell
terraform import citrixadc_aaakcdaccount.tf_aaakcdaccount my_kcdaccount
```
