---
subcategory: "Authentication"
---

# Resource: authenticationnoauthaction

The authenticationnoauthaction resource is used to create Authenticationnoauth action resource.


## Example usage

```hcl
resource "citrixadc_authenticationnoauthaction" "tf_noauthaction" {
  name                       = "tf_noauthaction"
  defaultauthenticationgroup = "group"
}
```


## Argument Reference

* `name` - (Required) Name for the new no-authentication action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `defaultauthenticationgroup` - (Optional) This is the group that is added to user sessions that match current policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationnoauthaction. It has the same value as the `name` attribute.


## Import

A authenticationnoauthaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationnoauthaction.tf_noauthaction tf_noauthaction
```
