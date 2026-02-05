---
subcategory: "NS"
---

# citrixadc_nslicenseserver (Data Source)

Data source for querying Citrix ADC license server configuration. This data source retrieves information about the configured license server that the ADC uses for licensing.

## Example Usage

```hcl
data "citrixadc_nslicenseserver" "example" {
  servername = "licenseserver.example.com"
}

# Output license server information
output "license_server_port" {
  value = data.citrixadc_nslicenseserver.example.port
}

output "license_mode" {
  value = data.citrixadc_nslicenseserver.example.licensemode
}
```

## Argument Reference

The following arguments are supported:

* `servername` - (Optional) Fully qualified domain name of the License server. Either `licenseserverip` or `servername` must be specified to identify the resource.
* `licenseserverip` - (Optional) IP address of the License server. Either `licenseserverip` or `servername` must be specified to identify the resource.

**Note:** At least one of `servername` or `licenseserverip` must be specified.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - The ID of the nslicenseserver datasource.
* `port` - License server port.
* `servername` - Fully qualified domain name of the License server.
* `licenseserverip` - IP address of the License server.
* `licensemode` - Type of license configured for the license server (e.g., POOLED, CICO).
* `nodeid` - Unique number that identifies the cluster node.
* `deviceprofilename` - Device profile created on ADM that contains the user name and password of the instance(s). ADM will use this info to add the NS for registration.
* `username` - Username to authenticate with ADM Agent for LAS licensing.
* `password` - Password to use when authenticating with ADM Agent for LAS licensing.
* `forceupdateip` - Flag indicating if existing config will be overwritten while adding the licenseserver.
