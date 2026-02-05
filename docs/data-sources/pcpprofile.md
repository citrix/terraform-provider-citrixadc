---
subcategory: "PCP"
---

# Data Source `pcpprofile`

The pcpprofile data source allows you to retrieve information about PCP (Port Control Protocol) profiles.


## Example usage

```terraform
data "citrixadc_pcpprofile" "tf_pcpprofile" {
  name = "my_pcpprofile"
}

output "mapping" {
  value = data.citrixadc_pcpprofile.tf_pcpprofile.mapping
}

output "peer" {
  value = data.citrixadc_pcpprofile.tf_pcpprofile.peer
}
```


## Argument Reference

* `name` - (Required) Name for the PCP Profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my pcpProfile" or my pcpProfile).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `announcemulticount` - Integer value that identify the number announce message to be send.
* `mapping` - This argument is for enabling/disabling the MAP opcode  of current PCP Profile
* `maxmaplife` - Integer value that identify the maximum mapping lifetime (in seconds) for a pcp profile. default(86400s = 24Hours).
* `minmaplife` - Integer value that identify the minimum mapping lifetime (in seconds) for a pcp profile. default(120s)
* `peer` - This argument is for enabling/disabling the PEER opcode of current PCP Profile
* `thirdparty` - This argument is for enabling/disabling the THIRD PARTY opcode of current PCP Profile

## Attribute Reference

* `id` - The id of the pcpprofile. It has the same value as the `name` attribute.


## Import

A pcpprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_pcpprofile.tf_pcpprofile my_pcpprofile
```

