---
subcategory: "NS"
---

# Resource: nstimeout

The nstimeout resource is used to create nstimeout.


## Example usage

```hcl
resource "citrixadc_nstimeout" "tf_nstimeout" {
  zombie     = 70
  client     = 2300
  server     = 2400
  httpclient = 2500
  reducedrsttimeout = 15
}
```


## Argument Reference

* `zombie` - (Optional) Interval, in seconds, at which the Citrix ADC zombie cleanup process must run. This process cleans up inactive TCP connections. Minimum value =  1 Maximum value =  600
* `client` - (Optional) Client idle timeout (in seconds). If zero, the service-type default value is taken when service is created. Minimum value =  0 Maximum value =  18000
* `server` - (Optional) Server idle timeout (in seconds).  If zero, the service-type default value is taken when service is created. Minimum value =  0 Maximum value =  18000
* `httpclient` - (Optional) Global idle timeout, in seconds, for client connections of HTTP service type. This value is over ridden by the client timeout that is configured on individual entities. Minimum value =  0 Maximum value =  18000
* `httpserver` - (Optional) Global idle timeout, in seconds, for server connections of HTTP service type. This value is over ridden by the server timeout that is configured on individual entities. Minimum value =  0 Maximum value =  18000
* `tcpclient` - (Optional) Global idle timeout, in seconds, for non-HTTP client connections of TCP service type. This value is over ridden by the client timeout that is configured on individual entities. Minimum value =  0 Maximum value =  18000
* `tcpserver` - (Optional) Global idle timeout, in seconds, for non-HTTP server connections of TCP service type. This value is over ridden by the server timeout that is configured on entities. Minimum value =  0 Maximum value =  18000
* `anyclient` - (Optional) Global idle timeout, in seconds, for non-TCP client connections. This value is over ridden by the client timeout that is configured on individual entities. Minimum value =  0 Maximum value =  31536000
* `anyserver` - (Optional) Global idle timeout, in seconds, for non TCP server connections. This value is over ridden by the server timeout that is configured on individual entities. Minimum value =  0 Maximum value =  31536000
* `anytcpclient` - (Optional) Global idle timeout, in seconds, for TCP client connections. This value takes precedence over  entity level timeout settings (vserver/service). This is applicable only to transport protocol TCP. Minimum value =  0 Maximum value =  31536000
* `anytcpserver` - (Optional) Global idle timeout, in seconds, for TCP server connections. This value takes precedence over entity level timeout settings ( vserver/service). This is applicable only to transport protocol TCP. Minimum value =  0 Maximum value =  31536000
* `halfclose` - (Optional) Idle timeout, in seconds, for connections that are in TCP half-closed state. Minimum value =  1 Maximum value =  600
* `nontcpzombie` - (Optional) Interval at which the zombie clean-up process for non-TCP connections should run. Inactive IP NAT connections will be cleaned up. Minimum value =  1 Maximum value =  600
* `reducedfintimeout` - (Optional) Alternative idle timeout, in seconds, for closed TCP NATPCB connections. Minimum value =  1 Maximum value =  300
* `reducedrsttimeout` - (Optional) Timer interval, in seconds, for abruptly terminated TCP NATPCB connections. Minimum value =  0 Maximum value =  300
* `newconnidletimeout` - (Optional) Timer interval, in seconds, for new TCP NATPCB connections on which no data was received. Minimum value =  1 Maximum value =  120


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nstimeout. It is a unique string prefixed with "tf-nstimeout-
