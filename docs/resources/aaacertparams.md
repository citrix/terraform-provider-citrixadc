---
subcategory: "AAA"
---

# Resource: aaacertparams

The aaacertparams resource is used to create aaacertparams.


## Example usage

```hcl
resource "citrixadc_aaacertparams" "tf_aaacertparams" {
  usernamefield              = "Subject:CN"
  groupnamefield             = "Subject:OU"
  defaultauthenticationgroup = 50
}
```


## Argument Reference

* `usernamefield` - (Optional) Client certificate field that contains the username, in the format <field>:<subfield>. .
* `groupnamefield` - (Optional) Client certificate field that specifies the group, in the format <field>:<subfield>.
* `defaultauthenticationgroup` - (Optional) This is the default group that is chosen when the authentication succeeds in addition to extracted groups. Maximum length =  64


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the aaacertparams. It is a unique string prefixed with `tf-aaacertparams-`.
