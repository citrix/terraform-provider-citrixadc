---
subcategory: "Ica"
---

# Resource: icalatencyprofile

The icalatencyprofile resource is used to create icalatencyprofile.


## Example usage

```hcl
resource "citrixadc_icalatencyprofile" "tf_icalatencyprofile" {
  name                     = "my_ica_latencyprofile"
  l7latencymonitoring      = "ENABLED"
  l7latencythresholdfactor = 120
  l7latencywaittime        = 100
}
```


## Argument Reference

* `name` - (Required) Name for the ICA latencyprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA latency profile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my ica l7latencyprofile" or 'my ica l7latencyprofile'). Minimum length =  1
* `l7latencymonitoring` - (Optional) Enable/Disable L7 Latency monitoring for L7 latency notifications. Possible values: [ ENABLED, DISABLED ]
* `l7latencythresholdfactor` - (Optional) L7 Latency threshold factor. This is the factor by which the active latency should be greater than the minimum observed value to determine that the latency is high and may need to be reported. Minimum value =  2 Maximum value =  65535
* `l7latencywaittime` - (Optional) L7 Latency Wait time. This is the time for which the Citrix ADC waits after the threshold is exceeded before it sends out a Notification to the Insight Center. Minimum value =  1 Maximum value =  65535
* `l7latencynotifyinterval` - (Optional) L7 Latency Notify Interval. This is the interval at which the Citrix ADC sends out notifications to the Insight Center after the wait time has passed. Minimum value =  1 Maximum value =  65535
* `l7latencymaxnotifycount` - (Optional) L7 Latency Max notify Count. This is the upper limit on the number of notifications sent to the Insight Center within an interval where the Latency is above the threshold. Minimum value =  1 Maximum value =  65535


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the icalatencyprofile. It has the same value as the `name` attribute.


## Import

A icalatencyprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_icalatencyprofile.tf_icalatencyprofile my_ica_latencyprofile
```
