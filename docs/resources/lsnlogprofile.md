---
subcategory: "Lsn"
---

# Resource: lsnlogprofile

The lsnlogprofile resource is used to create lsnlogprofile.


## Example usage

```hcl
resource "citrixadc_lsnlogprofile" "tf_lsnlogprofile" {
  logprofilename = "my_lsn_logprofile"
  logsubscrinfo   = "ENABLED"
  logcompact      = "ENABLED"
  logipfix        = "ENABLED"
}

```


## Argument Reference

* `logprofilename` - (Required) The name of the logging Profile.
* `analyticsprofile` - (Optional) Name of the Analytics Profile attached to this lsn profile.
* `logcompact` - (Optional) Logs in Compact Logging format if option is enabled.
* `logipfix` - (Optional) Logs in IPFIX  format if option is enabled.
* `logsessdeletion` - (Optional) LSN Session deletion will not be logged if disabled.
* `logsubscrinfo` - (Optional) Subscriber ID information is logged if option is enabled.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnlogprofile. It has the same value as the `logprofilename` attribute.


## Import

A lsnlogprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnlogprofile.tf_lsnlogprofile my_lsn_logprofile
```
