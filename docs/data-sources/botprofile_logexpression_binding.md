---
subcategory: "Bot"
---

# Data Source: botprofile_logexpression_binding

The botprofile_logexpression_binding data source allows you to retrieve information about the bindings between botprofile and log expressions.

## Example Usage

```terraform
data "citrixadc_botprofile_logexpression_binding" "tf_binding" {
  name                    = "tf_botprofile"
  bot_log_expression_name = "tf_logname"
}

output "bot_log_expression_value" {
  value = data.citrixadc_botprofile_logexpression_binding.tf_binding.bot_log_expression_value
}

output "bot_bind_comment" {
  value = data.citrixadc_botprofile_logexpression_binding.tf_binding.bot_bind_comment
}
```

## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `bot_log_expression_name` - (Required) Name of the log expression object.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_logexpression_binding. It is the concatenation of the `name` and `bot_log_expression_name` attributes separated by a comma.
* `logexpression` - Log expression binding.
* `bot_bind_comment` - Any comments about this binding.
* `bot_log_expression_enabled` - Enable or disable the log expression binding.
* `bot_log_expression_value` - Expression whose result to be logged when violation happened on the bot profile.
* `logmessage` - Message to be logged for this binding.
