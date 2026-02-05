---
subcategory: "ICA"
---

# Data Source `icalatencyprofile`

The icalatencyprofile data source allows you to retrieve information about an ICA latency profile configuration.


## Example usage

```terraform
resource "citrixadc_icalatencyprofile" "tf_icalatencyprofile" {
  name                     = "my_ica_latencyprofile"
  l7latencymonitoring      = "ENABLED"
  l7latencythresholdfactor = 120
  l7latencywaittime        = 100
}

data "citrixadc_icalatencyprofile" "tf_icalatencyprofile_ds" {
  name = citrixadc_icalatencyprofile.tf_icalatencyprofile.name
}

output "l7latencymonitoring" {
  value = data.citrixadc_icalatencyprofile.tf_icalatencyprofile_ds.l7latencymonitoring
}

output "l7latencythresholdfactor" {
  value = data.citrixadc_icalatencyprofile.tf_icalatencyprofile_ds.l7latencythresholdfactor
}
```


## Argument Reference

* `name` - (Required) Name for the ICA latencyprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA latency profile is added.

## Attribute Reference

The following attributes are available:

* `id` - The id of the icalatencyprofile. It is a system-generated identifier.
* `l7latencymaxnotifycount` - L7 Latency Max notify Count. This is the upper limit on the number of notifications sent to the Insight Center within an interval where the Latency is above the threshold.
* `l7latencymonitoring` - Enable/Disable L7 Latency monitoring for L7 latency notifications.
* `l7latencynotifyinterval` - L7 Latency Notify Interval. This is the interval at which the Citrix ADC sends out notifications to the Insight Center after the wait time has passed.
* `l7latencythresholdfactor` - L7 Latency threshold factor. This is the factor by which the active latency should be greater than the minimum observed value to determine that the latency is high and may need to be reported.
* `l7latencywaittime` - L7 Latency Wait time. This is the time for which the Citrix ADC waits after the threshold is exceeded before it sends out a Notification to the Insight Center.
* `name` - Name for the ICA latencyprofile.
