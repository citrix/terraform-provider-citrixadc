---
subcategory: "NS"
---

# Resource: nslicenseserver

The nslicenseserver resource is used to create license server entry in ADC.


## Example usage

```hcl
resource "citrixadc_nslicenseserver" "tf_licenseserver" {
  servername = "10.78.60.200"
  port       = 27000
}
```


## Argument Reference

* `servername` - (Required) Fully qualified domain name of the License server.
* `port` - (Optional) License server port.
* `forceupdateip` - (Optional) If this flag is used while adding the licenseserver, existing config will be overwritten. Use this flag only if you are sure that the new licenseserver has the required capacity.
* `nodeid` - (Optional) Unique number that identifies the cluster node.
* `deviceprofilename` - (Optional) Device profile is created on ADM and contains the user name and password of the instance(s). ADM will use this info to add the NS for registration
* `licensemode` - (Optional) This paramter indicates type of license customer interested while configuring add/set licenseserver
* `licenseserverip` - (Optional) IP address of the License server.
* `password` - (Optional) Password to use when authenticating with ADM Agent for LAS licensing.
* `username` - (Optional) Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nslicenseserver. It has the same value as the `servername` attribute.
