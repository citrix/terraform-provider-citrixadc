---
subcategory: "Authentication"
---

# Resource: authenticationadfsproxyprofile

The authenticationadfsproxyprofile resource is used to create an ADFS proxy profile on the Citrix ADC. The profile lets the ADC act as an ADFS proxy and register a trust with the ADFS server.


## Example usage

### Using password (sensitive attribute - persisted in state)

```hcl
variable "authenticationadfsproxyprofile_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationadfsproxyprofile" "tf_adfsproxyprofile" {
  name        = "tf_adfsproxyprofile"
  certkeyname = "servercert1"
  serverurl   = "https://adfs.example.com"
  username    = "adfsuser"
  password    = var.authenticationadfsproxyprofile_password
}
```

### Using password_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the ADFS account password. The value is sent to the Citrix ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `password_wo_version`.

```hcl
variable "authenticationadfsproxyprofile_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationadfsproxyprofile" "tf_adfsproxyprofile" {
  name                = "tf_adfsproxyprofile"
  certkeyname         = "servercert1"
  serverurl           = "https://adfs.example.com"
  username            = "adfsuser"
  password_wo         = var.authenticationadfsproxyprofile_password
  password_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_authenticationadfsproxyprofile" "tf_adfsproxyprofile" {
  name                = "tf_adfsproxyprofile"
  certkeyname         = "servercert1"
  serverurl           = "https://adfs.example.com"
  username            = "adfsuser"
  password_wo         = var.authenticationadfsproxyprofile_password
  password_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `name` - (Required) Name for the ADFS proxy profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created. Changing this attribute forces a new resource to be created.
* `certkeyname` - (Required) SSL certificate of the proxy that is registered at the ADFS server for trust.
* `serverurl` - (Required) Fully qualified URL of the ADFS server.
* `username` - (Required) Name of an account in the directory that is used to authenticate the trust request from the Citrix ADC acting as a proxy.
* `password` - (Optional, Sensitive) Password of an account in the directory that is used to authenticate the trust request from the Citrix ADC acting as a proxy. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative. Either `password` or `password_wo` must be specified.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationadfsproxyprofile. It has the same value as the `name` attribute.


## Import

An authenticationadfsproxyprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationadfsproxyprofile.tf_adfsproxyprofile tf_adfsproxyprofile
```
