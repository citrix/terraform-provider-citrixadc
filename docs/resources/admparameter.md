---
subcategory: "Adm"
---

# Resource: admparameter

The admparameter resource is used to update admparameter.


## Example usage

```hcl
resource "citrixadc_admparameter" "tf_admparameter" {
  admserviceconnect = "DISABLED"
}
```


## Argument Reference

* `admserviceconnect` - (Optional) Parameter to enable/disable Citrix ADM Service Connect. This feature helps you discover your Citrix ADC instances effortlessly on Citrix ADM service and get insights and curated machine learning based recommendations for applications and Citrix ADC infrastructure. This feature lets the Citrix ADC instance automatically send system, usage and telemetry data to Citrix ADM service. View here [https://docs.citrix.com/en-us/citrix-adc/13/data-governance.html] to learn more about this feature. Use of this feature is subject to the Citrix End User ServiceAgreement. View here [https://www.citrix.com/buy/licensing/agreements.html]. Possible values: [ ENABLED, DISABLED ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the admparameter. It is a unique string prefixed with  `tf-admparameter-` attribute.
