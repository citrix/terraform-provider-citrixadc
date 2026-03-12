---
subcategory: "VPN"
---

# Data Source: vpnurl

The vpnurl data source allows you to retrieve information about a VPN URL bookmark link.

## Example usage

```terraform
data "citrixadc_vpnurl" "example" {
  urlname = "Firsturl"
}

output "actualurl" {
  value = data.citrixadc_vpnurl.example.actualurl
}

output "linkname" {
  value = data.citrixadc_vpnurl.example.linkname
}
```

## Argument Reference

* `urlname` - (Required) Name of the bookmark link.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `actualurl` - Web address for the bookmark link.
* `appjson` - To store the template details in the json format.
* `applicationtype` - The type of application this VPN URL represents. Possible values are CVPN/SaaS/VPN.
* `clientlessaccess` - If clientless access to the resource hosting the link is allowed, also use clientless access for the bookmarked web address in the Secure Client Access based session. Allows single sign-on and other HTTP processing on Citrix Gateway for HTTPS resources.
* `comment` - Any comments associated with the bookmark link.
* `iconurl` - URL to fetch icon file for displaying this resource.
* `linkname` - Description of the bookmark link. The description appears in the Access Interface.
* `samlssoprofile` - Profile to be used for doing SAML SSO.
* `ssotype` - Single sign on type for unified gateway.
* `vservername` - Name of the associated LB/CS vserver.
