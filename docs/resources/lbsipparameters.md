---
subcategory: "Load Balancing"
---

# Resource: lbsipparameters

The lbsipparameters resource is used to configure the global SIP parameters.


## Example usage

```hcl
resource "citrixadc_lbsipparameters" "tf_lbsipparameters" {
	addrportvip = "ENABLED"
	retrydur = 100
	rnatdstport = 80
	rnatsecuredstport = 81
	rnatsecuresrcport = 82
	rnatsrcport = 83
	sip503ratethreshold = 15
}
```


## Argument Reference

* `rnatsrcport` - (Optional) Port number with which to match the source port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.
* `rnatdstport` - (Optional) Port number with which to match the destination port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.
* `retrydur` - (Optional) Time, in seconds, for which a client must wait before initiating a connection after receiving a 503 Service Unavailable response from the SIP server. The time value is sent in the "Retry-After" header in the 503 response.
* `addrportvip` - (Optional) Add the rport parameter to the VIA headers of SIP requests that virtual servers receive from clients or servers. Possible values: [ ENABLED, DISABLED ]
* `sip503ratethreshold` - (Optional) 
* `rnatsecuresrcport` - (Optional) Port number with which to match the source port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server. Range 1 - 65535 * in CLI is represented as 65535 in NITRO API
* `rnatsecuredstport` - (Optional) Port number with which to match the destination port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server. Range 1 - 65535 * in CLI is represented as 65535 in NITRO API


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbsipparameters. It is a unique string prefixed with "tf_lbsipparameters"
