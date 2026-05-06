---
subcategory: "AAA"
---

# Resource: aaassoprofile

The aaassoprofile resource is used to create an AAA SSO Profile for single sign-on configuration.


## Example usage

### Using `password` (sensitive attribute - persisted in state)

```hcl
variable "aaassoprofile_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaassoprofile" "tf_aaassoprofile" {
  name     = "myssoprofile"
  username = "john"
  password = var.aaassoprofile_password
}
```

### Using `password_wo` (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the SSO profile password. The value is sent to the Citrix ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `password_wo_version`.

```hcl
variable "aaassoprofile_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_aaassoprofile" "tf_aaassoprofile" {
  name               = "myssoprofile"
  username           = "john"
  password_wo        = var.aaassoprofile_password
  password_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_aaassoprofile" "tf_aaassoprofile" {
  name               = "myssoprofile"
  username           = "john"
  password_wo        = var.aaassoprofile_password
  password_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `name` - (Required) Name for the SSO Profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a SSO Profile is created. The following requirement applies only to the NetScaler CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `username` - (Required) Name for the user. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my group" or 'my group').
* `password` - (Optional, Sensitive) Password with which the user logs on. Required for Single sign on to external server. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative. Either `password` or `password_wo` must be specified.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaassoprofile. It has the same value as the `name` attribute.


## Import

A aaassoprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_aaassoprofile.tf_aaassoprofile myssoprofile
```
