---
subcategory: "Azure"
---

# Data Source: azurekeyvault

The azurekeyvault data source allows you to retrieve information about an Azure Key Vault entity that is configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_azurekeyvault" "tf_azurekeyvault" {
  name = "my_keyvault"
}

output "azureapplication" {
  value = data.citrixadc_azurekeyvault.tf_azurekeyvault.azureapplication
}

output "azurevaultname" {
  value = data.citrixadc_azurekeyvault.tf_azurekeyvault.azurevaultname
}
```


## Argument Reference

* `name` - (Required) Name of the Azure Key Vault configured on the Citrix ADC to look up.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `azureapplication` - Name of the Azure Application object created on the ADC appliance. This object is used for authentication with Azure Active Directory.
* `azurevaultname` - Name of the Key Vault configured in Azure cloud using either the Azure CLI or the Azure portal (GUI) with complete domain name. Example: `Test.vault.azure.net`.
* `id` - The id of the azurekeyvault. It has the same value as the `name` attribute.
