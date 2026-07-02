---
subcategory: "Utility"
---

# Data Source: callhome

The `citrixadc_callhome` data source is used to retrieve the Call Home configuration from the Citrix ADC. Because Call Home is a singleton feature, there is exactly one configuration per appliance, so no lookup argument is required.

## Example Usage

```hcl
data "citrixadc_callhome" "example" {}

output "callhome_details" {
  value = data.citrixadc_callhome.example
}
```

## Example Usage with Resource

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

* `id` - The id of the callhome data source. Because callhome is a singleton, this is always the static string `"callhome"`.
* `mode` - CallHome mode of operation. Possible values: `Default`, `CSP`, `Connector`.
* `emailaddress` - Email address of the contact administrator.
* `proxymode` - Enables or disables the proxy mode. The proxy server can be set by either specifying the IP address of the server or the name of the service representing the proxy server. Possible values: `YES`, `NO`.
* `ipaddress` - IP address of the proxy server.
* `proxyauthservice` - Name of the service that represents the proxy server.
* `port` - HTTP port on the Proxy server. This is a mandatory parameter for both IP address and service name based configuration.
* `hbcustominterval` - Interval (in days) between CallHome heartbeats.
* `nodeid` - Unique number that identifies the cluster node.
