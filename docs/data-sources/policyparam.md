---
subcategory: "Policy"
---

# Data Source `policyparam`

The policyparam data source allows you to retrieve information about the policy parameter configuration.

## Example usage

```terraform
data "citrixadc_policyparam" "tf_policyparam" {
}

output "timeout" {
  value = data.citrixadc_policyparam.tf_policyparam.timeout
}
```

## Argument Reference

This data source does not require any arguments.

## Attribute Reference

The following attributes are available:

* `id` - The id of the policyparam resource.
* `timeout` - Maximum time in milliseconds to allow for processing expressions and policies without interruption. If the timeout is reached then the evaluation causes an UNDEF to be raised and no further processing is performed.
