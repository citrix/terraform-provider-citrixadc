---
subcategory: "NS"
---

# Resource: nsmgmtparam

This resource is used to manage the management HTTP server (httpd) parameters.


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

* `id` - The id of the nsmgmtparam. It is set to `nsmgmtparam-config`.


## Import

A singleton resource is imported using the constant id `nsmgmtparam-config`:

```shell
terraform import citrixadc_nsmgmtparam.tf_nsmgmtparam nsmgmtparam-config
```
