---
subcategory: "Subscriber"
---

# Data Source: subscriberprofile

The subscriberprofile data source allows you to retrieve information about a subscriber profile configured on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_subscriberprofile" "tf_subscriberprofile" {
  ip   = "10.222.74.185"
  vlan = 1
}

output "ip" {
  value = data.citrixadc_subscriberprofile.tf_subscriberprofile.ip
}

output "vlan" {
  value = data.citrixadc_subscriberprofile.tf_subscriberprofile.vlan
}

output "subscriptionidtype" {
  value = data.citrixadc_subscriberprofile.tf_subscriberprofile.subscriptionidtype
}
```

## Argument Reference

* `ip` - (Required) Subscriber ip address
* `vlan` - (Required) The vlan number on which the subscriber is located.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the subscriberprofile. It has the composite value of `ip` and `vlan` attributes.
* `servicepath` - Name of the servicepath to be taken for this subscriber.
* `subscriberrules` - Rules configured for this subscriber. This is similar to rules received from PCRF for dynamic subscriber sessions.
* `subscriptionidtype` - Subscription-Id type
* `subscriptionidvalue` - Subscription-Id value
