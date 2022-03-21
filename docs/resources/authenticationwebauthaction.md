---
subcategory: "Authentication"
---

# Resource: authenticationwebauthaction

The authenticationwebauthaction resource is used to create authentication webauthaction resource.


## Example usage

```hcl
resource "citrixadc_authenticationwebauthaction" "tf_webauthaction" {
  name                       = "tf_webauthaction"
  serverip                   = "1.2.3.4"
  serverport                 = 8080
  fullreqexpr                = "TRUE"
  scheme                     = "http"
  successrule                = "http.RES.STATUS.EQ(200)"
  defaultauthenticationgroup = "new_group"
}
```


## Argument Reference

* `name` - (Required) Name for the Web Authentication action.  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the profile is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
* `scheme` - (Required) Type of scheme for the web server.
* `serverip` - (Required) IP address of the web server to be used for authentication.
* `serverport` - (Required) Port on which the web server accepts connections.
* `successrule` - (Required) Expression, that checks to see if authentication is successful.
* `attribute1` - (Optional) Expression that would be evaluated to extract attribute1 from the webauth response
* `attribute10` - (Optional) Expression that would be evaluated to extract attribute10 from the webauth response
* `attribute11` - (Optional) Expression that would be evaluated to extract attribute11 from the webauth response
* `attribute12` - (Optional) Expression that would be evaluated to extract attribute12 from the webauth response
* `attribute13` - (Optional) Expression that would be evaluated to extract attribute13 from the webauth response
* `attribute14` - (Optional) Expression that would be evaluated to extract attribute14 from the webauth response
* `attribute15` - (Optional) Expression that would be evaluated to extract attribute15 from the webauth response
* `attribute16` - (Optional) Expression that would be evaluated to extract attribute16 from the webauth response
* `attribute2` - (Optional) Expression that would be evaluated to extract attribute2 from the webauth response
* `attribute3` - (Optional) Expression that would be evaluated to extract attribute3 from the webauth response
* `attribute4` - (Optional) Expression that would be evaluated to extract attribute4 from the webauth response
* `attribute5` - (Optional) Expression that would be evaluated to extract attribute5 from the webauth response
* `attribute6` - (Optional) Expression that would be evaluated to extract attribute6 from the webauth response
* `attribute7` - (Optional) Expression that would be evaluated to extract attribute7 from the webauth response
* `attribute8` - (Optional) Expression that would be evaluated to extract attribute8 from the webauth response
* `attribute9` - (Optional) Expression that would be evaluated to extract attribute9 from the webauth response
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `fullreqexpr` - (Optional) Exact HTTP request, in the form of an expression, which the Citrix ADC sends to the authentication server. The Citrix ADC does not check the validity of this request. One must manually validate the request.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationwebauthaction. It has the same value as the `name` attribute.


## Import

A authenticationwebauthaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationwebauthaction.tf_webauthaction tf_webauthaction
```
