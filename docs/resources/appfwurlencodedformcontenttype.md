---
subcategory: "Application Firewall"
---

# Resource: appfwurlencodedformcontenttype

The appfwurlencodedformcontenttype resource is used to create appfw urlencoded form contenttype.


## Example usage

```hcl
resource "citrixadc_appfwurlencodedformcontenttype" "tf_urlencodedformcontenttype" {
  urlencodedformcontenttypevalue = "tf_urlencodedformcontenttype"
  isregex                        = "NOTREGEX"
}
```


## Argument Reference

* `urlencodedformcontenttypevalue` - (Required) Content type to be classified as urlencoded form
* `isregex` - (Optional) Is urlencoded form content type a regular expression?


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwurlencodedformcontenttype. It has the same value as the `urlencodedformcontenttypevalue` attribute.


## Import

A appfwurlencodedformcontenttype can be imported using its urlencodedformcontenttypevalue, e.g.

```shell
terraform import citrixadc_appfwurlencodedformcontenttype.tf_urlencodedformcontenttype tf_urlencodedformcontenttype
```
