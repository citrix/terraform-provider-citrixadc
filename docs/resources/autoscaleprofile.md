---
subcategory: "Autoscale"
---

# Resource: autoscaleprofile

The autoscaleprofile resource is used to create autoscaleprofile.


## Example usage

### Using apikey and sharedsecret (sensitive attributes - persisted in state)

```hcl
variable "autoscaleprofile_apikey" {
  type      = string
  sensitive = true
}

variable "autoscaleprofile_sharedsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
  name         = "my_autoscaleprofile"
  type         = "CLOUDSTACK"
  url          = "www.service.example.com"
  apikey       = var.autoscaleprofile_apikey
  sharedsecret = var.autoscaleprofile_sharedsecret
}
```

### Using apikey_wo and sharedsecret_wo (write-only/ephemeral - NOT persisted in state)

The `apikey_wo` and `sharedsecret_wo` attributes provide an ephemeral path for API credentials. The values are sent to the ADC but are **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when a value changes, increment the corresponding `_wo_version`.

```hcl
variable "autoscaleprofile_apikey" {
  type      = string
  sensitive = true
}

variable "autoscaleprofile_sharedsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
  name                   = "my_autoscaleprofile"
  type                   = "CLOUDSTACK"
  url                    = "www.service.example.com"
  apikey_wo              = var.autoscaleprofile_apikey
  apikey_wo_version      = 1
  sharedsecret_wo        = var.autoscaleprofile_sharedsecret
  sharedsecret_wo_version = 1
}
```

To rotate the secrets, update the variable values and bump the version numbers:

```hcl
resource "citrixadc_autoscaleprofile" "tf_autoscaleprofile" {
  name                   = "my_autoscaleprofile"
  type                   = "CLOUDSTACK"
  url                    = "www.service.example.com"
  apikey_wo              = var.autoscaleprofile_apikey
  apikey_wo_version      = 2  # Bumped to trigger update
  sharedsecret_wo        = var.autoscaleprofile_sharedsecret
  sharedsecret_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `apikey` - (Optional, Sensitive) Api key for authentication with service. The value is persisted in Terraform state (encrypted). See also `apikey_wo` for an ephemeral alternative. At least one of `apikey` or `apikey_wo` must be set.
* `apikey_wo` - (Optional, Sensitive, WriteOnly) Same as `apikey`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `apikey_wo_version`. If both `apikey` and `apikey_wo` are set, `apikey_wo` takes precedence. At least one of `apikey` or `apikey_wo` must be set.
* `apikey_wo_version` - (Optional) An integer version tracker for `apikey_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `name` - (Required) AutoScale profile name.
* `sharedsecret` - (Optional, Sensitive) Shared secret for authentication with service. The value is persisted in Terraform state (encrypted). See also `sharedsecret_wo` for an ephemeral alternative. At least one of `sharedsecret` or `sharedsecret_wo` must be set.
* `sharedsecret_wo` - (Optional, Sensitive, WriteOnly) Same as `sharedsecret`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `sharedsecret_wo_version`. If both `sharedsecret` and `sharedsecret_wo` are set, `sharedsecret_wo` takes precedence. At least one of `sharedsecret` or `sharedsecret_wo` must be set.
* `sharedsecret_wo_version` - (Optional) An integer version tracker for `sharedsecret_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `type` - (Required) The type of profile.
* `url` - (Required) URL providing the service


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the autoscaleprofile. It has the same value as the `name` attribute.


## Import

A autoscaleprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_autoscaleprofile.tf_autoscaleprofile my_autoscaleprofile
```
