---
subcategory: "Cloud"
---

# Resource: cloudngsparameter

Configures the global Citrix Gateway Service (CGS) NextGen Service parameters on the Citrix ADC. Use this resource to tune connection-level behavior for cloud gateway deployments, such as the UDT/EDT version negotiated with clients, DTLS 1.2 support, ticket-profile enforcement, and decoupling of content-switching virtual server state from the ticketing service. This is a singleton resource: a single configuration object always exists on the appliance, so applying this resource updates the existing global parameters rather than creating a new object.


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

* `id` - The id of the cloudngsparameter. Because this is a singleton resource, it is a fixed string set to `cloudngsparameter-config`.
