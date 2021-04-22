---
subcategory: "Transform"
---

# Resource: transformaction

The transformaction resource is used to create transform actions.


## Example usage

```hcl
resource "citrixadc_transformprofile" "tf_trans_profile" {
  name    = "tf_trans_profile"
  comment = "Some comment"
}

resource "citrixadc_transformaction" "tf_trans_action" {
  name        = "tf_trans_action"
  profilename = citrixadc_transformprofile.tf_trans_profile.name
  priority    = 100
  requrlfrom  = "http://m3.mydomain.com/(.*)"
  requrlinto  = "https://exp-proxy-v1.api.mydomain.com/$1"
  resurlfrom  = "https://exp-proxy-v1.api.mydomain.com/(.*)"
  resurlinto  = "https://m3.mydomain.com/$1"
}
```


## Argument Reference

* `name` - (Required) Name for the URL transformation action. Must begin with a letter, number, or the underscore character (\_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL Transformation action is added. 
* `profilename` - (Optional) Name of the URL Transformation profile with which to associate this action.
* `priority` - (Optional) Positive integer specifying the priority of the action within the profile. A lower number specifies a higher priority. Must be unique within the list of actions bound to the profile. Policies are evaluated in the order of their priority numbers, and the first policy that matches is applied.
* `state` - (Optional) Enable or disable this action. Possible values: [ ENABLED, DISABLED ]
* `requrlfrom` - (Optional) PCRE-format regular expression that describes the request URL pattern to be transformed.
* `requrlinto` - (Optional) PCRE-format regular expression that describes the transformation to be performed on URLs that match the reqUrlFrom pattern.
* `resurlfrom` - (Optional) PCRE-format regular expression that describes the response URL pattern to be transformed.
* `resurlinto` - (Optional) PCRE-format regular expression that describes the transformation to be performed on URLs that match the resUrlFrom pattern.
* `cookiedomainfrom` - (Optional) Pattern that matches the domain to be transformed in Set-Cookie headers.
* `cookiedomaininto` - (Optional) PCRE-format regular expression that describes the transformation to be performed on cookie domains that match the cookieDomainFrom pattern. NOTE: The cookie domain to be transformed is extracted from the request.
* `comment` - (Optional) Any comments to preserve information about this URL Transformation action.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the transformaction. It has the same value as the `name` attribute.


## Import

A transformaction can be imported using its name, e.g.

```shell
terraform import citrixadc_transformaction.tf_trans_action tf_trans_action
```
