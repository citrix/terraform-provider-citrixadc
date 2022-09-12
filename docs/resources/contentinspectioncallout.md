---
subcategory: "CI"
---

# Resource: contentinspectioncallout

The contentinspectioncallout resource is used to create contentinspectioncallout.


## Example usage

```hcl
resource "citrixadc_contentinspectioncallout" "tf_contentinspectioncalloout" {
  name        = "my_ci_callout"
  type        = "ICAP"
  profilename = "reqmod-profile"
  servername  = "icapsv1"
  returntype  = "TEXT"
  resultexpr  = "icap.res.header(\"ISTag\")"
}

```


## Argument Reference

* `name` - (Required) Name for the Content Inspection callout. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or callout.
* `type` - (Required) Type of the Content Inspection callout. It must be one of the following: * ICAP - Sends ICAP request to the configured ICAP server.
* `returntype` - (Required) Type of data that the target callout agent returns in response to the callout. Available settings function as follows: * TEXT - Treat the returned value as a text string. * NUM - Treat the returned value as a number. * BOOL - Treat the returned value as a Boolean value. Note: You cannot change the return type after it is set.
* `resultexpr` - (Required) Expression that extracts the callout results from the response sent by the CI callout agent. Must be a response based expression, that is, it must begin with ICAP.RES. The operations in this expression must match the return type. For example, if you configure a return type of TEXT, the result expression must be a text based expression, as in the following example: icap.res.header("ISTag")
* `comment` - (Optional) Any comments to preserve information about this Content Inspection callout.
* `profilename` - (Optional) Name of the Content Inspection profile. The type of the configured profile must match the type specified using -type argument.
* `serverip` - (Optional) IP address of Content Inspection server. Mutually exclusive with the server name parameter.
* `servername` - (Optional) Name of the load balancing or content switching virtual server or service to which the Content Inspection request is issued. Mutually exclusive with server IP address and port parameters. The service type must be TCP or SSL_TCP. If there are vservers and services with the same name, then vserver is selected.
* `serverport` - (Optional) Port of the Content Inspection server.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the contentinspectioncallout. It has the same value as the `name` attribute.


## Import

A contentinspectioncallout can be imported using its name, e.g.

```shell
terraform import citrixadc_contentinspectioncallout.tf_contentinspectioncallout my_ci_callout
```
