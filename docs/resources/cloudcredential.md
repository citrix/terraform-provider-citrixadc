---
subcategory: "Cloud"
---

# Resource: cloudcredential

Configures the Azure application (service principal) credentials that the Citrix ADC uses to authenticate to Azure for cloud integrations such as Azure Key Vault. Create this resource to store the tenant ID, application (client) ID, and application secret the ADC presents when requesting Azure access tokens. This is a singleton resource: only one set of cloud credentials exists on the appliance, so there is no name key and the resource is never deleted from the ADC (removing it from your configuration only removes it from Terraform state).

~> **Note** The NITRO GET for `cloudcredential` returns only an "is set" flag; it does not echo back `tenantidentifier`, `applicationid`, or the secret. The provider therefore manages all configured values in Terraform state from your configuration rather than reading them back from the ADC.


## Example usage

### Using applicationsecret (sensitive attribute - persisted in state)

```hcl
variable "cloudcredential_applicationsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_cloudcredential" "tf_cloudcredential" {
  tenantidentifier  = "00000000-0000-0000-0000-000000000000"
  applicationid     = "11111111-1111-1111-1111-111111111111"
  applicationsecret = var.cloudcredential_applicationsecret
}
```

### Using applicationsecret_wo (write-only/ephemeral - NOT persisted in state)

The `applicationsecret_wo` attribute provides an ephemeral path for the application secret. The value is sent to the ADC but is **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when the value changes, increment `applicationsecret_wo_version`.

```hcl
variable "cloudcredential_applicationsecret" {
  type      = string
  sensitive = true
}

resource "citrixadc_cloudcredential" "tf_cloudcredential" {
  tenantidentifier             = "00000000-0000-0000-0000-000000000000"
  applicationid                = "11111111-1111-1111-1111-111111111111"
  applicationsecret_wo         = var.cloudcredential_applicationsecret
  applicationsecret_wo_version = 1
}
```

To rotate the secret, update the variable value and bump the version:

```hcl
resource "citrixadc_cloudcredential" "tf_cloudcredential" {
  tenantidentifier             = "00000000-0000-0000-0000-000000000000"
  applicationid                = "11111111-1111-1111-1111-111111111111"
  applicationsecret_wo         = var.cloudcredential_applicationsecret
  applicationsecret_wo_version = 2  # Bumped to trigger update
}
```


## Argument Reference

* `tenantidentifier` - (Required) Tenant ID of the credentials. Maximum length = 63.
* `applicationid` - (Required) Application (client) ID of the credentials. Maximum length = 63.
* `applicationsecret` - (Optional, Sensitive) Application secret of the credentials. Maximum length = 63. The value is persisted in Terraform state (encrypted) but is never returned by the NITRO GET, so it is managed in state from your configuration. See also `applicationsecret_wo` for an ephemeral alternative. Either `applicationsecret` or `applicationsecret_wo` must be specified.
* `applicationsecret_wo` - (Optional, Sensitive, WriteOnly) Same as `applicationsecret`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `applicationsecret_wo_version`. If both `applicationsecret` and `applicationsecret_wo` are set, `applicationsecret_wo` takes precedence. Either `applicationsecret` or `applicationsecret_wo` must be specified. Maximum length = 63.
* `applicationsecret_wo_version` - (Optional) An integer version tracker for `applicationsecret_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger an update. Defaults to `1`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudcredential. Because this is a singleton resource, it is the fixed string `cloudcredential-config`.
