---
subcategory: "Cloud"
---

# Resource: cloudparameter

Configures the global Citrix Cloud / NGS Connector parameters on the Citrix ADC. These settings identify the appliance to Citrix Cloud (customer, instance, and resource location) and tell the SDProxy connector which controller to reach, so the ADC can register with and be managed from Citrix Cloud.

This is a singleton resource: a single `cloudparameter` configuration always exists on the appliance. Creating this resource updates the existing global configuration rather than adding a new object, so there is no delete operation and no name key.


## Example usage

```hcl
resource "citrixadc_cloudparameter" "tf_cloudparameter" {
  controllerfqdn     = "connector.citrixcloud.example.com"
  controllerport     = 443
  customerid         = "citrixcustomer1"
  instanceid         = "00000000-0000-0000-0000-000000000000"
  resourcelocation   = "11111111-1111-1111-1111-111111111111"
  activationcode     = "abcd-1234-efgh-5678"
  deployment         = "Production"
  connectorresidence = "Onprem"
}
```


## Argument Reference

* `controllerfqdn` - (Optional) FQDN of the controller to which the Citrix ADC SDProxy connects. Maximum length =  255
* `controllerport` - (Optional) Port number of the controller to which the Citrix ADC SDProxy connects. Minimum value =  1 Maximum value =  65535
* `customerid` - (Optional) Customer ID of the Citrix Cloud customer. Maximum length =  255
* `instanceid` - (Optional) Instance ID of the customer provided by Trust. Maximum length =  255
* `resourcelocation` - (Optional) Resource Location of the customer provided by Trust. Maximum length =  255
* `deployment` - (Optional) Describes whether the customer is a Staging, Production, or Dev Citrix Cloud customer. Possible values: [ Production, Staging, Dev ]
* `connectorresidence` - (Optional) Identifies where the connector is located. Possible values: [ None, Onprem, Aws, Azure, Cpx ]
* `activationcode` - (Optional) Activation code for the NGS Connector instance. Maximum length =  255. Note: This value is write-only on the appliance. The NITRO GET/show operation never returns it, so the provider does not read it back from the ADC; instead it preserves the value you configured in Terraform state. Because it cannot be read back, the provider cannot detect out-of-band changes to this value.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudparameter. It is set to `cloudparameter-config`.


## Import

A cloudparameter can be imported using its id, e.g.

```shell
terraform import citrixadc_cloudparameter.tf_cloudparameter cloudparameter-config
```
