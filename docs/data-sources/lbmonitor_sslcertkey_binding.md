---
subcategory: "Load Balancing"
---

# Data Source: lbmonitor_sslcertkey_binding

The lbmonitor_sslcertkey_binding data source allows you to retrieve information about the SSL certificate key bound to a load balancing monitor.


## Example usage

```terraform
data "citrixadc_lbmonitor_sslcertkey_binding" "tf_lbmonitor_sslcertkey_binding" {
  monitorname = "tf_monitor"
  certkeyname = "tf_sslcertkey"
  ca          = false
}

output "crlcheck" {
  value = data.citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding.crlcheck
}

output "ocspcheck" {
  value = data.citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding.ocspcheck
}
```


## Argument Reference

* `monitorname` - (Required) Name of the monitor.
* `certkeyname` - (Required) The name of the certificate bound to the monitor.
* `ca` - (Required) Boolean flag indicating that the bound certificate is a CA certificate, controlling the rule for use of the CRL corresponding to this CA certificate during client authentication. If crlCheck is set to Mandatory, the system will deny all SSL clients if the CRL is missing, expired (NextUpdate date is in the past), or is incomplete with remote CRL refresh enabled. If crlCheck is set to optional, the system will allow SSL clients in the above error cases. However, in any case if the client certificate is revoked in the CRL, the SSL client will be denied access.
* `crlcheck` - (Optional) The state of the CRL check parameter. Possible values: [ Mandatory, Optional ]
* `ocspcheck` - (Optional) The state of the OCSP check parameter. Possible values: [ Mandatory, Optional ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbmonitor_sslcertkey_binding. It is the concatenation of the `monitorname` and `certkeyname` attributes separated by a comma.
* `crlcheck` - The state of the CRL check parameter.
* `ocspcheck` - The state of the OCSP check parameter.
