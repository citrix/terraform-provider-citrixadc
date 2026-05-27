---
subcategory: "Authentication"
---

# Data Source: authenticationazurekeyvault

This data source is used to retrieve information about a specific `authenticationazurekeyvault` resource configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_authenticationazurekeyvault" "example" {
  name = "my_azurekeyvault"
}

output "azurekeyvault_vaultname" {
  value = data.citrixadc_authenticationazurekeyvault.example.vaultname
}
```

## Argument Reference

* `name` - (Required) Name of the Azure Key Vault profile to look up on the Citrix ADC.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the authenticationazurekeyvault. It has the same value as the `name` attribute.
* `authentication` - If authentication is disabled, otp checks are not performed after azure vault keys are obtained. This is useful to distinguish whether user has registered devices.
* `clientid` - Unique identity of the relying party requesting for authentication.
* `defaultauthenticationgroup` - This is the group that is added to user sessions that match current IdP policy. It can be used in policies to identify relying party trust.
* `pushservice` - Name of the service used to send push notifications.
* `refreshinterval` - Interval at which access token in obtained.
* `servicekeyname` - Friendly name of the Key to be used to compute signature.
* `signaturealg` - Algorithm to be used to sign/verify transactions.
* `tenantid` - TenantID of the application. This is usually specific to providers such as Microsoft and usually refers to the deployment identifier.
* `tokenendpoint` - URL endpoint on relying party to which the OAuth token is to be sent.
* `vaultname` - Name of the Azure vault account as configured in azure portal.

~> **Note** The `clientsecret`, `clientsecret_wo`, and `clientsecret_wo_version` attributes are sensitive/write-only and are not returned by the NITRO API, so they are not exposed by this data source.
