---
subcategory: "NS"
---

# Resource: nsappflowparam

Configures the global AppFlow parameters on the Citrix ADC. These settings control how IPFIX flow records are exported - the template refresh interval, the UDP packet MTU, which HTTP fields (URL, cookie, referer, method, host, user-agent) are included in the exported records, and whether records are generated only for client-side traffic. Use this resource to tune what application telemetry the appliance streams to an AppFlow collector.

This is a singleton resource: a single AppFlow parameter configuration always exists on the appliance, so this resource has no create or delete operation on the ADC - applying it updates the existing configuration, and destroying it only removes the resource from Terraform state.


## Example usage

```hcl
resource "citrixadc_nsappflowparam" "tf_nsappflowparam" {
  templaterefresh   = 600
  udppmtu           = 1472
  httpurl           = "ON"
  httpcookie        = "OFF"
  httpreferer       = "ON"
  httpmethod        = "ON"
  httphost          = "ON"
  httpuseragent     = "ON"
  clienttrafficonly = "NO"
}
```


## Argument Reference

* `templaterefresh` - (Optional) IPFIX template refresh interval, in seconds. Defaults to `600`.
* `udppmtu` - (Optional) MTU, in bytes, to be used for IPFIX UDP packets. Defaults to `1472`.
* `httpurl` - (Optional) Include the HTTP URL in the AppFlow records. Possible values: [ ON, OFF ]. Defaults to `OFF`.
* `httpcookie` - (Optional) Include the HTTP cookie in the AppFlow records. Possible values: [ ON, OFF ]. Defaults to `OFF`.
* `httpreferer` - (Optional) Include the HTTP referer in the AppFlow records. Possible values: [ ON, OFF ]. Defaults to `OFF`.
* `httpmethod` - (Optional) Include the HTTP method in the AppFlow records. Possible values: [ ON, OFF ]. Defaults to `OFF`.
* `httphost` - (Optional) Include the HTTP host in the AppFlow records. Possible values: [ ON, OFF ]. Defaults to `OFF`.
* `httpuseragent` - (Optional) Include the HTTP user-agent in the AppFlow records. Possible values: [ ON, OFF ]. Defaults to `OFF`.
* `clienttrafficonly` - (Optional) Generate AppFlow records only for client-side traffic. Possible values: [ YES, NO ].


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsappflowparam. Because this is a singleton, it is set to the constant string `nsappflowparam-config`.


## Import

A singleton resource is imported using the constant id `nsappflowparam-config`:

```shell
terraform import citrixadc_nsappflowparam.tf_nsappflowparam nsappflowparam-config
```
