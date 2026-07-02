---
subcategory: "Cloud"
---

# Data Source: cloudngsparameter

The cloudngsparameter data source allows you to retrieve the global Citrix Gateway Service (CGS) NextGen Service parameters configured on the Citrix ADC. Because this is a singleton configuration object, no lookup argument is required.


## Example usage

```terraform
data "citrixadc_cloudngsparameter" "tf_cloudngsparameter" {
}

output "allowedudtversion" {
  value = data.citrixadc_cloudngsparameter.tf_cloudngsparameter.allowedudtversion
}

output "allowdtls12" {
  value = data.citrixadc_cloudngsparameter.tf_cloudngsparameter.allowdtls12
}
```


## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `blockonallowedngstktprof` - Whether blocking connections authenticated with a ticket created by an entity not whitelisted in `allowedngstktprofile` is enabled. Possible values: [ YES, NO ]
* `allowedudtversion` - The required UDT version for EDT connections in the CGS deployment. Possible values: [ V4, V5, V6, V7 ]
* `csvserverticketingdecouple` - Whether decoupling the content-switching virtual server (CSVSERVER) state from the ticketing service state is enabled. Possible values: [ YES, NO ]
* `allowdtls12` - Whether DTLS 1.2 for client connections on CGS is enabled. Possible values: [ YES, NO ]
* `id` - The id of the cloudngsparameter. It is a fixed string set to `cloudngsparameter-config`.
