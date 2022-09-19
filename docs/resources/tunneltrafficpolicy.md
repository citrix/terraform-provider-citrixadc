---
subcategory: "Tunnel"
---

# Resource: tunneltrafficpolicy

The tunneltrafficpolicy resource is used to create tunneltrafficpolicy.


## Example usage

```hcl
resource "citrixadc_tunneltrafficpolicy" "tf_tunneltrafficpolicy" {
  name   = "my_tunneltrafficpolicy"
  rule   = "true"
  action = "COMPRESS"
}
```


## Argument Reference

* `name` - (Required) Name for the tunnel traffic policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy)'.
* `rule` - (Required) Expression, against which traffic is evaluated. The following requirements apply only to the Citrix ADC CLI: *  If the expression includes blank spaces, the entire expression must be enclosed in double quotation marks. *  If the expression itself includes double quotation marks, you must escape the quotations by using the \ character.  *  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `action` - (Required) Name of the built-in compression action to associate with the policy.
* `comment` - (Optional) Any comments to preserve information about this policy.
* `logaction` - (Optional) Name of the messagelog action to use for requests that match this policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the tunneltrafficpolicy. It has the same value as the `name` attribute.


## Import

A tunneltrafficpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_tunneltrafficpolicy.tf_tunneltrafficpolicy my_tunneltrafficpolicy
```
