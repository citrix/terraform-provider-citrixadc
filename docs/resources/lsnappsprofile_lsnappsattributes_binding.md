---
subcategory: "LSN"
---

# Resource: lsnappsprofile_lsnappsattributes_binding

The lsnappsprofile_lsnappsattributes_bindingresource is used to create lsnappsprofile_lsnappsattributes_binding.


## Example usage

```hcl
resource "citrixadc_lsnappsprofile_lsnappsattributes_binding" "tf_lsnappsprofile_lsnappsattributes_binding" {
  appsprofilename    = "my_lsn_profile"
  appsattributesname = "my_lsn_appattributes"
}

```


## Argument Reference

* `appsattributesname` - (Required) Name of the LSN application port ATTRIBUTES command to bind to the specified LSN Appsprofile. Properties of the Appsprofile will be applicable to this APPSATTRIBUTES
* `appsprofilename` - (Required) Name for the LSN application profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN application profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "lsn application profile1" or 'lsn application profile1').


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the lsnappsprofile_lsnappsattributes_binding. It is the concatenation of `appsprofilename` and `appsattributesname` attributes separated by a comma.


## Import

A lsnappsprofile_lsnappsattributes_binding can be imported using its name, e.g.

```shell
terraform import citrixadc_lsnappsprofile_lsnappsattributes_binding.tf_lsnappsprofile_lsnappsattributes_binding my_lsn_profile,my_lsn_appattributes
```
