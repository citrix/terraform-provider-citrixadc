---
subcategory: "VPN"
---

# Resource: vpnurl

The vpnurl resource is used to create a new vpn-url.


## Example usage

```hcl
resource "citrixadc_vpnurl" "tf_vpnurl" {
    actualurl = "www.citrix.com"
    appjson = "xyz"
    applicationtype = "CVPN"
    clientlessaccess = "OFF"
    comment = "Testing"
    linkname = "Description"
    ssotype = "unifiedgateway"
    urlname = "Firsturl"
    vservername = "server1"
}
```


## Argument Reference

* `urlname` - (Required) Name of the bookmark link.
* `actualurl` - (Required) Web address for the bookmark link.
* `linkname` - (Required) Description of the bookmark link. The description appears in the Access Interface.
* `appjson` - (Optional) To store the template details in the json format.
* `applicationtype` - (Optional) The type of application this VPN URL represents. Possible values are CVPN/SaaS/VPN
* `clientlessaccess` - (Optional) If clientless access to the resource hosting the link is allowed, also use clientless access for the bookmarked web address in the Secure Client Access based session. Allows single sign-on and other HTTP processing on Citrix Gateway for HTTPS resources.
* `comment` - (Optional) Any comments associated with the bookmark link.
* `iconurl` - (Optional) URL to fetch icon file for displaying this resource.
* `samlssoprofile` - (Optional) Profile to be used for doing SAML SSO
* `ssotype` - (Optional) Single sign on type for unified gateway
* `vservername` - (Optional) Name of the associated LB/CS vserver


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnurl. It has the same value as the `urlname` attribute.


## Import

A vpnurl can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnurl.tf_vpnurl Firsturl
```
