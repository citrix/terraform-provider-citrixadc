---
subcategory: "NS"
---

# Resource: nsservicepath_nsservicefunction_binding

The nsservicepath_nsservicefunction_binding resource is used to bind nsservice function to nsservice path.


## Example usage

```hcl
resource "citrixadc_nsservicepath" "tf_servicepath" {
  servicepathname = "tf_servicepath"
}
resource "citrixadc_vlan" "tf_vlan" {
  vlanid    = 20
  aliasname = "Management VLAN"
}
resource "citrixadc_nsservicefunction" "tf_servicefunc" {
  servicefunctionname = "tf_servicefunc"
  ingressvlan         = citrixadc_vlan.tf_vlan.vlanid
}
resource "citrixadc_nsservicepath_nsservicefunction_binding" "tf_binding" {
  servicepathname = citrixadc_nsservicepath.tf_servicepath.servicepathname
  servicefunction = citrixadc_nsservicefunction.tf_servicefunc.servicefunctionname
  index           = 2
}
```


## Argument Reference

* `servicepathname` - (Required) Name for the Service path. Must begin with an ASCII alphanumeric or underscore (_) character, and must       contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-)       characters. Minimum length =  1
* `servicefunction` - (Required) List of service functions constituting the chain. Minimum length =  1
* `index` - (Required) The serviceindex of each servicefunction in path. Minimum value =  1


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsservicepath_nsservicefunction_binding. It is the concatenation of `servicepathname` and `servicefunction` attributes separated by comma.


## Import

A nsservicepath_nsservicefunction_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_nsservicepath_nsservicefunction_binding.tf_binding tf_servicepath,tf_servicefunc
```
