---
subcategory: "NS"
---

# Data Source: nsservicefunction

The nsservicefunction data source allows you to retrieve information about service functions configured on the NetScaler appliance.


## Example usage

```terraform

data "citrixadc_nsservicefunction" "my_servicefunction_data" {
  servicefunctionname = "my_servicefunction"
}

output "servicefunction_vlan" {
  value = data.citrixadc_nsservicefunction.my_servicefunction_data.ingressvlan
}
```


## Argument Reference

* `servicefunctionname` - (Required) Name of the service function to be created. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `ingressvlan` - VLAN ID on which the traffic from service function reaches Citrix ADC.
* `id` - The id of the nsservicefunction. It has the same value as the `servicefunctionname` attribute.
