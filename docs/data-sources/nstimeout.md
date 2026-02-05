---
subcategory: "Network"
---

# Data Source `nstimeout`

The nstimeout data source allows you to retrieve information about global timeout settings configured on the Citrix ADC.


## Example usage

```terraform
data "citrixadc_nstimeout" "tf_nstimeout" {
}

output "zombie" {
  value = data.citrixadc_nstimeout.tf_nstimeout.zombie
}

output "client" {
  value = data.citrixadc_nstimeout.tf_nstimeout.client
}
```


## Argument Reference

This data source has no required arguments.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `anyclient` - Global idle timeout, in seconds, for non-TCP client connections.
* `anyserver` - Global idle timeout, in seconds, for non TCP server connections.
* `anytcpclient` - Global idle timeout, in seconds, for TCP client connections.
* `anytcpserver` - Global idle timeout, in seconds, for TCP server connections.
* `client` - Client idle timeout, in seconds.
* `halfclose` - Idle timeout, in seconds, for connections in the TCP half-closed state.
* `httpclient` - HTTP client idle timeout, in seconds.
* `httpserver` - HTTP server idle timeout, in seconds.
* `nontcpzombie` - Interval at which the zombie clean-up process for non-TCP connections should run.
* `reducedfintimeout` - Alternative idle timeout, in seconds, for closed TCP NATPCB connections.
* `reducedrsttimeout` - Timer interval, in seconds, for abruptly terminated TCP NATPCB connections.
* `server` - Server idle timeout, in seconds.
* `zombie` - Interval at which the zombie clean-up process for TCP connections should run.
