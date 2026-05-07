---
subcategory: "Authentication"
---

# Resource: authenticationpushservice

The authenticationpushservice resource is used to create authentication pushservice resource.


## Example usage

### Using clientsecret (sensitive attribute - persisted in state)

```hcl
variable "authenticationpushservice_clientsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationpushservice" "tf_pushservice" {
  name            = "tf_pushservice"
  clientid        = "cliId"
  clientsecret    = var.authenticationpushservice_clientsecret
  customerid      = "cusID"
  refreshinterval = 50
}
```

### Using clientsecret_wo (write-only/ephemeral - NOT persisted in state)

The `clientsecret_wo` attribute provides an ephemeral path for the client secret. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `clientsecret_wo_version`.

```hcl
variable "authenticationpushservice_clientsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationpushservice" "tf_pushservice" {
  name                   = "tf_pushservice"
  clientid               = "cliId"
  clientsecret_wo        = var.authenticationpushservice_clientsecret
  clientsecret_wo_version = 1
  customerid             = "cusID"
  refreshinterval        = 50
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_authenticationpushservice" "tf_pushservice" {
  name                   = "tf_pushservice"
  clientid               = "cliId"
  clientsecret_wo        = var.authenticationpushservice_clientsecret
  clientsecret_wo_version = 2  # Bumped to trigger update
  customerid             = "cusID"
  refreshinterval        = 50
}
```


## Argument Reference

* `name` - (Required) Name for the push service. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created. 	    CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my push service" or 'my push service').
* `clientid` - (Optional) Unique identity for communicating with Citrix Push server in cloud.
* `clientsecret` - (Optional, Sensitive) Unique secret for communicating with Citrix Push server in cloud. The value is persisted in Terraform state (encrypted). See also `clientsecret_wo` for an ephemeral alternative.
* `clientsecret_wo` - (Optional, Sensitive, WriteOnly) Same as `clientsecret`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `clientsecret_wo_version`. If both `clientsecret` and `clientsecret_wo` are set, `clientsecret_wo` takes precedence.
* `clientsecret_wo_version` - (Optional) An integer version tracker for `clientsecret_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `customerid` - (Optional) Customer id/name of the account in cloud that is used to create clientid/secret pair.
* `refreshinterval` - (Optional) Interval at which certificates or idtoken is refreshed.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationpushservice. It has the same value as the `name` attribute.


## Import

A authenticationpushservice can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationpushservice.tf_pushservice tf_pushservice
```
