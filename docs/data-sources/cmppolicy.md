---
subcategory: "Compression"
---

# Data Source: cmppolicy

The cmppolicy data source allows you to retrieve information about HTTP compression policies.


## Example usage

```terraform
data "citrixadc_cmppolicy" "tf_cmppolicy" {
  name = "my_cmppolicy"
}

output "rule" {
  value = data.citrixadc_cmppolicy.tf_cmppolicy.rule
}

output "resaction" {
  value = data.citrixadc_cmppolicy.tf_cmppolicy.resaction
}
```


## Argument Reference

* `name` - (Required) Name of the HTTP compression policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the policy is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `newname` - New name for the compression policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Choose a name that reflects the function that the policy performs.
* `resaction` - The built-in or user-defined compression action to apply to the response when the policy matches a request or response.
* `rule` - Expression that determines which HTTP requests or responses match the compression policy.
* `id` - The id of the cmppolicy. It has the same value as the `name` attribute.

### Read-only cmppolicy metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_cmppolicy` resource). They are Computed/GET-only, and any attribute the appliance does not return is `null`.

* `reqaction` - The compression action to be performed on requests.
* `hits` - Number of hits.
* `txbytes` - Number of bytes transferred.
* `rxbytes` - Number of bytes received.
* `clientttlb` - Total client TTLB value.
* `clienttransactions` - Number of client transactions.
* `serverttlb` - Total server TTLB value.
* `servertransactions` - Number of server transactions.
* `description` - Description of the policy.
* `builtin` - Flag to determine if compression policy is builtin or not (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this config.
* `isdefault` - A value of true is returned if it is a default policy.

## Import

A cmppolicy can be imported using its name, e.g.

```shell
terraform import citrixadc_cmppolicy.tf_cmppolicy my_cmppolicy
```
