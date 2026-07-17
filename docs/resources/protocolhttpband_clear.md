---
subcategory: "Protocol"
---

# Resource: protocolhttpband_clear

The protocolhttpband_clear resource clears (resets) the HTTP band statistics that the Citrix ADC accumulates for request and response payload sizes. Use it when you want to discard the previously collected size-distribution counters for a given band type so that the HTTP band statistics reports start accumulating from zero again (for example, after tuning `reqbandsize` / `respbandsize` on the `protocolhttpband` object, or before capturing a fresh measurement window).

~> **One-shot action.** This resource maps to the NITRO `clear` action (`POST ?action=clear`, CLI: `clear protocol httpBand -type <type>`); it does not create a persistent object on the appliance. Each `terraform apply` that creates or replaces this resource performs the clear once. There is no readable server-side object and no NITRO GET endpoint, so there is no corresponding data source: Read is a no-op, Delete only removes the resource from Terraform state, and changing `type` forces a new clear (replacement).


## Example usage

```hcl
resource "citrixadc_protocolhttpband_clear" "tf_protocolhttpband_clear" {
  type = "REQUEST"
}
```


## Argument Reference

* `type` - (Required) Type of HTTP band statistics to clear. Possible values: [ REQUEST, RESPONSE, MQTT_JUMBO_REQ ]. Changing this value forces the resource to be recreated (re-running the clear action against the new band type).


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The ID of the protocolhttpband_clear resource. It is a synthetic identifier with the format `protocolhttpband_clear-<type>` (for example, `protocolhttpband_clear-REQUEST`); it does not correspond to any object on the Citrix ADC.
