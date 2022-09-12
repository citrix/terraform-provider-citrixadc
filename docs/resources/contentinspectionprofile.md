---
subcategory: "CI"
---

# Resource: contentinspectionprofile

The contentinspectionprofile resource is used to create contentinspectionprofile.


## Example usage

```hcl
resource "citrixadc_contentinspectionprofile" "tf_contentinspectionprofile" {
  name             = "my_ci_profile"
  type             = "INLINEINSPECTION"
  ingressinterface = "LA/2"
  egressinterface  = "LA/3"
}
```


## Argument Reference

* `name` - (Required) Name of a ContentInspection profile. Must begin with a letter, number, or the underscore \(_\) character. Other characters allowed, after the first character, are the hyphen \(-\), period \(.\), hash \(\#\), space \( \), at \(@\), colon \(:\), and equal \(=\) characters. The name of a IPS profile cannot be changed after it is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my ips profile" or 'my ips profile'\).
* `type` - (Required) Type of ContentInspection profile. Following types are available to configure:            INLINEINSPECTION : To inspect the packets/requests using IPS. 	   MIRROR : To forward cloned packets.
* `egressinterface` - (Optional) Egress interface for CI profile.It is a mandatory argument while creating an ContentInspection profile of type INLINEINSPECTION or MIRROR.
* `egressvlan` - (Optional) Egress Vlan for CI
* `ingressinterface` - (Optional) Ingress interface for CI profile.It is a mandatory argument while creating an ContentInspection profile of IPS type.
* `ingressvlan` - (Optional) Ingress Vlan for CI
* `iptunnel` - (Optional) IP Tunnel for CI profile. It is used while creating a ContentInspection profile of type MIRROR when the IDS device is in a different network


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the contentinspectionprofile. It has the same value as the `name` attribute.


## Import

A contentinspectionprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_contentinspectionprofile.tf_contentinspectionprofile my_ci_profile
```
