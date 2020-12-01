---
subcategory: "Policy"
---

# Resource: policyexpression

The policyexpression resource is used to create policy expressions.


## Example usage

```hcl
resource "citrixadc_policyexpression" "tf_policyexpression" {
    name = "tf_policyexrpession"
    value = "HTTP.REQ.URL.SUFFIX.EQ(\"cgi\")"
    comment = "comment"
}
```


## Argument Reference

* `name` - (Optional) Unique name for the expression. Not case sensitive. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or HTTP callout.
* `value` - (Optional) Expression string. For example: http.req.body(100).contains("this").
* `description` - (Optional) Description for the expression.
* `comment` - (Optional) Any comments associated with the expression. Displayed upon viewing the policy expression.
* `clientsecuritymessage` - (Optional) Message to display if the expression fails. Allowed for classic end-point check expressions only.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policyexpression. It has the same value as the `name` attribute.


## Import

A policyexpression can be imported using its name, e.g.

```shell
terraform import citrixadc_policyexpression.tf_policyexrpession tf_policyexrpession
```
