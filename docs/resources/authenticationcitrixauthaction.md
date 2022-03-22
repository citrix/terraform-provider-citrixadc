---
subcategory: "Authentication"
---

# Resource: authenticationcitrixauthaction

The authenticationcitrixauthaction resource is used to create authentication citrixauthaction resource.


## Example usage

```hcl
resource "citrixadc_authenticationcitrixauthaction" "tf_citrixauthaction" {
  name               = "tf_citrixauthaction"
  authenticationtype = "CITRIXCONNECTOR"
  authentication     = "DISABLED"
}
```


## Argument Reference

* `name` - (Required) Name for the new Citrix Authentication action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after an action is created.  The following requirement applies only to the NetScaler CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action').
* `authentication` - (Optional) Authentication needs to be disabled for searching user object without performing authentication.
* `authenticationtype` - (Optional) Type of the Citrix Authentication implementation. Default implementation uses Citrix Cloud Connector.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationcitrixauthaction. It has the same value as the `name` attribute.


## Import

A authenticationcitrixauthaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationcitrixauthaction.tf_citrixauthaction tf_citrixauthaction
```
