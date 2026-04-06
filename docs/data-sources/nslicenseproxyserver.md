---
subcategory: "NS"
---

# citrixadc_nslicenseproxyserver (Data Source)

Data source for querying Citrix ADC license proxy server configuration. This data source retrieves information about the configured license proxy server that the ADC uses to communicate with the license server.

## Example Usage

```hcl
data "citrixadc_nslicenseproxyserver" "example" {
  servername = "proxy.example.com"
}

# Output license proxy server information
output "proxy_port" {
  value = data.citrixadc_nslicenseproxyserver.example.port
}
```

## Argument Reference

The following arguments are supported:

* `servername` - (Optional) Fully qualified domain name of the License proxy server. Either `serverip` or `servername` must be specified to identify the resource.
* `serverip` - (Optional) IP address of the License proxy server. Either `serverip` or `servername` must be specified to identify the resource.

**Note:** At least one of `servername` or `serverip` must be specified.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the nslicenseproxyserver datasource.
* `port` - License proxy server port.
* `servername` - Fully qualified domain name of the License proxy server.
* `serverip` - IP address of the License proxy server.
