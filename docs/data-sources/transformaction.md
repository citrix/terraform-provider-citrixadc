---
subcategory: "Transform"
---

# Data Source: transformaction

The transformaction data source allows you to retrieve information about a URL transformation action.

## Example usage

```terraform
data "citrixadc_transformaction" "tf_trans_action" {
  name = "tf_trans_action"
}

output "profilename" {
  value = data.citrixadc_transformaction.tf_trans_action.profilename
}

output "priority" {
  value = data.citrixadc_transformaction.tf_trans_action.priority
}
```

## Argument Reference

* `name` - (Required) Name for the URL transformation action.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Any comments to preserve information about this URL Transformation action.
* `cookiedomainfrom` - Pattern that matches the domain to be transformed in Set-Cookie headers.
* `cookiedomaininto` - PCRE-format regular expression that describes the transformation to be performed on cookie domains that match the cookieDomainFrom pattern. NOTE: The cookie domain to be transformed is extracted from the request.
* `priority` - Positive integer specifying the priority of the action within the profile. A lower number specifies a higher priority. Must be unique within the list of actions bound to the profile. Policies are evaluated in the order of their priority numbers, and the first policy that matches is applied.
* `profilename` - Name of the URL Transformation profile with which to associate this action.
* `requrlfrom` - PCRE-format regular expression that describes the request URL pattern to be transformed.
* `requrlinto` - PCRE-format regular expression that describes the transformation to be performed on URLs that match the reqUrlFrom pattern.
* `resurlfrom` - PCRE-format regular expression that describes the response URL pattern to be transformed.
* `resurlinto` - PCRE-format regular expression that describes the transformation to be performed on URLs that match the resUrlFrom pattern.
* `state` - Enable or disable this action.

## Attribute Reference

* `id` - The id of the transformaction. It has the same value as the `name` attribute.

## Import

A transformaction can be imported using its name, e.g.

```shell
terraform import citrixadc_transformaction.tf_trans_action tf_trans_action
```
