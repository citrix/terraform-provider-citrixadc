---
subcategory: "Content Switching"
---

# Data Source `csparameter`

The csparameter data source allows you to retrieve information about content switching parameters configuration.


## Example usage

```terraform
data "citrixadc_csparameter" "tf_csparameter" {
}

output "stateupdate" {
  value = data.citrixadc_csparameter.tf_csparameter.stateupdate
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `stateupdate` - Specifies whether the virtual server checks the attached load balancing server for state information. Possible values: `ENABLED`, `DISABLED`.

## Attribute Reference

* `id` - The id of the csparameter. It is a system-generated identifier.
