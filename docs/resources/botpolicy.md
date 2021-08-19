---
subcategory: "Bot"
---

# Resource: Botpolicy

The botpolicy resource is used to create a botpolicy.


## Example usage

```hcl
resource "citrixadc_botpolicy" "demo_botpolicy1" {
  name        = "demo_botpolicy1"
  profilename = "BOT_BYPASS"
  rule        = "true"
  comment     = "COMMENT FOR BOTPOLICY"
}
```


## Argument Reference

* `name` - (Required) Name for the bot policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the bot policy is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my bot policy" or 'my bot policy').
* `rule` - (Required) Expression that the policy uses to determine whether to apply bot profile on the specified request.
* `profilename` - (Required) Name of the bot profile to apply if the request matches this bot policy.
* `undefaction` - (Optional) Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an internal error condition.
* `comment` - (Optional) Any type of information about this bot policy.
* `logaction` - (Optional) Name of the messagelog action to use for requests that match this policy.
* `newname` - (Optional) New name for the bot policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my bot policy" or 'my bot policy').


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the Botpolicy. It has the same value as the `name` attribute.


## Import

A botpolicy resource can be imported using its name, e.g.

```shell
terraform import citrixadc_botpolicy.tf_botpolicy tf_botpolicy
```
