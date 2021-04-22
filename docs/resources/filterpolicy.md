---
subcategory: "Filter"
---

# Resource: filterpolicy

The filterpolicy resource is used to create filter policies.


## Example usage

```hcl
resource "citrixadc_filterpolicy" "tf_filterpolicy" {
    name = "tf_filterpolicy"
    reqaction = "DROP"
    rule = "REQ.HTTP.URL CONTAINS http://abcd.com"
}
```


## Argument Reference

* `name` - (Required) Name for the filtering action. Must begin with a letter, number, or the underscore character (\_). Other characters allowed, after the first character, are the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), and colon (:) characters. Choose a name that helps identify the type of action. The name cannot be updated after the policy is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my policy" or 'my policy').
* `rule` - (Optional) Citrix ADC classic expression specifying the type of connections that match this policy.
* `reqaction` - (Optional) Name of the action to be performed on requests that match the policy. Cannot be specified if the rule includes condition to be evaluated for responses.
* `resaction` - (Optional) The action to be performed on the response. The string value can be a filter action created filter action or a built-in action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the filterpolicy. It has the same value as the `name` attribute.


## Import

A filterpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_filterpolicy.tf_filterpolicy tf_filterpolicy
```
