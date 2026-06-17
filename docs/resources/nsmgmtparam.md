---
subcategory: "NS"
---

# Resource: nsmgmtparam

Configures the embedded management HTTP server (httpd) parameters on the Citrix ADC. Use this resource to change the ports the management GUI and API listen on (HTTP and HTTPS), and to tune the httpd worker capacity by setting either the maximum number of clients or the maximum number of request workers in `httpd.conf`.

This is a singleton resource: a single management parameter configuration always exists on the appliance, so this resource has no create or delete operation on the ADC - applying it updates the existing configuration, and destroying it only removes the resource from Terraform state.


## Example usage

```hcl
resource "citrixadc_nsmgmtparam" "tf_nsmgmtparam" {
  mgmthttpport    = 80
  mgmthttpsport   = 443
  httpdmaxclients = 30
}
```


## Argument Reference

* `mgmthttpport` - (Optional) Port on which the management HTTP server listens. Defaults to `80`.
* `mgmthttpsport` - (Optional) Port on which the management HTTPS server listens. Defaults to `443`.
* `httpdmaxclients` - (Optional) Sets the HTTPD Max Clients value in the `httpd.conf` file. You can configure either Max Clients or Max Request Workers. Minimum value = `1`. Maximum value = `255`.
* `httpdmaxreqworkers` - (Optional) Sets the HTTPD Max Request Workers value in the `httpd.conf` file. You can configure either Max Clients or Max Request Workers. Minimum value = `1`. Maximum value = `255`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsmgmtparam. Because this is a singleton, it is set to the constant string `nsmgmtparam-config`.


## Import

A singleton resource is imported using the constant id `nsmgmtparam-config`:

```shell
terraform import citrixadc_nsmgmtparam.tf_nsmgmtparam nsmgmtparam-config
```
