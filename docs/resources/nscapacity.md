---
subcategory: "NS"
---

# Resource: nscapacity

The nscapacity resource is used to apply licenses to a target ADC from a license server.


## Example usage

```hcl
# CICO
resource "citrixadc_nscapacity" "tf_cico" {
  platform = "VP10000"
}

# Pooled
resource "citrixadc_nscapacity" "tf_pooled" {
  bandwidth = 100
  unit      = "Mbps"
  edition   = "Platinum"
}

# vCPU
resource "citrixadc_nscapacity" "tf_vcpu" {
  vcpu    = true
  edition = "Standard"
}
```

### Using an ephemeral (write-only) password

The `password_wo` attribute sends the LAS licensing password to the ADC without storing it in Terraform state. Pair it with `password_wo_version`; increment the version to re-send a rotated password.

```hcl
variable "las_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_nscapacity" "tf_pooled" {
  bandwidth           = 100
  unit                = "Mbps"
  edition             = "Platinum"
  username            = "las_user"
  password_wo         = var.las_password
  password_wo_version = 1
}
```


## Argument Reference

* `bandwidth` - (Optional) System bandwidth limit.
* `platform` - (Optional) appliance platform type. Possible values: [ VS10, VE10, VP10, VS25, VE25, VP25, VS200, VE200, VP200, VS1000, VE1000, VP1000, VS3000, VE3000, VP3000, VS5000, VE5000, VP5000, VS8000, VE8000, VP8000, VS10000, VE10000, VP10000, VS15000, VE15000, VP15000, VS25000, VE25000, VP25000, VS40000, VE40000, VP40000, VS100000, VE100000, VP100000, CP1000 ]
* `vcpu` - (Optional) licensed using vcpu pool.
* `edition` - (Optional) Product edition. Possible values: [ Standard, Enterprise, Platinum ]
* `unit` - (Optional) Bandwidth unit. Possible values: [ Gbps, Mbps ]
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `password` - (Optional, Sensitive) Password to use when authenticating with NetScaler Console Agent for LAS licensing. The value is persisted in Terraform state. See also `password_wo` for an ephemeral alternative.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `username` - (Optional) Username to authenticate with NetScaler Console Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `ignoreexpiry` - (Optional) Value to mention if days to expire data needs to be fetched or not.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nscapacity. It is a unique string prefixed with `tf-nscapacity-`

## Import

A nscapacity can be imported using its id, e.g.

```shell
terraform import citrixadc_nscapacity.tf_pooled tf-nscapacity-<some_random_string>
```