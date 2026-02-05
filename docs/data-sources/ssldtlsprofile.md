---
subcategory: "SSL"
---

# Data Source: citrixadc_ssldtlsprofile

The ssldtlsprofile data source allows you to retrieve information about a DTLS profile.

## Example usage

```terraform
data "citrixadc_ssldtlsprofile" "tf_ssldtlsprofile" {
  name = "tf_ssldtlsprofile"
}

output "helloverifyrequest" {
  value = data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile.helloverifyrequest
}

output "maxbadmacignorecount" {
  value = data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile.maxbadmacignorecount
}

output "terminatesession" {
  value = data.citrixadc_ssldtlsprofile.tf_ssldtlsprofile.terminatesession
}
```

## Argument Reference

* `name` - (Required) Name for the DTLS profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `helloverifyrequest` - Send a Hello Verify request to validate the client.
* `id` - The id of the ssldtlsprofile. It has the same value as the `name` attribute.
* `initialretrytimeout` - Initial time out value to retransmit the last flight sent from the NetScaler.
* `maxbadmacignorecount` - Maximum number of bad MAC errors to ignore for a connection prior disconnect. Disabling parameter terminateSession terminates session immediately when bad MAC is detected in the connection.
* `maxholdqlen` - Maximum number of datagrams that can be queued at DTLS layer for processing.
* `maxpacketsize` - Maximum number of packets to reassemble. This value helps protect against a fragmented packet attack.
* `maxrecordsize` - Maximum size of records that can be sent if PMTU is disabled.
* `maxretrytime` - Wait for the specified time, in seconds, before resending the request.
* `pmtudiscovery` - Source for the maximum record size value. If ENABLED, the value is taken from the PMTU table. If DISABLED, the value is taken from the profile.
* `terminatesession` - Terminate the session if the message authentication code (MAC) of the client and server do not match.

## Import

A ssldtlsprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_ssldtlsprofile.tf_ssldtlsprofile tf_ssldtlsprofile
```
