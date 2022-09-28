---
subcategory: "Subscriber"
---

# Resource: subscriberprofile

The subscriberprofile resource is used to create subscriberprofile.


## Example usage

```hcl
resource "citrixadc_subscriberprofile" "tf_subscriberprofile" {
  ip                  = "10.222.74.185"
  subscriptionidtype  = "E164"
  subscriptionidvalue = 5
}
```


## Argument Reference

* `ip` - (Required) Subscriber ip address
* `servicepath` - (Optional) Name of the servicepath to be taken for this subscriber.
* `subscriberrules` - (Optional) Rules configured for this subscriber. This is similar to rules received from PCRF for dynamic subscriber sessions.
* `subscriptionidtype` - (Optional) Subscription-Id type
* `subscriptionidvalue` - (Optional) Subscription-Id value
* `vlan` - (Optional) The vlan number on which the subscriber is located.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the subscriberprofile. It has the same value as the `ip` attribute.


## Import

A subscriberprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_subscriberprofile.tf_subscriberprofile 10.222.74.185
```
