---
subcategory: "Cloud"
---

# Data Source: cloudprofile

The cloudprofile data source allows you to retrieve information about a configured cloud profile on the Citrix ADC, including the virtual server and bound service group settings that it provisions.


## Example usage

```terraform
data "citrixadc_cloudprofile" "tf_cloudprofile" {
  name = "tf_cloudprofile"
}

output "cloudprofile_type" {
  value = data.citrixadc_cloudprofile.tf_cloudprofile.type
}

output "cloudprofile_ipaddress" {
  value = data.citrixadc_cloudprofile.tf_cloudprofile.ipaddress
}
```


## Argument Reference

* `name` - (Required) Name for the cloud profile to look up.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `type` - Type of cloud profile, either based on a virtual server (autoscale) or based on Azure tags.
* `vservername` - Name of the virtual server provisioned by the cloud profile.
* `servicetype` - Protocol used by the service (also called the service type).
* `ipaddress` - IPv4 or IPv6 address assigned to the virtual server.
* `port` - Port number for the virtual server.
* `servicegroupname` - Name of the service group bound to the virtual server.
* `boundservicegroupsvctype` - The protocol type of the bound service.
* `vsvrbindsvcport` - The port number used for the bound service.
* `graceful` - Indicates graceful shutdown of the service. The system waits for all outstanding connections to this service to be closed before disabling the service.
* `delay` - Time, in seconds, after which all the services configured on the server are disabled.
* `azuretagname` - Azure tag name used to select the cloud back ends when `type` is `azuretags`.
* `azuretagvalue` - Azure tag value used to select the cloud back ends when `type` is `azuretags`.
* `azurepollperiod` - Azure polling period, in seconds, at which the ADC queries Azure for tag-matched back ends.
* `id` - The id of the cloudprofile. It has the same value as the `name` attribute.
