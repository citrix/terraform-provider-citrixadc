---
subcategory: "Basic"
---

# Resource: radiusnode

The radiusnode resource is used for Configuration of RADIUS Node resource


## Example usage

### Using radkey (sensitive attribute - persisted in state)

```hcl
variable "radiusnode_radkey" {
  type      = string
  sensitive = true
}

resource "citrixadc_radiusnode" "tf_radiusnode" {
  nodeprefix = "10.10.10.10/32"
  radkey     = var.radiusnode_radkey
}
```

### Using radkey_wo (write-only/ephemeral - NOT persisted in state)

The `radkey_wo` attribute provides an ephemeral path for the shared RADIUS key. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `radkey_wo_version`.

```hcl
variable "radiusnode_radkey" {
  type      = string
  sensitive = true
}

resource "citrixadc_radiusnode" "tf_radiusnode" {
  nodeprefix        = "10.10.10.10/32"
  radkey_wo         = var.radiusnode_radkey
  radkey_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_radiusnode" "tf_radiusnode" {
  nodeprefix        = "10.10.10.10/32"
  radkey_wo         = var.radiusnode_radkey
  radkey_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `nodeprefix` - (Required) IP address/IP prefix of radius node in CIDR format
* `radkey` - (Optional, Sensitive) The key shared between the RADIUS server and clients. Required for Citrix ADC to communicate with the RADIUS nodes. The value is persisted in Terraform state (encrypted). See also `radkey_wo` for an ephemeral alternative.
* `radkey_wo` - (Optional, Sensitive, WriteOnly) Same as `radkey`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `radkey_wo_version`. If both `radkey` and `radkey_wo` are set, `radkey_wo` takes precedence.
* `radkey_wo_version` - (Optional) An integer version tracker for `radkey_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the radiusnode. It has the same value as the `nodeprefix` attribute.


## Import

A radiusnode can be imported using its nodeprefix, e.g.

```shell
terraform import citrixadc_radiusnode.tf_radiusnode 10.10.10.10/32
```
