---
subcategory: "LLDP"
---

# Resource: lldpparam

The lldpparam resource is used to update lldpparam.


## Example usage

```hcl
resource "citrixadc_lldpparam" "tf_lldpparam" {
  holdtimetxmult = 3
  mode           = "TRANSMITTER"
  timer          = 40
}
```


## Argument Reference

* `holdtimetxmult` - (Optional) A multiplier for calculating the duration for which the receiving device stores the LLDP information in its database before discarding or removing it. The duration is calculated as the holdtimeTxMult (Holdtime Multiplier) parameter value multiplied by the timer (Timer) parameter value. Minimum value =  1 Maximum value =  20
* `timer` - (Optional) Interval, in seconds, between LLDP packet data units (LLDPDUs).  that the Citrix ADC sends to a directly connected device. Minimum value =  1 Maximum value =  3000
* `mode` - (Optional) Global mode of Link Layer Discovery Protocol (LLDP) on the Citrix ADC. The resultant LLDP mode of an interface depends on the LLDP mode configured at the global and the interface levels. Possible values: [ NONE, TRANSMITTER, RECEIVER, TRANSCEIVER ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lldpparam. It is a unique string prefixed with `tf-lldpparam-` attribute.