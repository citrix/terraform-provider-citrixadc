---
subcategory: "Bot"
---

# Resource: botprofile_kmdetectionexpr_binding

The botprofile_kmdetectionexpr_binding resource is used to bind a keyboard-mouse (KM) based detection expression to a bot profile. The bound expression injects a JavaScript payload for keyboard-mouse detection when its condition evaluates to true, allowing the bot profile to distinguish human interaction from automated traffic.


## Example usage

```hcl
resource "citrixadc_botprofile" "tf_botprofile" {
  name                     = "tf_botprofile"
  errorurl                 = "http://www.citrix.com"
  trapurl                  = "/http://www.citrix.com"
  comment                  = "tf_botprofile comment"
  bot_enable_white_list    = "ON"
  bot_enable_black_list    = "ON"
  bot_enable_rate_limit    = "ON"
  devicefingerprint        = "ON"
  devicefingerprintaction  = ["LOG", "RESET"]
  bot_enable_ip_reputation = "ON"
  trap                     = "ON"
  trapaction               = ["LOG", "RESET"]
  bot_enable_tps           = "ON"
}

resource "citrixadc_botprofile_kmdetectionexpr_binding" "tf_binding" {
  name                     = citrixadc_botprofile.tf_botprofile.name
  kmdetectionexpr          = true
  bot_km_expression_name   = "tf_kmexpr"
  bot_km_expression_value  = "HTTP.REQ.URL.CONTAINS(\"/login\")"
  bot_km_detection_enabled = "ON"
  logmessage               = "KM detection triggered"
  bot_bind_comment         = "KM detection binding"
}
```


## Argument Reference

* `name` - (Required) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `kmdetectionexpr` - (Required) Keyboard-mouse based detection binding. For each name, only one binding is allowed. To update the values of an existing binding, first unbind the binding, then bind again with the new values. A maximum of 30 bindings can be configured per profile.
* `bot_km_expression_name` - (Required) Name of the keyboard-mouse expression object.
* `bot_km_expression_value` - (Required) JavaScript file for keyboard-mouse detection, inserted if the result of the expression is true.
* `bot_km_detection_enabled` - (Optional) Enable or disable the keyboard-mouse based binding. Possible values: [ ON, OFF ]. Defaults to `OFF`.
* `logmessage` - (Optional) Message to be logged for this binding.
* `bot_bind_comment` - (Optional) Any comments about this binding.

~> **NOTE:** This binding is immutable. Changing any attribute forces recreation (unbind and re-bind) of the resource.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile_kmdetectionexpr_binding. It is the concatenation of the `name` and `bot_km_expression_name` attributes separated by a comma.


## Import

A botprofile_kmdetectionexpr_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_botprofile_kmdetectionexpr_binding.tf_binding tf_botprofile,tf_kmexpr
```
