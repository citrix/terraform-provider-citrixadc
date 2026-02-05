---
subcategory: "NS"
---

# Data Source `nsservicepath`

The nsservicepath data source allows you to retrieve information about a service path.


## Example usage

```terraform
data "citrixadc_nsservicepath" "my_servicepath" {
  servicepathname = "my_servicepath"
}

output "servicepathname" {
  value = data.citrixadc_nsservicepath.my_servicepath.servicepathname
}
```


## Argument Reference

* `servicepathname` - (Required) Name for the Service path. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsservicepath. It has the same value as the `servicepathname` attribute.
