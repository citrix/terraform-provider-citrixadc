---
subcategory: "authorization"
---

# Resource: authorizationpolicylabel

The authorizationpolicylabel resource is used to create authorization policy label.


## Example usage

```hcl
resource "citrixadc_authorizationpolicylabel" "authorizationpolicylabel" {
  labelname = "trans_http_url"
}

```


## Argument Reference

* `labelname` - (Required) Name for the new authorization policy label.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authorization policy label" or 'authorization policy label').
* `newname` - (Optional) The new name of the auth policy label


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authorizationpolicylabel. It has the same value as the `labelname` attribute.


## Import

A authorizationpolicylabel can be imported using its name, e.g.

```shell
terraform import citrixadc_authorizationpolicylabel.authorizationpolicylabel trans_http_url
```
