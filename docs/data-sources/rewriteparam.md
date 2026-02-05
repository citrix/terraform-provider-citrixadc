---
subcategory: "Rewrite"
---

# Data Source: citrixadc_rewriteparam

This data source retrieves the global rewrite parameter configuration.

## Example Usage

```hcl
data "citrixadc_rewriteparam" "test" {
}

output "rewrite_timeout" {
  value = data.citrixadc_rewriteparam.test.timeout
}
```

## Argument Reference

This data source takes no arguments.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the rewrite parameter configuration.
* `timeout` - Maximum time in milliseconds to allow for processing all the policies and their selected actions without interruption. If the timeout is reached then the evaluation causes an UNDEF to be raised and no further processing is performed.
* `undefaction` - Action to perform if the result of policy evaluation is undefined (UNDEF). Possible values: NOREWRITE, RESET, DROP.
