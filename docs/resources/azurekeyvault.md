---
subcategory: "Azure"
---

# Resource: azurekeyvault

The azurekeyvault resource is used to configure an Azure Key Vault entity on the Citrix ADC. The Key Vault, in combination with an Azure application registration, allows the ADC to use keys and certificates stored in Azure for SSL operations.


## Example usage

```hcl
resource "citrixadc_azureapplication" "tf_azureapplication" {
  name          = "my_azure_app"
  clientid      = "11111111-2222-3333-4444-555555555555"
  tenantid      = "66666666-7777-8888-9999-000000000000"
  tokenendpoint = "https://login.microsoftonline.com/66666666-7777-8888-9999-000000000000"
  vaultresource = "vault.azure.net"
}

resource "citrixadc_azurekeyvault" "tf_azurekeyvault" {
  name             = "my_keyvault"
  azurevaultname   = "Test.vault.azure.net"
  azureapplication = citrixadc_azureapplication.tf_azureapplication.name
}
```


## Argument Reference

* `name` - (Required) Name for the Key Vault. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, `"my keyvault"` or `'my keyvault'`).
* `azureapplication` - (Required) Name of the Azure Application object created on the ADC appliance. This object is used for authentication with Azure Active Directory.
* `azurevaultname` - (Optional) Name of the Key Vault configured in Azure cloud using either the Azure CLI or the Azure portal (GUI) with complete domain name. Example: `Test.vault.azure.net`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the azurekeyvault. It has the same value as the `name` attribute.


## Import

An azurekeyvault can be imported using its name, e.g.

```shell
terraform import citrixadc_azurekeyvault.tf_azurekeyvault my_keyvault
```
