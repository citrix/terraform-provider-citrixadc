---
subcategory: "Transform"
---

# Resource: transformpolicy

The transformpolicy resource is used to create transform policies


## Example usage

```hcl
resource "citrixadc_transformprofile" "tf_trans_profile" {
  name    = "tf_trans_profile"
  comment = "Some comment"
}

resource "citrixadc_transformaction" "tf_trans_action1" {
  name        = "tf_trans_action1"
  profilename = citrixadc_transformprofile.tf_trans_profile.name
  priority    = 100
  requrlfrom  = "http://m3.mydomain.com/(.*)"
  requrlinto  = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom  = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto  = "https://m3.mydomain.com/$1"
}

resource "citrixadc_transformpolicy" "tf_trans_policy" {
  name        = "tf_trans_policy"
  profilename = citrixadc_transformprofile.tf_trans_profile.name
  rule        = "http.REQ.URL.CONTAINS(\"test_url\")"
}
```


## Argument Reference

* `name` - (Required) Name for the URL Transformation policy. Must begin with a letter, number, or the underscore character (\_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the URL Transformation policy is added.
* `rule` - (Optional) Expression, or name of a named expression, against which to evaluate traffic. The following requirements apply only to the Citrix ADC CLI: * If the expression includes blank spaces, the entire expression must be enclosed in double quotation marks. * If the expression itself includes double quotation marks, you must escape the quotations by using the \ character. * Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.
* `profilename` - (Optional) Name of the URL Transformation profile to use to transform requests and responses that match the policy.
* `comment` - (Optional) Any comments to preserve information about this URL Transformation policy.
* `logaction` - (Optional) Log server to use to log connections that match this policy.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the transformpolicy. It has the same value as the `name` attribute.


## Import

A transformpolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_transformpolicy.tf_trans_policy tf_trans_policy
```
