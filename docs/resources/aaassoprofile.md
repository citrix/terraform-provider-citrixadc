---
subcategory: "AAA"
---

# Resource: aaassoprofile

The aaassoprofile resource is used to create aaassoprofile.


## Example usage

```hcl
resource "citrixadc_aaassoprofile" "tf_aaassoprofile" {
  name = "myssoprofile"
  username = "john"
  password = "my_password"
}
```


## Argument Reference

* `name` - (Required) Name for the SSO Profile. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after a SSO Profile is created. The following requirement applies only to the NetScaler CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my action" or 'my action'). Minimum length =  1
* `username` - (Required) Name for the user. Must begin with a letter, number, or the underscore (_) character, and must contain only alphanumeric, hyphen (-), period (.), hash (#), space ( ), at (@), equal (=), colon (:), and underscore characters. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my group" or 'my group'). Minimum length =  1
* `password` - (Required) Password with which the user logs on. Required for Single sign on to  external server. Minimum length =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaassoprofile. It has the same value as the `name` attribute.


## Import

A aaassoprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_aaassoprofile.tf_aaassoprofile myssoprofile
```
