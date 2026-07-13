---
subcategory: "Authentication"
---

# Resource: authenticationazurekeyvault

The authenticationazurekeyvault resource is used to create and manage Azure Key Vault profiles on the Citrix ADC. This profile is used by the Citrix ADC to authenticate to Azure Key Vault for retrieving keys used to sign authentication transactions.


## Example usage

### Using clientsecret (sensitive attribute - persisted in state)

```hcl
variable "authenticationazurekeyvault_clientsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationazurekeyvault" "tf_azurekeyvault" {
  name           = "tf_azurekeyvault"
  clientid       = "00000000-0000-0000-0000-000000000000"
  clientsecret   = var.authenticationazurekeyvault_clientsecret
  servicekeyname = "myKey"
  vaultname      = "myvault"
  tenantid       = "11111111-1111-1111-1111-111111111111"
  tokenendpoint  = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
}
```

### Using clientsecret_wo (write-only/ephemeral - NOT persisted in state)

The `clientsecret_wo` attribute provides an ephemeral path for the client secret. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `clientsecret_wo_version`.

```hcl
variable "authenticationazurekeyvault_clientsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_authenticationazurekeyvault" "tf_azurekeyvault" {
  name                    = "tf_azurekeyvault"
  clientid                = "00000000-0000-0000-0000-000000000000"
  clientsecret_wo         = var.authenticationazurekeyvault_clientsecret
  clientsecret_wo_version = 1
  servicekeyname          = "myKey"
  vaultname               = "myvault"
  tenantid                = "11111111-1111-1111-1111-111111111111"
  tokenendpoint           = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_authenticationazurekeyvault" "tf_azurekeyvault" {
  name                    = "tf_azurekeyvault"
  clientid                = "00000000-0000-0000-0000-000000000000"
  clientsecret_wo         = var.authenticationazurekeyvault_clientsecret
  clientsecret_wo_version = 2  # Bumped to trigger update
  servicekeyname          = "myKey"
  vaultname               = "myvault"
  tenantid                = "11111111-1111-1111-1111-111111111111"
  tokenendpoint           = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
}
```


## Argument Reference

* `name` - (Required) Name for the new Azure Key Vault profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `clientid` - (Required) Unique identity of the relying party requesting for authentication.
* `servicekeyname` - (Required) Friendly name of the Key to be used to compute signature.
* `vaultname` - (Required) Name of the Azure vault account as configured in azure portal.
* `clientsecret` - (Optional, Sensitive) Unique secret string to authorize relying party at authorization server. The value is persisted in Terraform state (encrypted). See also `clientsecret_wo` for an ephemeral alternative. At least one of `clientsecret` or `clientsecret_wo` must be set.
* `clientsecret_wo` - (Optional, Sensitive, WriteOnly) Same as `clientsecret`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `clientsecret_wo_version`. If both `clientsecret` and `clientsecret_wo` are set, `clientsecret_wo` takes precedence. At least one of `clientsecret` or `clientsecret_wo` must be set.
* `clientsecret_wo_version` - (Optional) An integer version tracker for `clientsecret_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.
* `authentication` - (Optional) If authentication is disabled, otp checks are not performed after azure vault keys are obtained. This is useful to distinguish whether user has registered devices. Defaults to `"ENABLED"`.
* `defaultauthenticationgroup` - (Optional) This is the group that is added to user sessions that match current IdP policy. It can be used in policies to identify relying party trust.
* `pushservice` - (Optional) Name of the service used to send push notifications.
* `refreshinterval` - (Optional) Interval at which access token in obtained. Defaults to `50`.
* `signaturealg` - (Optional) Algorithm to be used to sign/verify transactions. Defaults to `"RS256"`.
* `tenantid` - (Optional) TenantID of the application. This is usually specific to providers such as Microsoft and usually refers to the deployment identifier.
* `tokenendpoint` - (Optional) URL endpoint on relying party to which the OAuth token is to be sent.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationazurekeyvault. It has the same value as the `name` attribute.


## Import

An authenticationazurekeyvault can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationazurekeyvault.tf_azurekeyvault tf_azurekeyvault
```
