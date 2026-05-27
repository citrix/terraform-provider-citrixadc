---
subcategory: "Azure"
---

# Data Source: azureapplication

The azureapplication data source allows you to retrieve information about an Azure Active Directory application registration that is configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_azureapplication" "tf_azureapplication" {
  name = "my_azure_app"
}

output "clientid" {
  value = data.citrixadc_azureapplication.tf_azureapplication.clientid
}

output "tenantid" {
  value = data.citrixadc_azureapplication.tf_azureapplication.tenantid
}

output "vaultresource" {
  value = data.citrixadc_azureapplication.tf_azureapplication.vaultresource
}
```


## Argument Reference

* `name` - (Required) Name of the Azure application configured on the Citrix ADC to look up.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `clientid` - Application ID that is generated when an application is created in Azure Active Directory using either the Azure CLI or the Azure portal (GUI).
* `tenantid` - ID of the directory inside Azure Active Directory in which the application was created.
* `tokenendpoint` - URL from where the access token can be obtained. If the token end point is not specified, the default value is `https://login.microsoftonline.com/<tenant id>`.
* `vaultresource` - Vault resource for which the access token is granted. Example: `vault.azure.net`.
* `id` - The id of the azureapplication. It has the same value as the `name` attribute.

Note: The `clientsecret` value is not returned by the Citrix ADC NITRO API and therefore is not exposed by this data source.
