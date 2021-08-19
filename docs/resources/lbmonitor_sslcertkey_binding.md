---
subcategory: "Load Balancing"
---

# Resource: lbmonitor_sslcertkey_binding

The lbmonitor_sslcertkey_bindingresource is used to add an sslcertkey to lbmonitor.


## Example usage

```hcl
resource "citrixadc_lbmonitor_sslcertkey_binding" "tf_lbmonitor_sslcertkey_binding" {
	monitorname = citrixadc_lbmonitor.tf_monitor.monitorname
	certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
}

resource "citrixadc_lbmonitor" "tf_monitor" {
	monitorname = "tf_monitor"
	type = "HTTP"
	sslprofile = "ns_default_ssl_profile_backend"
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	certkey = "tf_sslcertkey"
	cert = "/var/tmp/certificate1.crt"
	key = "/var/tmp/key1.pem"
}
```


## Argument Reference

* `certkeyname` - (Required) The name of the certificate bound to the monitor.
* `ca` - (Optional) The rule for use of CRL corresponding to this CA certificate during client authentication. If crlCheck is set to Mandatory, the system will deny all SSL clients if the CRL is missing, expired - NextUpdate date is in the past, or is incomplete with remote CRL refresh enabled. If crlCheck is set to optional, the system will allow SSL clients in the above error cases.However, in any case if the client certificate is revoked in the CRL, the SSL client will be denied access.
* `crlcheck` - (Optional) The state of the CRL check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]
* `ocspcheck` - (Optional) The state of the OCSP check parameter. (Mandatory/Optional). Possible values: [ Mandatory, Optional ]
* `monitorname` - (Required) Name of the monitor.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lbmonitor_sslcertkey_binding. It has the same value as the concatenation of the `monitorname` and `certkeyname` attributes separated by a comma.


## Import

A lbmonitor\_sslcertkey\_binding can be imported using its id, e.g.

```shell
terraform import citrixadc_lbmonitor_sslcertkey_binding.tf_lbmonitor_sslcertkey_binding tf_monitor,tf_sslcertkey
```

