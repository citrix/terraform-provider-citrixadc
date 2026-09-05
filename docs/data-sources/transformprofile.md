---
subcategory: "Transform"
---

# Data Source: transformprofile

The transformprofile data source allows you to retrieve information about a URL Transformation profile.


## Example usage

```terraform
data "citrixadc_transformprofile" "tf_trans_profile" {
  name = "tf_trans_profile"
}

output "comment" {
  value = data.citrixadc_transformprofile.tf_trans_profile.comment
}

output "onlytransformabsurlinbody" {
  value = data.citrixadc_transformprofile.tf_trans_profile.onlytransformabsurlinbody
}
```


## Argument Reference

* `name` - (Required) Name for the URL transformation profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Any comments to preserve information about this URL Transformation profile.
* `onlytransformabsurlinbody` - In the HTTP body, transform only absolute URLs. Relative URLs are ignored.
* `type` - Type of transformation. Always URL for URL Transformation profiles.
* `id` - The id of the transformprofile. It has the same value as the `name` attribute.

### Read-only transformprofile metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_transformprofile` resource). They are GET-only/Computed and are `null` when the appliance does not return them.

* `regexforfindingurlinjavascript` - Patclass having regexes to find the URLs in JavaScript.
* `regexforfindingurlincss` - Patclass having regexes to find the URLs in CSS.
* `regexforfindingurlinxcomponent` - Patclass having regexes to find the URLs in X-Component.
* `regexforfindingurlinxml` - Patclass having regexes to find the URLs in XML.
* `additionalreqheaderslist` - Patclass having a list of additional request header names that should transformed.
* `additionalrespheaderslist` - Patclass having a list of additional response header names that should transformed.
