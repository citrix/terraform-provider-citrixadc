---
subcategory: "SMPP"
---

# Resource: smppparam

The smppparam resource is used to update smppparam.


## Example usage

```hcl
resource "citrixadc_smppparam" "tf_smppparam" {
  clientmode = "TRANSCEIVER"
  msgqueue   = "OFF"
  addrnpi    = 40
  addrton    = 40
}
```


## Argument Reference

* `addrnpi` - (Optional) Numbering Plan Indicator, such as landline, data, or WAP client, used in the ESME address sent in the bind request.
* `addrrange` - (Optional) Set of SME addresses, sent in the bind request, serviced by the ESME.
* `addrton` - (Optional) Type of Number, such as an international number or a national number, used in the ESME address sent in the bind request.
* `clientmode` - (Optional) Mode in which the client binds to the ADC. Applicable settings function as follows: * TRANSCEIVER - Client can send and receive messages to and from the message center. * TRANSMITTERONLY - Client can only send messages. * RECEIVERONLY - Client can only receive messages.
* `msgqueue` - (Optional) Queue SMPP messages if a client that is capable of receiving the destination address messages is not available.
* `msgqueuesize` - (Optional) Maximum number of SMPP messages that can be queued. After the limit is reached, the Citrix ADC sends a deliver_sm_resp PDU, with an appropriate error message, to the message center.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the smppparam. It is a unique string prefixed with `tf-smppparam-`.
