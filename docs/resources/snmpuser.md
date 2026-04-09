---
subcategory: "SNMP"
---

# Resource: snmpuser

The snmpuser resource is used to create snmpuser.


## Example usage

### Basic usage

```hcl
resource "citrixadc_snmpuser" "tf_snmpuser" {
  name     = "test_user"
  group    = "test_group"
  authtype = "MD5"
  privtype = "DES"
}
```

### Using authpasswd and privpasswd (sensitive attributes - persisted in state)

```hcl
variable "snmpuser_authpasswd" {
  type      = string
  sensitive = true
}

variable "snmpuser_privpasswd" {
  type      = string
  sensitive = true
}

resource "citrixadc_snmpuser" "tf_snmpuser" {
  name       = "test_user"
  group      = "test_group"
  authtype   = "MD5"
  authpasswd = var.snmpuser_authpasswd
  privtype   = "DES"
  privpasswd = var.snmpuser_privpasswd
}
```

### Using authpasswd_wo and privpasswd_wo (write-only/ephemeral - NOT persisted in state)

The `authpasswd_wo` and `privpasswd_wo` attributes provide an ephemeral path for the authentication and encryption passwords. The values are sent to the ADC but are **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when a value changes, increment the corresponding `_wo_version`.

```hcl
variable "snmpuser_authpasswd" {
  type      = string
  sensitive = true
}

variable "snmpuser_privpasswd" {
  type      = string
  sensitive = true
}

resource "citrixadc_snmpuser" "tf_snmpuser" {
  name                  = "test_user"
  group                 = "test_group"
  authtype              = "MD5"
  authpasswd_wo         = var.snmpuser_authpasswd
  authpasswd_wo_version = 1
  privtype              = "DES"
  privpasswd_wo         = var.snmpuser_privpasswd
  privpasswd_wo_version = 1
}
```

To rotate the passwords, update the variable values and bump the versions:

```hcl
resource "citrixadc_snmpuser" "tf_snmpuser" {
  name                  = "test_user"
  group                 = "test_group"
  authtype              = "MD5"
  authpasswd_wo         = var.snmpuser_authpasswd
  authpasswd_wo_version = 2  # Bumped to trigger update
  privtype              = "DES"
  privpasswd_wo         = var.snmpuser_privpasswd
  privpasswd_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `name` - (Required) Name for the SNMPv3 user. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose it in double or single quotation marks (for example, "my user" or 'my user').
* `group` - (Required) Name of the configured SNMPv3 group to which to bind this SNMPv3 user. The access rights (bound SNMPv3 views) and security level set for this group are assigned to this user.
* `authpasswd` - (Optional, Sensitive) Plain-text pass phrase to be used by the authentication algorithm specified by the authType (Authentication Type) parameter. Can consist of 8 to 63 characters. The value is persisted in Terraform state (encrypted). See also `authpasswd_wo` for an ephemeral alternative.
* `authpasswd_wo` - (Optional, Sensitive, WriteOnly) Same as `authpasswd`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `authpasswd_wo_version`. If both `authpasswd` and `authpasswd_wo` are set, `authpasswd_wo` takes precedence.
* `authpasswd_wo_version` - (Optional) An integer version tracker for `authpasswd_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the pass phrase has changed and trigger an update. Defaults to `1`.
* `authtype` - (Optional) Authentication algorithm used by the Citrix ADC and the SNMPv3 user for authenticating the communication between them. You must specify the same authentication algorithm when you configure the SNMPv3 user in the SNMP manager.
* `privpasswd` - (Optional, Sensitive) Encryption key to be used by the encryption algorithm specified by the privType (Encryption Type) parameter. Can consist of 8 to 63 characters. The value is persisted in Terraform state (encrypted). See also `privpasswd_wo` for an ephemeral alternative.
* `privpasswd_wo` - (Optional, Sensitive, WriteOnly) Same as `privpasswd`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `privpasswd_wo_version`. If both `privpasswd` and `privpasswd_wo` are set, `privpasswd_wo` takes precedence.
* `privpasswd_wo_version` - (Optional) An integer version tracker for `privpasswd_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the encryption key has changed and trigger an update. Defaults to `1`.
* `privtype` - (Optional) Encryption algorithm used by the Citrix ADC and the SNMPv3 user for encrypting the communication between them. You must specify the same encryption algorithm when you configure the SNMPv3 user in the SNMP manager.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the snmpuser. It has the same value as the `name` attribute.


## Import

A snmpuser can be imported using its name, e.g.

```shell
terraform import citrixadc_snmpuser.tf_snmpuser test_user
```
