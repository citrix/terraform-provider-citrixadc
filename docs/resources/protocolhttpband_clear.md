---
subcategory: "Protocol"
---

# Resource: protocolhttpband_clear

This resource is used to clear HTTP band statistics on the Citrix ADC.


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

* `id` - The ID of the protocolhttpband_clear resource. It has the format `protocolhttpband_clear-<type>` (for example, `protocolhttpband_clear-REQUEST`).
