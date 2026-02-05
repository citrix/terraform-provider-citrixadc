---
subcategory: "Load Balancing"
---

# Data Source: citrixadc_lbsipparameters

The lbsipparameters data source allows you to retrieve information about the global SIP parameters configuration.


## Example Usage

### Standalone usage

```terraform
data "citrixadc_lbsipparameters" "tf_lbsipparameters" {
}

output "addrportvip" {
  value = data.citrixadc_lbsipparameters.tf_lbsipparameters.addrportvip
}

output "retrydur" {
  value = data.citrixadc_lbsipparameters.tf_lbsipparameters.retrydur
}

output "sip503ratethreshold" {
  value = data.citrixadc_lbsipparameters.tf_lbsipparameters.sip503ratethreshold
}
```

### Usage with resource dependency

```terraform
resource "citrixadc_lbsipparameters" "tf_lbsipparameters" {
  addrportvip         = "ENABLED"
  retrydur            = 100
  rnatdstport         = 80
  rnatsecuredstport   = 81
  rnatsecuresrcport   = 82
  rnatsrcport         = 83
  sip503ratethreshold = 15
}

data "citrixadc_lbsipparameters" "tf_lbsipparameters" {
  depends_on = [citrixadc_lbsipparameters.tf_lbsipparameters]
}
```


## Argument Reference

This data source does not require any arguments.


## Attribute Reference

The following attributes are exported:

* `id` - The id of the lbsipparameters resource. It is a unique string identifier.
* `addrportvip` - (String) Add the rport parameter to the VIA headers of SIP requests that virtual servers receive from clients or servers. Possible values: `ENABLED`, `DISABLED`.
* `retrydur` - (Number) Time, in seconds, for which a client must wait before initiating a connection after receiving a 503 Service Unavailable response from the SIP server. The time value is sent in the "Retry-After" header in the 503 response.
* `rnatdstport` - (Number) Port number with which to match the destination port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.
* `rnatsecuredstport` - (Number) Port number with which to match the destination port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching destination port number, and CALL-ID based persistence is implemented for the responses received by the virtual server. Range: 1-65535.
* `rnatsecuresrcport` - (Number) Port number with which to match the source port in server-initiated SIP over SSL traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server. Range: 1-65535.
* `rnatsrcport` - (Number) Port number with which to match the source port in server-initiated SIP traffic. The rport parameter is added, without a value, to SIP packets that have a matching source port number, and CALL-ID based persistence is implemented for the responses received by the virtual server.
* `sip503ratethreshold` - (Number) Maximum number of 503 Service Unavailable responses to generate, once every 10 milliseconds, when a SIP virtual server becomes unavailable.
