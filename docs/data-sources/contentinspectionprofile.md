---
subcategory: "CI"
---

# Data Source `contentinspectionprofile`

The contentinspectionprofile data source allows you to retrieve information about content inspection profiles.


## Example usage

```terraform
data "citrixadc_contentinspectionprofile" "tf_contentinspectionprofile" {
  name = "my_ci_profile"
}

output "type" {
  value = data.citrixadc_contentinspectionprofile.tf_contentinspectionprofile.type
}

output "ingressinterface" {
  value = data.citrixadc_contentinspectionprofile.tf_contentinspectionprofile.ingressinterface
}

output "egressinterface" {
  value = data.citrixadc_contentinspectionprofile.tf_contentinspectionprofile.egressinterface
}
```


## Argument Reference

* `name` - (Required) Name of a ContentInspection profile. Must begin with a letter, number, or the underscore \(_\) character. Other characters allowed, after the first character, are the hyphen \(-\), period \(.\), hash \(\#\), space \( \), at \(@\), colon \(:\), and equal \(=\) characters. The name of a IPS profile cannot be changed after it is created.  CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my ips profile" or 'my ips profile'\).

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the contentinspectionprofile. It has the same value as the `name` attribute.
* `type` - Type of ContentInspection profile. Following types are available to configure:            INLINEINSPECTION : To inspect the packets/requests using IPS. 	   MIRROR : To forward cloned packets.
* `egressinterface` - Egress interface for CI profile.It is a mandatory argument while creating an ContentInspection profile of type INLINEINSPECTION or MIRROR.
* `egressvlan` - Egress Vlan for CI
* `ingressinterface` - Ingress interface for CI profile.It is a mandatory argument while creating an ContentInspection profile of IPS type.
* `ingressvlan` - Ingress Vlan for CI
* `iptunnel` - IP Tunnel for CI profile. It is used while creating a ContentInspection profile of type MIRROR when the IDS device is in a different network
