---
subcategory: "LSN"
---

# Resource: lsnappsprofile_port_binding

The lsnappsprofile_port_binding resource is used to createlsnappsprofile_port_binding.


## Example usage

```hcl
resource "citrixadc_lsnappsprofile_port_binding" "tf_lsnappsprofile_port_binding" {
  appsprofilename = "my_lsn_profile"
  lsnport         = "80"
}

```


## Argument Reference

* `appsprofilename` - (Optional) Name for the LSN application profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN application profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn application profile1" or 'lsn application profile1').
* `lsnport` - (Optional) Port numbers or range of port numbers to match against the destination port of the incoming packet from a subscriber. When the destination port is matched, the LSN application profile is applied for the LSN session. Separate a range of ports with a hyphen. For example, 40-90.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnappsprofile_port_binding. It is the concatenation of `appsprofilename` and `lsnport` attributes separated by a comma.


## Import

A lsnappsprofile_port_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnappsprofile_port_binding.tf_lsnappsprofile_port_binding my_lsn_profile,80
```
