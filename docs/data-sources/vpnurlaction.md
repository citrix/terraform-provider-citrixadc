---
subcategory: "VPN"
---

# Data Source: vpnurlaction

The vpnurlaction data source allows you to retrieve information about a VPN URL action.

## Example usage

```terraform
data "citrixadc_vpnurlaction" "example" {
  name = "tf_vpnurlaction"
}

output "actualurl" {
  value = data.citrixadc_vpnurlaction.example.actualurl
}

output "linkname" {
  value = data.citrixadc_vpnurlaction.example.linkname
}
```

## Argument Reference

* `name` - (Required) Name of the bookmark link.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `actualurl` - Web address for the bookmark link.
* `applicationtype` - The type of application this VPN URL represents. Possible values are CVPN/SaaS/VPN.
* `clientlessaccess` - If clientless access to the resource hosting the link is allowed, also use clientless access for the bookmarked web address in the Secure Client Access based session. Allows single sign-on and other HTTP processing on NetScaler Gateway for HTTPS resources.
* `comment` - Any comments associated with the bookmark link.
* `iconurl` - URL to fetch icon file for displaying this resource.
* `linkname` - Description of the bookmark link. The description appears in the Access Interface.
* `newname` - New name for the vpn urlAction. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
* `samlssoprofile` - Profile to be used for doing SAML SSO.
* `ssotype` - Single sign on type for unified gateway.
* `vservername` - Name of the associated vserver to handle selfAuth SSO.
