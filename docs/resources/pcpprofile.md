---
subcategory: "Pcp"
---

# Resource: pcpprofile

The pcpprofile resource is used to create pcpprofile.


## Example usage

```hcl
resource "citrixadc_pcpprofile" "tf_pcpprofile" {
  name               = "my_pcpprofile"
  mapping            = "ENABLED"
  peer               = "ENABLED"
}
```


## Argument Reference

* `name` - (Required) Name for the PCP Profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my pcpProfile" or my pcpProfile).
* `announcemulticount` - (Optional) Integer value that identify the number announce message to be send.
* `mapping` - (Optional) This argument is for enabling/disabling the MAP opcode  of current PCP Profile
* `maxmaplife` - (Optional) Integer value that identify the maximum mapping lifetime (in seconds) for a pcp profile. default(86400s = 24Hours).
* `minmaplife` - (Optional) Integer value that identify the minimum mapping lifetime (in seconds) for a pcp profile. default(120s)
* `peer` - (Optional) This argument is for enabling/disabling the PEER opcode of current PCP Profile
* `thirdparty` - (Optional) This argument is for enabling/disabling the THIRD PARTY opcode of current PCP Profile


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the pcpprofile. It has the same value as the `name` attribute.


## Import

A pcpprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_pcpprofile.tf_pcpprofile my_pcpprofile
```
