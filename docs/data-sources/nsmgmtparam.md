---
subcategory: "NS"
---

# Data Source: nsmgmtparam

The nsmgmtparam data source allows you to retrieve the embedded management HTTP server (httpd) parameters configured on the Citrix ADC appliance, such as the management HTTP and HTTPS ports and the httpd worker capacity settings.


## Example usage

```terraform
data "citrixadc_nsmgmtparam" "example" {
}

output "mgmt_http_port" {
  value = data.citrixadc_nsmgmtparam.example.mgmthttpport
}
```


## Argument Reference

This datasource is a singleton and does not require any arguments. All attributes are computed.

## Attribute Reference

The following attributes are available:

* `id` - The id of the nsmgmtparam datasource. It is set to `nsmgmtparam-config`.
* `mgmthttpport` - Port on which the management HTTP server listens.
* `mgmthttpsport` - Port on which the management HTTPS server listens.
* `httpdmaxclients` - The HTTPD Max Clients value in the `httpd.conf` file. Minimum value: `1`, Maximum value: `255`.
* `httpdmaxreqworkers` - The HTTPD Max Request Workers value in the `httpd.conf` file. Minimum value: `1`, Maximum value: `255`.
