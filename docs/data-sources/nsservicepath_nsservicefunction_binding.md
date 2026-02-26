---
subcategory: "NS"
---

# Data Source: nsservicepath_nsservicefunction_binding

The nsservicepath_nsservicefunction_binding data source allows you to retrieve information about the binding between nsservice function and nsservice path.

## Example Usage

```terraform
data "citrixadc_nsservicepath_nsservicefunction_binding" "tf_binding" {
  servicepathname = "tf_servicepath"
  servicefunction = "tf_servicefunc"
}

output "index" {
  value = data.citrixadc_nsservicepath_nsservicefunction_binding.tf_binding.index
}

output "id" {
  value = data.citrixadc_nsservicepath_nsservicefunction_binding.tf_binding.id
}
```

## Argument Reference

* `servicepathname` - (Required) Name for the Service path. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.
* `servicefunction` - (Required) List of service functions constituting the chain.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsservicepath_nsservicefunction_binding. It is a system-generated identifier.
* `index` - The serviceindex of each servicefunction in path.
