---
subcategory: "RDP"
---

# Resource: rdpserverprofile

The rdpserverprofile resource is used to create rdpserverprofile.


## Example usage

### Using psk (sensitive attribute - persisted in state)

```hcl
variable "rdpserverprofile_psk" {
  type      = string
  sensitive = true
}

resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile" {
  name           = "my_rdpserverprofile"
  psk            = var.rdpserverprofile_psk
  rdpredirection = "ENABLE"
  rdpport        = 4000
}
```

### Using psk_wo (write-only/ephemeral - NOT persisted in state)

The `psk_wo` attribute provides an ephemeral path for the pre-shared key. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `psk_wo_version`.

```hcl
variable "rdpserverprofile_psk" {
  type      = string
  sensitive = true
}

resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile" {
  name           = "my_rdpserverprofile"
  psk_wo         = var.rdpserverprofile_psk
  psk_wo_version = 1
  rdpredirection = "ENABLE"
  rdpport        = 4000
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_rdpserverprofile" "tf_rdpserverprofile" {
  name           = "my_rdpserverprofile"
  psk_wo         = var.rdpserverprofile_psk
  psk_wo_version = 2  # Bumped to trigger update
  rdpredirection = "ENABLE"
  rdpport        = 4000
}
```


## Argument Reference

* `name` - (Required) The name of the rdp server profile
* `psk` - (Optional, Sensitive) Pre shared key value. The value is persisted in Terraform state (encrypted). See also `psk_wo` for an ephemeral alternative.
* `psk_wo` - (Optional, Sensitive, WriteOnly) Same as `psk`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `psk_wo_version`. If both `psk` and `psk_wo` are set, `psk_wo` takes precedence.
* `psk_wo_version` - (Optional) An integer version tracker for `psk_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `rdpip` - (Optional) IPv4 or IPv6 address of RDP listener. This terminates client RDP connections.
* `rdpport` - (Optional) TCP port on which the RDP connection is established.
* `rdpredirection` - (Optional) Enable/Disable RDP redirection support. This needs to be enabled in presence of connection broker or session directory with IP cookie(msts cookie) based redirection support


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the rdpserverprofile. It has the same value as the `name` attribute.


## Import

A rdpserverprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_rdpserverprofile.tf_rdpserverprofile my_rdpserverprofile
```
