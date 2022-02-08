---
subcategory: "Vpn"
---

# Resource: vpnurlaction

The vpnurlaction resource is used to create vpn url action.


## Example usage

```hcl
resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
  name             = "tf_vpnurlaction"
  linkname         = "new_link"
  actualurl        = "www.citrix.com"
  applicationtype  = "CVPN"
  clientlessaccess = "OFF"
  comment          = "Testing"
  ssotype          = "unifiedgateway"
  vservername      = "vserver1"
}
```


## Argument Reference

* `name` - (Required) Name of the bookmark link.
* `linkname` - (Required) Description of the bookmark link. The description appears in the Access Interface.
* `actualurl` - (Required) Web address for the bookmark link.
* `applicationtype` - (Optional) The type of application this VPN URL represents. Possible values are CVPN/SaaS/VPN
* `clientlessaccess` - (Optional) If clientless access to the resource hosting the link is allowed, also use clientless access for the bookmarked web address in the Secure Client Access based session. Allows single sign-on and other HTTP processing on NetScaler Gateway for HTTPS resources.
* `comment` - (Optional) Any comments associated with the bookmark link.
* `iconurl` - (Optional) URL to fetch icon file for displaying this resource.
* `newname` - (Optional) New name for the vpn urlAction. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.  The following requirement applies only to the NetScaler CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my vpnurl action" or 'my vpnurl action').
* `samlssoprofile` - (Optional) Profile to be used for doing SAML SSO
* `ssotype` - (Optional) Single sign on type for unified gateway
* `vservername` - (Optional) Name of the associated vserver to handle selfAuth SSO


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnurlaction. It has the same value as the `name` attribute.


## Import

A vpnurlaction can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnurlaction.tf_vpnurlaction tf_vpnurlaction
```
