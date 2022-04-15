---
subcategory: "Application Firewall"
---

# Resource: appfwconfidfield

The appfwconfidfield resource is used to create appfw confidfield.


## Example usage

```hcl
resource "citrixadc_appfwconfidfield" "tf_confidfield" {
  fieldname = "tf_confidfield"
  url       = "www.example.com/"
  isregex   = "REGEX"
  comment   = "Testing"
  state     = "DISABLED"
}
```


## Argument Reference

* `fieldname` - (Required) Name of the form field to designate as confidential.
* `url` - (Required) URL of the web page that contains the web form.
* `comment` - (Optional) Any comments to preserve information about the form field designation.
* `isregex` - (Optional) Method of specifying the form field name. Available settings function as follows: * REGEX. Form field is a regular expression. * NOTREGEX. Form field is a literal string.
* `state` - (Optional) Enable or disable the confidential field designation.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the appfwconfidfield. It is the concatenation of `fieldname` and `url` attribute.


## Import

A appfwconfidfield can be imported using its name, e.g.

```shell
terraform import citrixadc_appfwconfidfield.tf_confidfield tf_confidfield,www.example.com/
```
