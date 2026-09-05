---
subcategory: "Utility"
---

# Data Source: callhome

The callhome data source allows you to retrieve information about the Call Home configuration.


## Example usage

```hcl
data "citrixadc_callhome" "example" {}

output "callhome_details" {
  value = data.citrixadc_callhome.example
}
```

## Example usage with Resource

```hcl
data "citrixadc_callhome" "tf_callhome" {
  depends_on = [citrixadc_callhome.tf_callhome]
}

output "configured_callhome_mode" {
  value = data.citrixadc_callhome.tf_callhome.mode
}
```

## Argument Reference

This data source does not require any arguments. It retrieves the current Call Home configuration from the Citrix ADC.

## Attribute Reference

In addition to the above arguments, the following attributes are exported:

* `id` - The id of the callhome data source. It is set to `callhome`.
* `mode` - CallHome mode of operation. Possible values: `Default`, `CSP`, `Connector`.
* `emailaddress` - Email address of the contact administrator.
* `proxymode` - Enables or disables the proxy mode. The proxy server can be set by either specifying the IP address of the server or the name of the service representing the proxy server. Possible values: `YES`, `NO`.
* `ipaddress` - IP address of the proxy server.
* `proxyauthservice` - Name of the service that represents the proxy server.
* `port` - HTTP port on the Proxy server. This is a mandatory parameter for both IP address and service name based configuration.
* `hbcustominterval` - Interval (in days) between CallHome heartbeats.
* `nodeid` - Unique number that identifies the cluster node.

### Read-only callhome metadata

These attributes are returned by the appliance on a GET (they are not configurable on the `citrixadc_callhome` resource). They are GET-only / Computed, and any attribute the appliance does not return is `null`.

* `sslcardfirstfailure` - First occurrence SSL card failure.
* `sslcardlatestfailure` - Latest occurrence SSL card failure.
* `powfirstfail` - First occurrence power supply unit failure.
* `powlatestfailure` - Latest occurrence power supply unit failure.
* `hddfirstfail` - First occurrence hard disk drive failure.
* `hddlatestfailure` - Latest occurrence hard disk drive failure.
* `flashfirstfail` - First occurrence compact flash failure.
* `flashlatestfailure` - Latest occurrence compact flush failure.
* `rlfirsthighdrop` - First occurence of high rate limit drops.
* `rllatesthighdrop` - Latest occurence of high rate limit drops.
* `restartlatestfail` - Latest occurrence warm restart failure.
* `memthrefirstanomaly` - First occurrence of memory anomaly.
* `memthrelatestanomaly` - Latest occurrence of memory anomaly.
* `callhomestatus` - Callhome feature enabled/disable, register with upload server successful/failed. A list of strings.
* `anomalydetection` - Enables or disables anomaly detection.
