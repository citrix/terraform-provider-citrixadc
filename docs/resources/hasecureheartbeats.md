---
subcategory: "HA"
---

# Resource: hasecureheartbeats

The hasecureheartbeats resource is used to configure the HA secure heartbeats
parameters. This is a singleton (global) configuration resource.


## Example usage

```hcl
resource "citrixadc_hasecureheartbeats" "tf_hasecureheartbeats" {
  state = "ENABLED"
  hapsk = "presharedkey123"
}
```

### Write-only (ephemeral) secret

Use `hapsk_wo` to keep the pre-shared key out of Terraform state. Bump
`hapsk_wo_version` whenever the secret value changes so the provider re-applies
it:

```hcl
variable "ha_psk" {
  type      = string
  sensitive = true
}

resource "citrixadc_hasecureheartbeats" "tf_hasecureheartbeats" {
  state            = "ENABLED"
  hapsk_wo         = var.ha_psk
  hapsk_wo_version = 1
}
```


## Argument Reference

* `state` - (Optional) By enabling this option, HA heartbeats are securely exchanged between nodes. Possible values: [ ENABLED, DISABLED ]. Default value: DISABLED
* `hapsk` - (Optional) Pre shared key to be used for securing HA heartbeats. This is stored in Terraform state; prefer `hapsk_wo` for ephemeral (write-only) handling.
* `hapsk_wo` - (Optional, Write-only) Write-only (ephemeral) equivalent of `hapsk`. The value is used only during the apply and is never persisted to Terraform state. Pair it with `hapsk_wo_version` to trigger updates when the secret rotates. (Requires Terraform 1.11+.)
* `hapsk_wo_version` - (Optional) Version tracker for `hapsk_wo`. Increment this value to signal that the write-only secret changed so the provider pushes the new value to the appliance. Default value: 1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the hasecureheartbeats. Because this is a singleton resource, it has a fixed identifier.


## Import

A hasecureheartbeats can be imported using its id, e.g.

```shell
terraform import citrixadc_hasecureheartbeats.tf_hasecureheartbeats hasecureheartbeats-config
```
