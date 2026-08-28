---
subcategory: "DPS"
---

# Resource: dpsparameter

The dpsparameter resource is used to configure the DPS (Citrix Cloud connectivity)
parameters. This is a singleton (global) configuration resource.


## Example usage

```hcl
resource "citrixadc_dpsparameter" "tf_dpsparameter" {
  customerid = "customer123"
  deployment = "COMMERCIAL"
  serviceurl = "https://example.citrixcloud.net"
}
```


## Argument Reference

* `customerid` - (Optional) Customer ID of the Citrix Cloud customer. Minimum length = 1. Default value: "None"
* `deployment` - (Optional) Describes if the customer is connecting to Commerical/JapanCloud/Gov Citrix Cloud customer. Possible values: [ COMMERCIAL, GOV, JAPANCLOUD ]. Default value: COMMERCIAL
* `serviceurl` - (Optional) Service URL of the Citrix Cloud customer. Minimum length = 1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the dpsparameter. Because this is a singleton resource, it has a fixed identifier.
