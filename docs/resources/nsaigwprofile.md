---
subcategory: "NS"
---

# Resource: nsaigwprofile

The nsaigwprofile resource is used to create and manage AI GW (AI Gateway) profiles.


## Example usage

```hcl
resource "citrixadc_nsaigwprofile" "tf_nsaigwprofile" {
  name                  = "my_nsaigwprofile"
  endpointtype          = "azureopenai"
  profiletype           = "frontend"
  tokenquota            = 1000
  quotarefreshfrequency = 60
}
```

### Using authtoken (sensitive attribute - persisted in state)

```hcl
variable "nsaigwprofile_authtoken" {
  type      = string
  sensitive = true
}

resource "citrixadc_nsaigwprofile" "tf_nsaigwprofile" {
  name        = "my_nsaigwprofile"
  profiletype = "frontend"
  authtoken   = var.nsaigwprofile_authtoken
}
```

### Using authtoken_wo (write-only/ephemeral - NOT persisted in state)

The `authtoken_wo` attribute provides an ephemeral path for the authentication token / API key. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `authtoken_wo_version`.

```hcl
variable "nsaigwprofile_authtoken" {
  type      = string
  sensitive = true
}

resource "citrixadc_nsaigwprofile" "tf_nsaigwprofile" {
  name                 = "my_nsaigwprofile"
  profiletype          = "frontend"
  authtoken_wo         = var.nsaigwprofile_authtoken
  authtoken_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_nsaigwprofile" "tf_nsaigwprofile" {
  name                 = "my_nsaigwprofile"
  profiletype          = "frontend"
  authtoken_wo         = var.nsaigwprofile_authtoken
  authtoken_wo_version = 2 # Bumped to trigger update
}
```


## Argument Reference

* `name` - (Required) Name of the AIGW Profile. Changing this forces a new resource to be created.
* `endpointtype` - (Optional) The type of AI GW endpoint type. Possible values = azureopenai
* `profiletype` - (Optional) The binding entity for the aigw profile. Possible values = frontend, backend
* `tokenquota` - (Optional) Token capacity of the backend server.
* `quotarefreshfrequency` - (Optional) Quota refresh rate, in minutes.
* `authtoken` - (Optional, Sensitive) Authentication token/API Key for the AI GW Endpoint. The value is persisted in Terraform state. See also `authtoken_wo` for an ephemeral alternative.
* `authtoken_wo` - (Optional, Sensitive, WriteOnly) Same as `authtoken`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `authtoken_wo_version`. If both `authtoken` and `authtoken_wo` are set, `authtoken_wo` takes precedence.
* `authtoken_wo_version` - (Optional) An integer version tracker for `authtoken_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsaigwprofile. It has the same value as the `name` attribute.


## Import

A nsaigwprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_nsaigwprofile.tf_nsaigwprofile my_nsaigwprofile
```
