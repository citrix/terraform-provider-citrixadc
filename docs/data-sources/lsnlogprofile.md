---
subcategory: "LSN"
---

# Data Source `lsnlogprofile`

The lsnlogprofile data source allows you to retrieve information about LSN (Large Scale NAT) logging profiles.


## Example usage

```terraform
data "citrixadc_lsnlogprofile" "tf_lsnlogprofile" {
  logprofilename = "my_lsn_logprofile"
}

output "logsubscrinfo" {
  value = data.citrixadc_lsnlogprofile.tf_lsnlogprofile.logsubscrinfo
}

output "logcompact" {
  value = data.citrixadc_lsnlogprofile.tf_lsnlogprofile.logcompact
}

output "logipfix" {
  value = data.citrixadc_lsnlogprofile.tf_lsnlogprofile.logipfix
}
```


## Argument Reference

* `logprofilename` - (Required) The name of the logging Profile.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `analyticsprofile` - Name of the Analytics Profile attached to this lsn profile.
* `logcompact` - Logs in Compact Logging format if option is enabled.
* `logipfix` - Logs in IPFIX  format if option is enabled.
* `logsessdeletion` - LSN Session deletion will not be logged if disabled.
* `logsubscrinfo` - Subscriber ID information is logged if option is enabled.

## Attribute Reference

* `id` - The id of the lsnlogprofile. It has the same value as the `logprofilename` attribute.


## Import

A lsnlogprofile can be imported using its logprofilename, e.g.

```shell
terraform import citrixadc_lsnlogprofile.tf_lsnlogprofile my_lsn_logprofile
```
