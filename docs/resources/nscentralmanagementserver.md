---
subcategory: "NS"
---

# Resource: nscentralmanagementserver

The nscentralmanagementserver resource registers the Citrix ADC with a central management server (Citrix Application Delivery Management, ADM) so the instance can be managed centrally. Use it to onboard the ADC to either an on-premises ADM deployment or the ADM cloud service, supplying the management server address and the credentials ADM uses to create the device profile for the instance.

All attributes are immutable: changing any of them forces a new resource to be created (the registration is deleted and re-added), because the NITRO API exposes only add, delete, and get operations for this resource.


## Example usage

### Using password / adcpassword (sensitive attributes - persisted in state)

```hcl
variable "nscentralmanagementserver_password" {
  type      = string
  sensitive = true
}

variable "nscentralmanagementserver_adcpassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_nscentralmanagementserver" "tf_centralmgmtserver" {
  type              = "ONPREM"
  ipaddress         = "192.168.10.20"
  username          = "admuser"
  password          = var.nscentralmanagementserver_password
  deviceprofilename = "ns_nsroot_profile"
  adcusername       = "nsroot"
  adcpassword       = var.nscentralmanagementserver_adcpassword
  validatecert      = "YES"
}
```

### Using password_wo / adcpassword_wo (write-only/ephemeral - NOT persisted in state)

The `password_wo` and `adcpassword_wo` attributes provide an ephemeral path for the management-server and ADC credentials. The values are sent to the Citrix ADC but are **not stored in Terraform state**, reducing the risk of secret exposure. To trigger an update when a value changes, increment the corresponding `_wo_version`.

```hcl
variable "nscentralmanagementserver_password" {
  type      = string
  sensitive = true
}

variable "nscentralmanagementserver_adcpassword" {
  type      = string
  sensitive = true
}

resource "citrixadc_nscentralmanagementserver" "tf_centralmgmtserver" {
  type                   = "CLOUD"
  servername             = "adm.cloud.com"
  username               = "admuser"
  password_wo            = var.nscentralmanagementserver_password
  password_wo_version    = 1
  activationcode         = "00000000-0000-0000-0000-000000000000"
  deviceprofilename      = "ns_nsroot_profile"
  adcusername            = "nsroot"
  adcpassword_wo         = var.nscentralmanagementserver_adcpassword
  adcpassword_wo_version = 1
  validatecert           = "YES"
}
```

To rotate a secret, update the variable value and bump the version:

```hcl
resource "citrixadc_nscentralmanagementserver" "tf_centralmgmtserver" {
  type                = "CLOUD"
  servername          = "adm.cloud.com"
  username            = "admuser"
  password_wo         = var.nscentralmanagementserver_password
  password_wo_version = 2  # Bumped to trigger update
  # ... other required attrs
}
```


## Argument Reference

* `type` - (Required) Type of the central management server. Must be either `CLOUD` or `ONPREM` depending on whether the server is on the cloud or on premise. Changing this attribute forces a new resource to be created. Possible values: [ CLOUD, ONPREM ]
* `ipaddress` - (Optional) IP address of the central management server. Exactly one of `ipaddress` or `servername` must be set, and they are mutually exclusive (setting both is rejected). Changing this attribute forces a new resource to be created.
* `servername` - (Optional) Fully qualified domain name of the central management server, or the service-url used to locate the ADM service. Exactly one of `ipaddress` or `servername` must be set, and they are mutually exclusive (setting both is rejected). Changing this attribute forces a new resource to be created.
* `username` - (Optional) Username for access to the central management server. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Changing this attribute forces a new resource to be created.
* `password` - (Optional, Sensitive) Password for access to the central management server. Required for any user account. The value is persisted in Terraform state (encrypted). See also `password_wo` for an ephemeral alternative. Changing this attribute forces a new resource to be created.
* `password_wo` - (Optional, Sensitive, WriteOnly) Same as `password`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `password_wo_version`. If both `password` and `password_wo` are set, `password_wo` takes precedence. Changing this attribute forces a new resource to be created.
* `password_wo_version` - (Optional) An integer version tracker for `password_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger a replacement. Defaults to `1`.
* `adcusername` - (Optional) ADC username used to create the device profile on ADM. Changing this attribute forces a new resource to be created.
* `adcpassword` - (Optional, Sensitive) ADC password used to create the device profile on ADM. The value is persisted in Terraform state (encrypted). See also `adcpassword_wo` for an ephemeral alternative. Changing this attribute forces a new resource to be created.
* `adcpassword_wo` - (Optional, Sensitive, WriteOnly) Same as `adcpassword`, but the value is **not persisted in Terraform state**. Use this for improved secret hygiene. Must be used together with `adcpassword_wo_version`. If both `adcpassword` and `adcpassword_wo` are set, `adcpassword_wo` takes precedence. Changing this attribute forces a new resource to be created.
* `adcpassword_wo_version` - (Optional) An integer version tracker for `adcpassword_wo`. Because write-only values are not stored in state, Terraform cannot detect when the value changes. Increment this version number to signal that the value has changed and trigger a replacement. Defaults to `1`.
* `deviceprofilename` - (Optional) Device profile created on ADM that contains the user name and password of the instance(s). Changing this attribute forces a new resource to be created.
* `activationcode` - (Optional) Activation code used to register to the ADM service. Changing this attribute forces a new resource to be created.
* `validatecert` - (Optional) Validate the server certificate for secure SSL connections. Defaults to `YES`. Changing this attribute forces a new resource to be created.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nscentralmanagementserver. It has the same value as the `type` attribute.


## Import

A nscentralmanagementserver can be imported using its type, e.g.

```shell
terraform import citrixadc_nscentralmanagementserver.tf_centralmgmtserver ONPREM
```
