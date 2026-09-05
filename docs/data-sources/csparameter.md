---
subcategory: "Content Switching"
---

# Data Source: csparameter

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
* `id` - The id of the csparameter. It is a system-generated identifier.

### Read-only csparameter metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_csparameter` resource). They are GET-only/Computed, and any attribute the appliance does not return is `null`.

* `builtin` - Flag to determine if CS param is built-in or not (for example `MODIFIABLE`, `DELETABLE`, `IMMUTABLE`, `PARTITION_ALL`). A list of strings.
* `feature` - The feature to be checked while applying this config.
