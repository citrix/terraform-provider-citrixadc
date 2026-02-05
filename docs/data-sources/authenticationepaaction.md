---
subcategory: "Authentication"
---

# Data Source `authenticationepaaction`

The authenticationepaaction data source allows you to retrieve information about authentication EPA (Endpoint Analysis) actions.


## Example usage

```terraform
data "citrixadc_authenticationepaaction" "tf_epaaction" {
  name = "tf_epaaction"
}

output "csecexpr" {
  value = data.citrixadc_authenticationepaaction.tf_epaaction.csecexpr
}

output "defaultepagroup" {
  value = data.citrixadc_authenticationepaaction.tf_epaaction.defaultepagroup
}
```


## Argument Reference

* `name` - (Required) Name for the epa action. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after epa action is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `csecexpr` - It holds the ClientSecurityExpression to be sent to the client.
* `defaultepagroup` - This is the default group that is chosen when the EPA check succeeds.
* `deletefiles` - String specifying the path(s) and name(s) of the files to be deleted by the endpoint analysis (EPA) tool. Multiple files to be delimited by comma.
* `deviceposture` - Parameter to enable/disable device posture service scan.
* `killprocess` - String specifying the name of a process to be terminated by the endpoint analysis (EPA) tool. Multiple processes to be delimited by comma.
* `quarantinegroup` - This is the quarantine group that is chosen when the EPA check fails if configured.

## Attribute Reference

* `id` - The id of the authenticationepaaction. It has the same value as the `name` attribute.


## Import

A authenticationepaaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationepaaction.tf_epaaction tf_epaaction
```
