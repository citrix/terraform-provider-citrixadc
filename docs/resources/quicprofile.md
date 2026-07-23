---
subcategory: "QUIC"
---

# Resource: quicprofile

This resource is used to manage QUIC profiles.


## Example usage

```hcl
resource "citrixadc_quicprofile" "tf_quicprofile" {
  name                           = "quic_profile1"
  activeconnectionmigration      = "ENABLED"
  congestionctrlalgorithm        = "CUBIC"
  statelessaddressvalidation     = "ENABLED"
  ackdelayexponent               = 3
  activeconnectionidlimit        = 3
  initialmaxdata                 = 1048576
  initialmaxstreamdatabidilocal  = 262144
  initialmaxstreamdatabidiremote = 262144
  initialmaxstreamdatauni        = 262144
  initialmaxstreamsbidi          = 100
  initialmaxstreamsuni           = 10
  maxackdelay                    = 20
  maxidletimeout                 = 180
  maxudpdatagramsperburst        = 8
  maxudppayloadsize              = 1472
  newtokenvalidityperiod         = 300
  retrytokenvalidityperiod       = 10
}
```


## Argument Reference

* `name` - (Required) Name for the QUIC profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals sign (=), and hyphen (-) characters. Maximum length = 255. Cannot be changed after the profile is created. Changing this attribute forces a new resource to be created.
* `ackdelayexponent` - (Optional) An integer value advertised by the Citrix ADC to the remote QUIC endpoint, indicating an exponent that the remote QUIC endpoint should use to decode the ACK Delay field in QUIC ACK frames sent by the Citrix ADC. Minimum value = 0. Maximum value = 20. Defaults to `3`.
* `activeconnectionidlimit` - (Optional) An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum number of QUIC connection IDs from the remote QUIC endpoint that the Citrix ADC is willing to store. Minimum value = 2. Maximum value = 8. Defaults to `3`.
* `activeconnectionmigration` - (Optional) Specify whether the Citrix ADC should allow the remote QUIC endpoint to perform active QUIC connection migration. Possible values: [ ENABLED, DISABLED ]. Defaults to `"ENABLED"`.
* `congestionctrlalgorithm` - (Optional) Congestion control algorithm to be used for QUIC connections. Possible values: [ Default, NewReno, CUBIC, BBR ]. Defaults to `"Default"`.
* `initialmaxdata` - (Optional) An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial value, in bytes, for the maximum amount of data that can be sent on a QUIC connection. Minimum value = 8192. Maximum value = 67108864. Defaults to `1048576`.
* `initialmaxstreamdatabidilocal` - (Optional) An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the Citrix ADC. Minimum value = 8192. Maximum value = 8388608. Defaults to `262144`.
* `initialmaxstreamdatabidiremote` - (Optional) An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for bidirectional QUIC streams initiated by the remote QUIC endpoint. Minimum value = 8192. Maximum value = 8388608. Defaults to `262144`.
* `initialmaxstreamdatauni` - (Optional) An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial flow control limit, in bytes, for unidirectional streams initiated by the remote QUIC endpoint. Minimum value = 8192. Maximum value = 8388608. Defaults to `262144`.
* `initialmaxstreamsbidi` - (Optional) An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial maximum number of bidirectional streams the remote QUIC endpoint may initiate. Minimum value = 1. Maximum value = 500. Defaults to `100`.
* `initialmaxstreamsuni` - (Optional) An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the initial maximum number of unidirectional streams the remote QUIC endpoint may initiate. Minimum value = 1. Maximum value = 500. Defaults to `10`.
* `maxackdelay` - (Optional) An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum amount of time, in milliseconds, by which the Citrix ADC will delay sending acknowledgments. Minimum value = 10. Maximum value = 2000. Defaults to `20`.
* `maxidletimeout` - (Optional) An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the maximum idle timeout, in seconds, for a QUIC connection. A QUIC connection will be silently discarded by the Citrix ADC if it remains idle for longer than the minimum of the idle timeout values advertised by the Citrix ADC and the remote QUIC endpoint, and three times the current Probe Timeout (PTO). Minimum value = 1. Maximum value = 3600. Defaults to `180`.
* `maxudpdatagramsperburst` - (Optional) An integer value specifying the maximum number of UDP datagrams that can be transmitted by the Citrix ADC in a single transmission burst on a QUIC connection. Minimum value = 1. Maximum value = 256. Defaults to `8`.
* `maxudppayloadsize` - (Optional) An integer value advertised by the Citrix ADC to the remote QUIC endpoint, specifying the size of the largest UDP datagram payload, in bytes, that the Citrix ADC is willing to receive on a QUIC connection. Minimum value = 1252. Maximum value = 9188. Defaults to `1472`.
* `newtokenvalidityperiod` - (Optional) An integer value specifying the validity period, in seconds, of address validation tokens issued through QUIC NEW_TOKEN frames sent by the Citrix ADC. Minimum value = 1. Maximum value = 3600. Defaults to `300`.
* `retrytokenvalidityperiod` - (Optional) An integer value specifying the validity period, in seconds, of address validation tokens issued through QUIC Retry packets sent by the Citrix ADC. Minimum value = 1. Maximum value = 120. Defaults to `10`.
* `statelessaddressvalidation` - (Optional) Specify whether the Citrix ADC should perform stateless address validation for QUIC clients, by sending tokens in QUIC Retry packets during QUIC connection establishment, and by sending tokens in QUIC NEW_TOKEN frames after QUIC connection establishment. Possible values: [ ENABLED, DISABLED ]. Defaults to `"ENABLED"`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the quicprofile. It has the same value as the `name` attribute.


## Import

A quicprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_quicprofile.tf_quicprofile quic_profile1
```
