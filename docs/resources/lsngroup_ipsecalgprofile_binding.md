---
subcategory: "LSN"
---

# Resource: lsngroup_ipsecalgprofile_binding

Associates an IPSec ALG profile with a Large Scale NAT (LSN) group so that IPSec (IKE/ESP) traffic traversing that group is handled by the ALG. Binding the profile applies its IPSec ALG settings (such as connection failover and IKE session timeout) to all subscribers served by the LSN group.


## Example usage

```hcl
resource "citrixadc_ipsecalgprofile" "tf_ipsecalgprofile" {
  name             = "ipsecalgprofile1"
  ikesessiontimeout = 60
  espsessiontimeout = 60
  espgatetimeout    = 30
}

resource "citrixadc_lsngroup" "tf_lsngroup" {
  groupname  = "lsngroup1"
  clientname = "lsnclient1"
  nattype    = "DYNAMIC"
}

resource "citrixadc_lsngroup_ipsecalgprofile_binding" "tf_binding" {
  groupname       = citrixadc_lsngroup.tf_lsngroup.groupname
  ipsecalgprofile = citrixadc_ipsecalgprofile.tf_ipsecalgprofile.name
}
```


## Argument Reference

* `groupname` - (Required) Name for the LSN group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN group is created. Changing this value forces a new resource to be created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, `"lsn group1"` or `'lsn group1'`).
* `ipsecalgprofile` - (Required) Name of the IPSec ALG profile to bind to the specified LSN group. Changing this value forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the lsngroup_ipsecalgprofile_binding resource. It is a comma-separated string of `key:value` pairs (URL-encoded), composed of the `groupname` and `ipsecalgprofile` attributes — for example, `groupname:lsngroup1,ipsecalgprofile:ipsecalgprofile1`.


## Import

A lsngroup_ipsecalgprofile_binding can be imported using its ID, which is the concatenation of the `groupname` and `ipsecalgprofile` attributes formatted as `key:value` pairs separated by a comma, e.g.

```shell
terraform import citrixadc_lsngroup_ipsecalgprofile_binding.tf_binding "groupname:lsngroup1,ipsecalgprofile:ipsecalgprofile1"
```
