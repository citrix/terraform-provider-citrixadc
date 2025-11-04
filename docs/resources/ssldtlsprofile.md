---
subcategory: "SSL"
---

# Resource: ssldtlsprofile

The ssldtlsprofile resource is used to create a new DTLS profile.


## Example usage

```hcl
resource "citrixadc_ssldtlsprofile" "tf_ssldtlsprofile" {
	name = "tf_ssldtlsprofile"
	helloverifyrequest = "ENABLED"
	maxbadmacignorecount = 128
	maxholdqlen = 64
	maxpacketsize = 125
	maxrecordsize = 250
	maxretrytime = 5
	pmtudiscovery = "DISABLED"
	terminatesession = "ENABLED"
}
```


## Argument Reference

* `name` - (Required) Name for the DTLS profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.
* `pmtudiscovery` - (Optional) Source for the maximum record size value. If ENABLED, the value is taken from the PMTU table. If DISABLED, the value is taken from the profile. Possible values: [ ENABLED, DISABLED ]
* `maxrecordsize` - (Optional) 
* `maxretrytime` - (Optional) Wait for the specified time, in seconds, before resending the request.
* `helloverifyrequest` - (Optional) Send a Hello Verify request to validate the client. Possible values: [ ENABLED, DISABLED ]
* `terminatesession` - (Optional) Terminate the session if the message authentication code (MAC) of the client and server do not match. Possible values: [ ENABLED, DISABLED ]
* `maxpacketsize` - (Optional) 
* `initialretrytimeout` - (Optional) Initial time out value to retransmit the last flight sent from the NetScaler.
* `maxbadmacignorecount` - (Optional) Maximum number of bad MAC errors to ignore for a connection prior disconnect. Disabling parameter terminateSession terminates session immediately when bad MAC is detected in the connection.
* `maxholdqlen` - (Optional) Maximum number of datagrams that can be queued at DTLS layer for processing


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the ssldtlsprofile. It has the same value as the `name` attribute.


## Import

A ssldtlsprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_ssldtlsprofile.tf_ssldtlsprofile tf_ssldtlsprofile
```
