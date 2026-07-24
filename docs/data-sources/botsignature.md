---
subcategory: "Bot"
---

# Data Source: botsignature

The botsignature data source allows you to retrieve information about a bot signature.


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
