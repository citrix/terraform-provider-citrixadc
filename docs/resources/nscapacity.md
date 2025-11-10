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


## Argument Reference

* `bandwidth` - (Optional) System bandwidth limit.
* `platform` - (Optional) appliance platform type. Possible values: [ VS10, VE10, VP10, VS25, VE25, VP25, VS200, VE200, VP200, VS1000, VE1000, VP1000, VS3000, VE3000, VP3000, VS5000, VE5000, VP5000, VS8000, VE8000, VP8000, VS10000, VE10000, VP10000, VS15000, VE15000, VP15000, VS25000, VE25000, VP25000, VS40000, VE40000, VP40000, VS100000, VE100000, VP100000, CP1000 ]
* `vcpu` - (Optional) licensed using vcpu pool.
* `edition` - (Optional) Product edition. Possible values: [ Standard, Enterprise, Platinum ]
* `unit` - (Optional) Bandwidth unit. Possible values: [ Gbps, Mbps ]
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `password` - (Optional) Password to use when authenticating with ADM Agent for LAS licensing.
* `username` - (Optional) Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nscapacity. It is a unique string prefixed with `tf-nscapacity-`

## Import

A nscapacity can be imported using its id, e.g.

```shell
terraform import citrixadc_nscapacity.tf_pooled tf-nscapacity-<some_random_string>
```