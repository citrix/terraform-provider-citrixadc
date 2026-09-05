---
subcategory: "Transform"
---

# Data Source: transformpolicy

The transformpolicy data source allows you to retrieve information about a URL Transformation policy.


## Example usage

```terraform
data "citrixadc_transformpolicy" "tf_trans_policy" {
  name = "tf_trans_policy"
}

output "profilename" {
  value = data.citrixadc_transformpolicy.tf_trans_policy.profilename
}

output "rule" {
  value = data.citrixadc_transformpolicy.tf_trans_policy.rule
}
```


## Argument Reference

* `name` - (Required) Name for the URL Transformation policy.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `comment` - Any comments to preserve information about this URL Transformation policy.
* `logaction` - Log server to use to log connections that match this policy.
* `newname` - New name for the policy. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `profilename` - Name of the URL Transformation profile to use to transform requests and responses that match the policy.
* `rule` - Expression, or name of a named expression, against which to evaluate traffic.
* `id` - The id of the transformpolicy. It has the same value as the `name` attribute.

### Read-only transformpolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_transformpolicy` resource). They are GET-only/Computed and are `null` when the appliance does not return them.

* `hits` - Number of hits.
* `isdefault` - A value of true is returned if it is a default transform policy.
* `builtin` - Flag to determine if Transform policy is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL. A list of strings.
* `feature` - The feature to be checked while applying this config.
