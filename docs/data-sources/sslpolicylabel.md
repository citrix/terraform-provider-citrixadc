---
subcategory: "SSL"
---

# Data Source: sslpolicylabel

The sslpolicylabel data source allows you to retrieve information about an existing sslpolicylabel.


## Example usage

```terraform
data "citrixadc_sslpolicylabel" "tf_sslpolicylabel" {
  labelname = "tf_sslpolicylabel"
}

output "labelname" {
  value = data.citrixadc_sslpolicylabel.tf_sslpolicylabel.labelname
}

output "type" {
  value = data.citrixadc_sslpolicylabel.tf_sslpolicylabel.type
}
```


## Argument Reference

* `labelname` - (Required) Name for the SSL policy label. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy label is created.

The following requirement applies only to the Citrix ADC CLI:
If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my label" or 'my label').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the sslpolicylabel. It has the same value as the `labelname` attribute.
* `type` - Type of policies that the policy label can contain.

### Read-only sslpolicylabel metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_sslpolicylabel` resource). Any attribute the appliance does not return is `null`.

* `numpol` - Number of polices bound to label.
* `hits` - Number of times policy label was invoked.
* `policyname` - Name of the SSL policy to bind to the policy label.
* `priority` - Specifies the priority of the policy.
* `gotopriorityexpression` - Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.
* `invoke` - Invoke flag.
* `labeltype` - Type of policy label invocation. Possible values = vserver, service, policylabel.
* `invoke_labelname` - Name of the label to invoke if the current policy rule evaluates to TRUE.
* `flowtype` - Flowtype of the bound SSL policy.
* `description` - Description of the policylabel.
