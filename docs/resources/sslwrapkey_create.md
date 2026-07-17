---
subcategory: "SSL"
---

# Resource: sslwrapkey_create

The sslwrapkey_create resource creates a named SSL wrap key on the Citrix ADC through the NITRO `create` action. A wrap key is a password- and salt-derived key used by the appliance's crypto subsystem to wrap (encrypt) and unwrap other key material — for example, when exporting or importing protected keys. The key is immutable: any change to its attributes forces a new wrap key to be created. The resource manages a full lifecycle — it creates the wrap key on apply and deletes it on destroy.

~> **WARNING: FIPS / crypto subsystem required.**
> This resource depends on the Citrix ADC FIPS / cryptographic subsystem being available on the appliance. It is intended for FIPS-capable appliances and may not be supported on appliances where the crypto subsystem is unavailable; the create action will fail in that case. Validate against your target platform before applying.

## Example usage

### Using the sensitive password and salt attributes (persisted in state)

```hcl
variable "sslwrapkey_create_password" {
  type      = string
  sensitive = true
}

variable "sslwrapkey_create_salt" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslwrapkey_create" "tf_sslwrapkey_create" {
  wrapkeyname = "mywrapkey"
  password    = var.sslwrapkey_create_password
  salt        = var.sslwrapkey_create_salt
}
```

### Using the write-only password and salt attributes (NOT persisted in state)

The `password_wo` and `salt_wo` attributes provide an ephemeral path for the wrap key secrets. The values are sent to the Citrix ADC but are **not stored in Terraform state**, reducing the risk of secret exposure. Because the wrap key is immutable, changing any of these values (including bumping a `*_wo_version`) forces the wrap key to be replaced.

```hcl
variable "sslwrapkey_create_password" {
  type      = string
  sensitive = true
}

variable "sslwrapkey_create_salt" {
  type      = string
  sensitive = true
}

resource "citrixadc_sslwrapkey_create" "tf_sslwrapkey_create" {
  wrapkeyname         = "mywrapkey"
  password_wo         = var.sslwrapkey_create_password
  password_wo_version = 1
  salt_wo             = var.sslwrapkey_create_salt
  salt_wo_version     = 1
}
```

To rotate a secret, update the variable value and bump the corresponding version (this replaces the wrap key):

```hcl
resource "citrixadc_sslwrapkey_create" "tf_sslwrapkey_create" {
  wrapkeyname         = "mywrapkey"
  password_wo         = var.sslwrapkey_create_password
  password_wo_version = 2  # Bumped to trigger replacement
  salt_wo             = var.sslwrapkey_create_salt
  salt_wo_version     = 1
}
```

## Argument Reference

* `wrapkeyname` - (Required) Name for the wrap key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the wrap key is created. The following requirement applies only to the Citrix ADC CLI: if the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my key" or 'my key'). Changing this attribute forces a new resource to be created.
* `password` - (Optional, Sensitive) Password string for the wrap key. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative. Either `password` or `password_wo` must be specified. Changing this attribute forces a new resource to be created.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence. Changing this attribute forces a new resource to be created.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed; because the wrap key is immutable, this forces the resource to be replaced. Defaults to `1`.
* `salt` - (Optional, Sensitive) Salt string for the wrap key. The value is persisted in Terraform state (encrypted). See also `salt_wo` for an ephemeral alternative. Either `salt` or `salt_wo` must be specified. Changing this attribute forces a new resource to be created.
* `salt_wo` - (Optional, Sensitive, WriteOnly) Same as `salt`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `salt_wo_version`. If both `salt` and `salt_wo` are set, `salt_wo` takes precedence. Changing this attribute forces a new resource to be created.
* `salt_wo_version` - (Optional) An integer version tracker for `salt_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed; because the wrap key is immutable, this forces the resource to be replaced. Defaults to `1`.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslwrapkey_create resource. It has the same value as the `wrapkeyname` attribute.

## Import

A sslwrapkey_create resource can be imported using its wrapkeyname, e.g.

```shell
terraform import citrixadc_sslwrapkey_create.tf_sslwrapkey_create mywrapkey
```
