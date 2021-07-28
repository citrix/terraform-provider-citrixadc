---
subcategory: "Bot"
---

# Resource: botprofile

The Botprofile resource is used to create a collection of profile settings to configure bot management on the appliance.


## Example usage

```hcl
resource "citrixadc_botprofile" "tf_botprofile_name" {

  name                   = "botprofile_name"
  comment                = "My botprofile"
  bot_enable_white_list  = "ON"
  devicefingerprint      = "ON"
}

```


## Argument Reference

* `name` - (Optional) Name for the profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.), pound (#), space ( ), at (@), equals (=), colon (:), and underscore (_) characters. Cannot be changed after the profile is added. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `signature` - (Optional) Name of object containing bot static signature details.
* `errorurl` - (Optional) URL that Bot protection uses as the Error URL.
* `trapurl` - (Optional) URL that Bot protection uses as the Trap URL.
* `comment` - (Optional) Any comments about the purpose of profile, or other useful information about the profile.
* `bot_enable_white_list` - (Optional) Enable white-list bot detection. Possible values: [ on, off ]
* `bot_enable_black_list` - (Optional) Enable black-list bot detection. Possible values: [ on, off ]
* `bot_enable_rate_limit` - (Optional) Enable rate-limit bot detection. Possible values: [ on, off ]
* `devicefingerprint` - (Optional) Enable device-fingerprint bot detection. Possible values: [ on, off ]
* `devicefingerprintaction` - (Optional) Action to be taken for device-fingerprint based bot detection. Possible values: [ NONE, LOG, DROP, REDIRECT, RESET, MITIGATION ]
* `bot_enable_ip_reputation` - (Optional) Enable IP-reputation bot detection. Possible values: [ on, off ]
* `trap` - (Optional) Enable trap bot detection. Possible values: [ on, off ]
* `trapaction` - (Optional) Action to be taken for bot trap based bot detection. Possible values: [ NONE, LOG, DROP, REDIRECT, RESET ]
* `bot_enable_tps` - (Optional) Enable TPS. Possible values: [ on, off ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the botprofile. It has the same value as the `name` attribute.

