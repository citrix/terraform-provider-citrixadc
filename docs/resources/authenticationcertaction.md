---
subcategory: "Authentication"
---

# Resource: authenticationcertaction

The authenticationcertaction resource is used to create authentication certaction resource.


## Example usage

```hcl
resource "citrixadc_authenticationcertaction" "tf_certaction" {
  name                       = "tf_certaction"
  twofactor                  = "ON"
  defaultauthenticationgroup = "new_group"
  usernamefield              = "Subject:CN"
  groupnamefield             = "subject:grp"
}
```


## Argument Reference

* `name` - (Required) Name for the client cert authentication server profile (action).  Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after certifcate action is created.  The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my authentication action" or 'my authentication action').
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups.
* `groupnamefield` - (Optional) Client-cert field from which the group is extracted.  Must be set to either ""Subject"" and ""Issuer"" (include both sets of double quotation marks). Format: <field>:<subfield>
* `twofactor` - (Optional) Enables or disables two-factor authentication.  Two factor authentication is client cert authentication followed by password authentication.
* `usernamefield` - (Optional) Client-cert field from which the username is extracted. Must be set to either ""Subject"" and ""Issuer"" (include both sets of double quotation marks). Format: <field>:<subfield>.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the authenticationcertaction. It has the same value as the `name` attribute.


## Import

A authenticationcertaction can be imported using its name, e.g.

```shell
terraform import citrixadc_authenticationcertaction.tf_certaction tf_certaction
```
