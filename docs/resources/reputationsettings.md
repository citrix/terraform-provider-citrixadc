---
subcategory: "Reputation"
---

# Resource: reputationsettings

The reputationsettings resource is used to create reputationsettings.


## Example usage

### Basic usage

```hcl
resource "citrixadc_reputationsettings" "tf_reputationsettings" {
  proxyserver = "my_proxyserver"
  proxyport   = 3500
}
```

### Using proxypassword (sensitive attribute - persisted in state)

```hcl
variable "reputationsettings_proxypassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_reputationsettings" "tf_reputationsettings" {
  proxyserver   = "my_proxyserver"
  proxyport     = 3500
  proxyusername = "proxyuser"
  proxypassword = var.reputationsettings_proxypassword
}
```

### Using proxypassword_wo (write-only/ephemeral - NOT persisted in state)

The `proxypassword_wo` attribute provides an ephemeral path for the proxy password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the password value changes, increment `proxypassword_wo_version`.

```hcl
variable "reputationsettings_proxypassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_reputationsettings" "tf_reputationsettings" {
  proxyserver          = "my_proxyserver"
  proxyport            = 3500
  proxyusername        = "proxyuser"
  proxypassword_wo     = var.reputationsettings_proxypassword
  proxypassword_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_reputationsettings" "tf_reputationsettings" {
  proxyserver          = "my_proxyserver"
  proxyport            = 3500
  proxyusername        = "proxyuser"
  proxypassword_wo     = var.reputationsettings_proxypassword
  proxypassword_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `proxyport` - (Optional) Proxy server port.
* `proxyserver` - (Optional) Proxy server IP to get Reputation data.
* `proxypassword` - (Optional, Sensitive) Password with which user logs on. The value is persisted in Terraform state (encrypted). See also `proxypassword_wo` for an ephemeral alternative.
* `proxypassword_wo` - (Optional, Sensitive, WriteOnly) Same as `proxypassword`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `proxypassword_wo_version`. If both `proxypassword` and `proxypassword_wo` are set, `proxypassword_wo` takes precedence.
* `proxypassword_wo_version` - (Optional) An integer version tracker for `proxypassword_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `proxyusername` - (Optional) Proxy Username


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the reputationsettings. It is a unique string prefixed with `tf-reputationsettings-` attribute.


## Import

A reputationsettings can be imported using its id, e.g.

```shell
terraform import citrixadc_reputationsettings.tf_reputationsettings reputationsettings-config
```