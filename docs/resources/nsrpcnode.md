---
subcategory: "NS"
---

# Resource: nsrpcnode

The nsrpcnode resource is used to manage rpc nodes.


## Example usage

### Basic usage

```hcl
resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "10.78.60.201"
    secure    = "ON"
    srcip     = "10.78.60.201"
}
```

### Using password (sensitive attribute - persisted in state)

```hcl
variable "nsrpcnode_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress = "10.78.60.201"
    password  = var.nsrpcnode_password
    secure    = "ON"
    srcip     = "10.78.60.201"
}
```

### Using password_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` attribute provides an ephemeral path for the RPC node password. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the password value changes, increment `password_wo_version`.

```hcl
variable "nsrpcnode_password" {
  type      = string
  sensitive = true
}

resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress           = "10.78.60.201"
    password_wo         = var.nsrpcnode_password
    password_wo_version = 1
    secure              = "ON"
    srcip               = "10.78.60.201"
}
```

To rotate the password, update the variable value and bump the version:

```hcl
resource "citrixadc_nsrpcnode" "tf_nsrpcnode" {
    ipaddress           = "10.78.60.201"
    password_wo         = var.nsrpcnode_password
    password_wo_version = 2  # Bumped to trigger update
    secure              = "ON"
    srcip               = "10.78.60.201"
}
```


## Argument Reference

* `ipaddress` - (Required) IP address of the node. This has to be in the same subnet as the NSIP address.
* `password` - (Optional, Sensitive) Password to be used in authentication with the peer system node. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the password has changed and trigger an update. Defaults to `1`.
* `srcip` - (Optional) Source IP address to be used to communicate with the peer system node. The default value is 0, which means that the appliance uses the NSIP address as the source IP address.
* `secure` - (Optional) State of the channel when talking to the node. Default value: ON | Possible values: [ ON, OFF ]
* `validatecert` - (Optional) validate the server certificate for secure SSL connections. Default value: NO | Possible values: [ YES, NO ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsrpcnode. It has the same value as the `ipaddress` attribute.


## Import

A nsrpcnode can be imported using its ipaddress, e.g.

```shell
terraform import citrixadc_nsrpcnode.tf_nsrpcnode 10.78.60.201
```
