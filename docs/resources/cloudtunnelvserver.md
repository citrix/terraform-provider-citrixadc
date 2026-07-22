---
subcategory: "Cloud"
---

# Resource: cloudtunnelvserver

The cloudtunnelvserver resource configures a Cloud Tunnel virtual server on the Citrix ADC. A Cloud Tunnel virtual server acts as the listener that accepts client traffic and tunnels it through a cloud tunnel connection, letting you extend on-premises services into a cloud back end without exposing them directly. Use it to define which traffic (by service type and an optional listen policy) is forwarded over the cloud tunnel.

~> **Prerequisite:** The Cloud Tunnel feature is license- and feature-gated. Ensure the cloud tunnel feature is licensed and enabled on the target Citrix ADC before creating this resource; otherwise the NITRO API call will fail.


## Example usage

```hcl
resource "citrixadc_cloudtunnelvserver" "tf_cloudtunnelvserver" {
  name           = "ct_vserver1"
  servicetype    = "TCP"
  listenpolicy   = "CLIENT.IP.SRC.IN_SUBNET(10.0.0.0/8)"
  listenpriority = 50
}
```

### Minimal example (defaults applied)

```hcl
resource "citrixadc_cloudtunnelvserver" "tf_cloudtunnelvserver_min" {
  name        = "ct_vserver2"
  servicetype = "UDP"
}
```


## Argument Reference

* `name` - (Required) Name for the Cloud Tunnel virtual server. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Maximum length = 127. Changing this attribute forces a new resource to be created.
* `servicetype` - (Required) Service type of the listener using which traffic will be tunneled through the cloud tunnel server. Possible values: [ TCP, UDP ]. Changing this attribute forces a new resource to be created.
* `listenpolicy` - (Optional) String specifying the listen policy for the Cloud Tunnel virtual server. Can be either a named expression or an expression. The Cloud Tunnel virtual server processes only the traffic for which the expression evaluates to true. Maximum length = 1499. Defaults to `"none"`.
* `listenpriority` - (Optional) Integer specifying the priority of the listen policy. A higher number specifies a lower priority. If a request matches the listen policies of more than one virtual server, the virtual server whose listen policy has the highest priority (the lowest priority number) accepts the request. Minimum value = 0. Maximum value = 100. Defaults to `101`.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the cloudtunnelvserver. It has the same value as the `name` attribute.


## Import

A cloudtunnelvserver can be imported using its name, e.g.

```shell
terraform import citrixadc_cloudtunnelvserver.tf_cloudtunnelvserver ct_vserver1
```
