---
subcategory: "VPN"
---

# Resource: vpnsecureprivateaccessprofile

The vpnsecureprivateaccessprofile resource is used to create and manage Secure Private Access profiles.


## Example usage

### Using sharedsecret (sensitive attribute - persisted in state)

```hcl
variable "spa_sharedsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_vpnsecureprivateaccessprofile" "tf_vpnsecureprivateaccessprofile" {
  name            = "my_spaprofile"
  url             = "https://spa.example.com"
  forceclienttype = "ON"
  sharedsecret    = var.spa_sharedsecret
}
```

### Using sharedsecret_wo (write-only/ephemeral - NOT persisted in state)

The `sharedsecret_wo` attribute provides an ephemeral path for the shared secret. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `sharedsecret_wo_version`.

```hcl
variable "spa_sharedsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_vpnsecureprivateaccessprofile" "tf_vpnsecureprivateaccessprofile" {
  name                    = "my_spaprofile"
  url                     = "https://spa.example.com"
  forceclienttype         = "ON"
  sharedsecret_wo         = var.spa_sharedsecret
  sharedsecret_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_vpnsecureprivateaccessprofile" "tf_vpnsecureprivateaccessprofile" {
  name                    = "my_spaprofile"
  url                     = "https://spa.example.com"
  forceclienttype         = "ON"
  sharedsecret_wo         = var.spa_sharedsecret
  sharedsecret_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `name` - (Required) name of Secure Private Access profile.
* `url` - (Optional) Public URL for your Secure Private Access site or load balancing server.
* `customerid` - (Optional) Customer ID of the citrix cloud customer.
* `chromeenterprisepremiummode` - (Optional) Secure Private Access Chrome Enterprise Premium mode of operation. Possible values = OFF, WITH_PARTNER_CONNECTOR, WITHOUT_PARTNER_CONNECTOR
* `googlecustomerid` - (Optional) Your organization's unique ID on Google's Admin console Profile settings.
* `googlesecuritygatewayid` - (Optional) The ID of the Google Secure Gateway.
* `forceclienttype` - (Optional) Automatically configures the session for Citrix Secure Access client connectivity. Possible values = ON, OFF
* `sharedsecret` - (Optional, Sensitive) Secure Private Access Shared Secret. The value is persisted in Terraform state (encrypted). See also `sharedsecret_wo` for an ephemeral alternative.
* `sharedsecret_wo` - (Optional, Sensitive, WriteOnly) Same as `sharedsecret`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `sharedsecret_wo_version`. If both `sharedsecret` and `sharedsecret_wo` are set, `sharedsecret_wo` takes precedence.
* `sharedsecret_wo_version` - (Optional) An integer version tracker for `sharedsecret_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnsecureprivateaccessprofile. It has the same value as the `name` attribute.


## Import

A vpnsecureprivateaccessprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnsecureprivateaccessprofile.tf_vpnsecureprivateaccessprofile my_spaprofile
```
