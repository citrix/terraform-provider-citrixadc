---
subcategory: "Azure"
---

# Resource: azureapplication

The azureapplication resource is used to configure an Azure Active Directory application registration on the Citrix ADC. The application is used by the ADC to authenticate to Azure and obtain access tokens for Azure resources such as Azure Key Vault.


## Example usage

### Basic usage

```hcl
resource "citrixadc_azureapplication" "tf_azureapplication" {
  name          = "my_azure_app"
  clientid      = "11111111-2222-3333-4444-555555555555"
  tenantid      = "66666666-7777-8888-9999-000000000000"
  tokenendpoint = "https://login.microsoftonline.com/66666666-7777-8888-9999-000000000000"
  vaultresource = "vault.azure.net"
}
```

### Using clientsecret (sensitive attribute - persisted in state)

```hcl
variable "azureapplication_clientsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_azureapplication" "tf_azureapplication" {
  name          = "my_azure_app"
  clientid      = "11111111-2222-3333-4444-555555555555"
  clientsecret  = var.azureapplication_clientsecret
  tenantid      = "66666666-7777-8888-9999-000000000000"
  vaultresource = "vault.azure.net"
}
```

### Using clientsecret_wo (write-only/ephemeral - NOT persisted in state)

The `clientsecret_wo` attribute provides an ephemeral path for the Azure application client secret. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the client secret value changes, increment `clientsecret_wo_version`.

```hcl
variable "azureapplication_clientsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_azureapplication" "tf_azureapplication" {
  name                    = "my_azure_app"
  clientid                = "11111111-2222-3333-4444-555555555555"
  clientsecret_wo         = var.azureapplication_clientsecret
  clientsecret_wo_version = 1
  tenantid                = "66666666-7777-8888-9999-000000000000"
  vaultresource           = "vault.azure.net"
}
```

To rotate the client secret, update the variable value and bump the version:

```hcl
resource "citrixadc_azureapplication" "tf_azureapplication" {
  name                    = "my_azure_app"
  clientid                = "11111111-2222-3333-4444-555555555555"
  clientsecret_wo         = var.azureapplication_clientsecret
  clientsecret_wo_version = 2  # Bumped to trigger update
  tenantid                = "66666666-7777-8888-9999-000000000000"
  vaultresource           = "vault.azure.net"
}
```


## Argument Reference

* `name` - (Required) Name for the application. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, `"my application"` or `'my application'`).
* `vaultresource` - (Required) Vault resource for which the access token is granted. Example: `vault.azure.net`.
* `clientid` - (Optional) Application ID that is generated when an application is created in Azure Active Directory using either the Azure CLI or the Azure portal (GUI).
* `clientsecret` - (Optional, Sensitive) Password for the application configured in Azure Active Directory. The password is specified in the Azure CLI or generated in the Azure portal (GUI). The value is persisted in Terraform state (encrypted). See also `clientsecret_wo` for an ephemeral alternative.
* `clientsecret_wo` - (Optional, Sensitive, WriteOnly) Same as `clientsecret`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `clientsecret_wo_version`. If both `clientsecret` and `clientsecret_wo` are set, `clientsecret_wo` takes precedence.
* `clientsecret_wo_version` - (Optional) An integer version tracker for `clientsecret_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the client secret has changed and trigger an update. Defaults to `1`.
* `tenantid` - (Optional) ID of the directory inside Azure Active Directory in which the application was created.
* `tokenendpoint` - (Optional) URL from where the access token can be obtained. If the token end point is not specified, the default value is `https://login.microsoftonline.com/<tenant id>`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the azureapplication. It has the same value as the `name` attribute.


## Import

An azureapplication can be imported using its name, e.g.

```shell
terraform import citrixadc_azureapplication.tf_azureapplication my_azure_app
```
