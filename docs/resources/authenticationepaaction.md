---
subcategory: "Authentication"
---

# Resource: authenticationepaaction

The authenticationepaaction resource is used to create authentication epaaction Resource.


## Example usage

```hcl
resource "citrixadc_authenticationepaaction" "tf_epaaction" {
  name            = "tf_epaaction"
  csecexpr        = "sys.client_expr (\"app_0_MAC-BROWSER_1001_VERSION_<=_10.0.3\")"
  defaultepagroup = "new_group"
  deletefiles     = "old_files"
  killprocess     = "old_process"
}
```


## Argument Reference

* `name` - (Required) Name for the epa action. Must begin with a 	    letter, number, or the underscore character (_), and must consist 	    only of letters, numbers, and the hyphen (-), period (.) pound 	    (#), space ( ), at (@), equals (=), colon (:), and underscore 		    characters. Cannot be changed after epa action is created.The following requirement applies only to the Citrix ADC CLI:If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my aaa action" or 'my aaa action').
* `csecexpr` - (Required) it holds the ClientSecurityExpression to be sent to the client
* `defaultepagroup` - (Optional) This is the default group that is chosen when the EPA check succeeds.
* `deletefiles` - (Optional) String specifying the path(s) and name(s) of the files to be deleted by the endpoint analysis (EPA) tool. Multiple files to be delimited by comma
* `killprocess` - (Optional) String specifying the name of a process to be terminated by the endpoint analysis (EPA) tool. Multiple processes to be delimited by comma
* `quarantinegroup` - (Optional) This is the quarantine group that is chosen when the EPA check fails if configured.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationepaaction. It has the same value as the `name` attribute.


## Import

A authenticationepaaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationepaaction.tf_epaaction tf_epaaction
```
