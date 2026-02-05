---
subcategory: "Bot"
---

# Data Source: citrixadc_botsignature

Use this data source to retrieve information about an existing Bot Signature.

The `citrixadc_botsignature` data source allows you to retrieve details of a bot signature by its name. This is useful for referencing existing bot signatures in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing bot signature
data "citrixadc_botsignature" "example" {
  name = "my_bot_signature"
}

# Reference signature attributes
output "signature_src" {
  value = data.citrixadc_botsignature.example.src
}

output "signature_comment" {
  value = data.citrixadc_botsignature.example.comment
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name to assign to the bot signature file object on the Citrix ADC.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the bot signature (same as name).

* `comment` - Any comments to preserve information about the signature file object.

* `overwrite` - Overwrites the existing file.

* `src` - Local path to and name of, or URL (protocol, host, path, and file name) for, the file in which to store the imported signature file. NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.

## Common Use Cases

### Retrieve Signature for Bot Profile Configuration

```hcl
data "citrixadc_botsignature" "bot_sigs" {
  name = "production_bot_signatures"
}

# Use the retrieved signature in bot profile configuration
output "bot_signature_source" {
  value = data.citrixadc_botsignature.bot_sigs.src
}
```

### Reference Signature for Validation

```hcl
data "citrixadc_botsignature" "existing_signature" {
  name = "existing_bot_signature"
}

# Verify signature exists before creating dependent resources
resource "citrixadc_botprofile" "app_bot_profile" {
  name = "app_profile"
  signature = data.citrixadc_botsignature.existing_signature.name
  # ... other profile configuration
}
```

## Notes

- Bot signatures are used to identify and categorize bots accessing your applications
- The signature file should be in JSON format and contain bot detection rules
- Make sure the signature file is available at the specified source location before importing
