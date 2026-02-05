---
subcategory: "Bot"
---

# Data Source: citrixadc_botpolicylabel

Use this data source to retrieve information about an existing Bot Policy Label.

The `citrixadc_botpolicylabel` data source allows you to retrieve details of a bot policy label by its name. This is useful for referencing existing bot policy labels in your Terraform configurations without managing them directly.

## Example usage

```hcl
# Retrieve an existing bot policy label
data "citrixadc_botpolicylabel" "example" {
  labelname = "my_bot_policy_label"
}

# Reference policy label attributes
output "label_comment" {
  value = data.citrixadc_botpolicylabel.example.comment
}

output "label_name" {
  value = data.citrixadc_botpolicylabel.example.labelname
}
```

## Argument Reference

The following arguments are supported:

* `labelname` - (Required) Name for the bot policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the responder policy label is added.

## Attribute Reference

In addition to the argument above, the following attributes are exported:

* `id` - The ID of the bot policy label (same as labelname).

* `comment` - Any comments to preserve information about this bot policy label.

* `newname` - New name for the bot policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Common Use Cases

### Retrieve Policy Label for Bindings

```hcl
data "citrixadc_botpolicylabel" "bot_label" {
  labelname = "bot_protection_label"
}

# Use the policy label in your configuration
output "bot_label_id" {
  value = data.citrixadc_botpolicylabel.bot_label.id
}
```

### Reference Policy Label Details

```hcl
data "citrixadc_botpolicylabel" "existing_label" {
  labelname = "existing_bot_label"
}

# Display label information
output "label_info" {
  value = {
    name    = data.citrixadc_botpolicylabel.existing_label.labelname
    comment = data.citrixadc_botpolicylabel.existing_label.comment
  }
}
```
