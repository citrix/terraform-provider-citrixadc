---
subcategory: "Bot"
---

# Data Source: botprofile_kmdetectionexpr_binding

The botprofile_kmdetectionexpr_binding data source allows you to retrieve information about the bindings between a bot profile and keyboard-mouse (KM) detection expressions.

## Example usage

```terraform
data "citrixadc_botprofile_kmdetectionexpr_binding" "tf_binding" {
  name                   = "tf_botprofile"
  kmdetectionexpr        = true
  bot_km_expression_name = "tf_kmexpr"
}

output "bot_km_expression_value" {
  value = data.citrixadc_botprofile_kmdetectionexpr_binding.tf_binding.bot_km_expression_value
}

output "bot_km_detection_enabled" {
  value = data.citrixadc_botprofile_kmdetectionexpr_binding.tf_binding.bot_km_detection_enabled
}
```

## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `kmdetectionexpr` - (Required) Keyboard-mouse based detection binding. For each name, only one binding is allowed.
* `bot_km_expression_name` - (Required) Name of the keyboard-mouse expression object.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_kmdetectionexpr_binding. It is the concatenation of the `name` and `bot_km_expression_name` attributes separated by a comma.
* `bot_km_expression_value` - JavaScript file for keyboard-mouse detection, inserted if the result of the expression is true.
* `bot_km_detection_enabled` - Enable or disable the keyboard-mouse based binding.
* `logmessage` - Message to be logged for this binding.
* `bot_bind_comment` - Any comments about this binding.
