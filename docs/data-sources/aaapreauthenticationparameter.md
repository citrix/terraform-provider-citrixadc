---
subcategory: "AAA"
---

# Data Source `aaapreauthenticationparameter`

The aaapreauthenticationparameter data source allows you to retrieve information about AAA preauthentication parameters configuration.


## Example usage

```terraform
data "citrixadc_aaapreauthenticationparameter" "tf_aaapreauthenticationparameter" {
}

output "preauthenticationaction" {
  value = data.citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter.preauthenticationaction
}

output "deletefiles" {
  value = data.citrixadc_aaapreauthenticationparameter.tf_aaapreauthenticationparameter.deletefiles
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `preauthenticationaction` - Deny or allow login after endpoint analysis (EPA) results. Possible values: [ ALLOW, DENY ]
* `deletefiles` - String specifying the path(s) to the file(s) to be deleted by the endpoint analysis (EPA) tool.
* `rule` - Expression that the Citrix ADC evaluates to allow or deny the user from logging on.
* `killprocess` - String specifying the name of a process to be terminated by the endpoint analysis (EPA) tool.

## Attribute Reference

* `id` - The id of the aaapreauthenticationparameter. It is a system-generated identifier.
