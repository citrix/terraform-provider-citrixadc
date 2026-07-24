---
subcategory: "Cloud"
---

# Resource: cloudngsparameter

This resource is used to manage the global Citrix Gateway Service NextGen Service parameters.


## Example usage

```hcl
resource "citrixadc_cloudngsparameter" "tf_cloudngsparameter" {
  blockonallowedngstktprof   = "NO"
  allowedudtversion          = "V4"
  csvserverticketingdecouple = "NO"
  allowdtls12                = "NO"
}
```


## Argument Reference

* `blockonallowedngstktprof` - (Optional) Enables blocking connections authenticated with a ticket created by an entity not whitelisted in `allowedngstktprofile`. Defaults to `"NO"`. Possible values: [ YES, NO ]
* `allowedudtversion` - (Optional) Enables the required UDT version for EDT connections in the CGS deployment. Defaults to `"V4"`. Possible values: [ V4, V5, V6, V7 ]
* `csvserverticketingdecouple` - (Optional) Enables decoupling the content-switching virtual server (CSVSERVER) state from the ticketing service state in the CGS deployment. Defaults to `"NO"`. Possible values: [ YES, NO ]
* `allowdtls12` - (Optional) Enables DTLS 1.2 for client connections on CGS. Defaults to `"NO"`. Possible values: [ YES, NO ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudngsparameter. It is set to `cloudngsparameter-config`.


## Import

A cloudngsparameter can be imported using its id, e.g.

```shell
terraform import citrixadc_cloudngsparameter.tf_cloudngsparameter cloudngsparameter-config
```
