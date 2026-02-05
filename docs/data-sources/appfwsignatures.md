---
subcategory: "Application Firewall"
---

# Data Source: citrixadc_appfwsignatures

Use this data source to retrieve information about an existing Application Firewall Signatures object.

The `citrixadc_appfwsignatures` data source allows you to retrieve details of an Application Firewall signatures object by its name. This is useful for referencing existing signature objects in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing appfwsignatures object
data "citrixadc_appfwsignatures" "example" {
  name = "my_signature_object"
}

# Use the retrieved signature data in a profile binding
output "signature_source" {
  value = data.citrixadc_appfwsignatures.example.src
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the signature object to retrieve. Must match an existing signature object name.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the Application Firewall signatures object (same as name).

### Signature Configuration

* `src` - URL or file path for the location of the imported signatures object.
* `comment` - Any comments to preserve information about the signatures object.
* `overwrite` - Overwrite any existing signatures object of the same name.
* `merge` - Merges the existing Signature with new signature rules.
* `mergedefault` - Merges signature file with default signature file.
* `preservedefactions` - Preserves default actions of signature rules.
* `sha1` - File path for sha1 file to validate signature file.
* `xslt` - XSLT file source.
* `vendortype` - Third party vendor type for which WAF signatures has to be generated.

### Signature Management

* `autoenablenewsignatures` - Flag used to enable/disable auto enable new signatures.
* `enabled` - Flag used to enable/disable signature rule IDs/Signature Category.
* `category` - Signature category to be Enabled/Disabled.
* `ruleid` - Signature rule IDs to be Enabled/Disabled.
* `action` - Signature action.

## Import

Application Firewall signatures can be imported using the signature object name:

```shell
terraform import citrixadc_appfwsignatures.example my_signature_object
```

## Notes

* The `src` attribute returned by the data source may differ from the input format. For example, a signature imported from `local://appfw_signatures.xml` may be returned as just `appfw_signatures.xml`.
* Some attributes like `comment` may not be returned by the API in the GET response, even if they were set during creation.
* Signature objects are typically used in conjunction with Application Firewall profiles to provide protection against web application attacks.
