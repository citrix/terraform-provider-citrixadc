---
subcategory: "Cloud"
---

# Resource: cloudprofile

The cloudprofile resource configures a cloud profile on the Citrix ADC. A cloud profile provisions a load balancing virtual server together with its bound service group in a single step, either as an autoscale front end or driven by Azure resource tags, so that the ADC can automatically balance traffic across dynamically scaled cloud back ends.

~> **Note:** The cloudprofile resource is immutable. Every attribute is immutable, so changing any argument (including optional ones) destroys the existing cloud profile and creates a new one. There is no in-place update path; only create and delete operations are supported by the underlying NITRO API.


## Example usage

### Autoscale profile

```hcl
resource "citrixadc_cloudprofile" "tf_cloudprofile" {
  name                     = "tf_cloudprofile"
  type                     = "autoscale"
  vservername              = "tf_lbvserver"
  servicetype              = "HTTP"
  ipaddress                = "10.222.74.128"
  port                     = 80
  servicegroupname         = "tf_servicegroup"
  boundservicegroupsvctype = "HTTP"
  vsvrbindsvcport          = 80
  graceful                 = "NO"
  delay                    = 10
}
```

### Azure tags profile

```hcl
resource "citrixadc_cloudprofile" "tf_azure_cloudprofile" {
  name                     = "tf_azure_cloudprofile"
  type                     = "azuretags"
  vservername              = "tf_lbvserver"
  servicetype              = "SSL"
  ipaddress                = "10.222.74.129"
  port                     = 443
  servicegroupname         = "tf_servicegroup"
  boundservicegroupsvctype = "SSL"
  vsvrbindsvcport          = 443
  azuretagname             = "environment"
  azuretagvalue            = "production"
  azurepollperiod          = 120
}
```


## Argument Reference

All arguments below force a new resource to be created when changed (the resource is immutable).

* `name` - (Required) Name for the cloud profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created. Changing this attribute forces a new resource to be created.
* `type` - (Required) Type of cloud profile that you want to create, either based on a virtual server (autoscale) or based on Azure tags. Possible values: [ autoscale, azuretags ]. Changing this attribute forces a new resource to be created.
* `vservername` - (Required) Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Changing this attribute forces a new resource to be created.
* `servicetype` - (Required) Protocol used by the service (also called the service type). Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, MONGO, MONGO_TLS, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP, QUIC, IPFIX, LOGSTREAM, LOGSTREAM_SSL ]. Changing this attribute forces a new resource to be created.
* `ipaddress` - (Required) IPv4 or IPv6 address to assign to the virtual server. Changing this attribute forces a new resource to be created.
* `port` - (Required) Port number for the virtual server. Changing this attribute forces a new resource to be created.
* `servicegroupname` - (Required) Name of the service group to bind to this virtual server. Changing this attribute forces a new resource to be created.
* `boundservicegroupsvctype` - (Required) The protocol type of the bound service. Possible values: [ HTTP, FTP, TCP, UDP, SSL, SSL_BRIDGE, SSL_TCP, DTLS, NNTP, RPCSVR, DNS, ADNS, SNMP, RTSP, DHCPRA, ANY, SIP_UDP, SIP_TCP, SIP_SSL, DNS_TCP, ADNS_TCP, MYSQL, MSSQL, ORACLE, MONGO, MONGO_TLS, RADIUS, RADIUSListener, RDP, DIAMETER, SSL_DIAMETER, TFTP, SMPP, PPTP, GRE, SYSLOGTCP, SYSLOGUDP, FIX, SSL_FIX, USER_TCP, USER_SSL_TCP, QUIC, IPFIX, LOGSTREAM, LOGSTREAM_SSL ]. Changing this attribute forces a new resource to be created.
* `vsvrbindsvcport` - (Required) The port number to be used for the bound service. Changing this attribute forces a new resource to be created.
* `graceful` - (Optional) Indicates graceful shutdown of the service. The system waits for all outstanding connections to this service to be closed before disabling the service. Possible values: [ YES, NO ]. Defaults to `NO`. Changing this attribute forces a new resource to be created.
* `delay` - (Optional) Time, in seconds, after which all the services configured on the server are disabled. Changing this attribute forces a new resource to be created.
* `azuretagname` - (Optional) Azure tag name used to select the cloud back ends when `type` is `azuretags`. Maximum length = 511. Changing this attribute forces a new resource to be created.
* `azuretagvalue` - (Optional) Azure tag value used to select the cloud back ends when `type` is `azuretags`. Maximum length = 255. Changing this attribute forces a new resource to be created.
* `azurepollperiod` - (Optional) Azure polling period, in seconds, at which the ADC queries Azure for tag-matched back ends. Minimum value = 60. Maximum value = 3600. Defaults to `60`. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudprofile. It has the same value as the `name` attribute.


## Import

A cloudprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_cloudprofile.tf_cloudprofile tf_cloudprofile
```
