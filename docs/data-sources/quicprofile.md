---
subcategory: "QUIC"
---

# Data Source: quicprofile

The quicprofile data source allows you to retrieve information about an existing QUIC profile configured on the Citrix ADC, looked up by its name.


## Example usage

```terraform
data "citrixadc_quicprofile" "example" {
  name = "quic_profile1"
}

output "congestionctrlalgorithm" {
  value = data.citrixadc_quicprofile.example.congestionctrlalgorithm
}
```


## Argument Reference

* `name` - (Required) Name of the QUIC profile to look up.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the quicprofile. It has the same value as the `name` attribute.
* `ackdelayexponent` - Exponent that the remote QUIC endpoint should use to decode the ACK Delay field in QUIC ACK frames sent by the Citrix ADC.
* `activeconnectionidlimit` - Maximum number of QUIC connection IDs from the remote QUIC endpoint that the Citrix ADC is willing to store.
* `activeconnectionmigration` - Whether the Citrix ADC allows the remote QUIC endpoint to perform active QUIC connection migration.
* `congestionctrlalgorithm` - Congestion control algorithm used for QUIC connections.
* `initialmaxdata` - Initial value, in bytes, for the maximum amount of data that can be sent on a QUIC connection.
* `initialmaxstreamdatabidilocal` - Initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the Citrix ADC.
* `initialmaxstreamdatabidiremote` - Initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the remote QUIC endpoint.
* `initialmaxstreamdatauni` - Initial flow control limit, in bytes, for unidirectional streams initiated by the remote QUIC endpoint.
* `initialmaxstreamsbidi` - Initial maximum number of bidirectional streams the remote QUIC endpoint may initiate.
* `initialmaxstreamsuni` - Initial maximum number of unidirectional streams the remote QUIC endpoint may initiate.
* `maxackdelay` - Maximum amount of time, in milliseconds, by which the Citrix ADC will delay sending acknowledgments.
* `maxidletimeout` - Maximum idle timeout, in seconds, for a QUIC connection.
* `maxudpdatagramsperburst` - Maximum number of UDP datagrams that can be transmitted by the Citrix ADC in a single transmission burst on a QUIC connection.
* `maxudppayloadsize` - Size of the largest UDP datagram payload, in bytes, that the Citrix ADC is willing to receive on a QUIC connection.
* `newtokenvalidityperiod` - Validity period, in seconds, of address validation tokens issued through QUIC NEW_TOKEN frames sent by the Citrix ADC.
* `retrytokenvalidityperiod` - Validity period, in seconds, of address validation tokens issued through QUIC Retry packets sent by the Citrix ADC.
* `statelessaddressvalidation` - Whether the Citrix ADC performs stateless address validation for QUIC clients.
