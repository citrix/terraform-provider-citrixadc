---
subcategory: "SMPP"
---

# Data Source: citrixadc_smppparam

The smppparam data source allows you to retrieve information about SMPP parameters configuration.

## Example Usage

```terraform
data "citrixadc_smppparam" "tf_smppparam" {
}

output "clientmode" {
  value = data.citrixadc_smppparam.tf_smppparam.clientmode
}

output "msgqueue" {
  value = data.citrixadc_smppparam.tf_smppparam.msgqueue
}
```

## Argument Reference

This datasource does not require any arguments.

## Attribute Reference

The following attributes are available:

* `id` - The id of the smppparam. It is a system-generated identifier.
* `clientmode` - Mode in which the client binds to the ADC. Modes are: `TRANSMITTERONLY`, `RECEIVERONLY`, `TRANSCEIVER`.
* `msgqueue` - Queue messages if the SMPP server is unavailable. Possible values: `ON`, `OFF`.
* `msgqueuesize` - Maximum number of messages that can be queued.
* `addrton` - Type of Number, such as an international number or a national number, used in the ESME address sent in the bind request.
* `addrnpi` - Numbering Plan Indicator, such as landline, data, or WAP client, used in the ESME address sent in the bind request.
* `addrrange` - Set of SME addresses, sent in the bind request, serviced by the ESME.
