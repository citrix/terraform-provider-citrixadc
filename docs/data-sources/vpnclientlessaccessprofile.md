---
subcategory: "VPN"
---

# Data Source: citrixadc_vpnclientlessaccessprofile

The vpnclientlessaccessprofile data source allows you to retrieve information about a VPN clientless access profile configured on the Citrix ADC.

## Example usage

```terraform
data "citrixadc_vpnclientlessaccessprofile" "tf_vpnclientlessaccessprofile" {
  profilename = "my_clientless_profile"
}

output "requirepersistentcookie" {
  value = data.citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile.requirepersistentcookie
}

output "javascriptrewritepolicylabel" {
  value = data.citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile.javascriptrewritepolicylabel
}
```

## Argument Reference

* `profilename` - (Required) Name for the Citrix Gateway clientless access profile. Must begin with an ASCII alphabetic or underscore (_) character, and must consist only of ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnclientlessaccessprofile. It has the same value as the `profilename` attribute.
* `clientconsumedcookies` - Specify the name of the pattern set containing the names of the cookies, which are allowed between the client and the server. If a pattern set is not specified, Citrix Gateway does not allow any cookies between the client and the server. A cookie that is not specified in the pattern set is handled by Citrix Gateway on behalf of the client.
* `javascriptrewritepolicylabel` - Name of the configured JavaScript rewrite policy label.  If you do not specify a policy label name, then JAVA scripts are not rewritten.
* `regexforfindingcustomurls` - Name of the pattern set that contains the regular expressions, which match the URLs in the custom content type other than HTML, CSS, XML, XCOMP, and JavaScript. The custom content type should be included in the patset ns_cvpn_custom_content_types.
* `regexforfindingurlincss` - Name of the pattern set that contains the regular expressions, which match the URL in the CSS.
* `regexforfindingurlinjavascript` - Name of the pattern set that contains the regular expressions, which match the URL in Java script.
* `regexforfindingurlinxcomponent` - Name of the pattern set that contains the regular expressions, which match the URL in X Component.
* `regexforfindingurlinxml` - Name of the pattern set that contains the regular expressions, which match the URL in XML.
* `reqhdrrewritepolicylabel` - Name of the configured Request rewrite policy label.  If you do not specify a policy label name, then requests are not rewritten.
* `requirepersistentcookie` - Specify whether a persistent session cookie is set and accepted for clientless access. If this parameter is set to ON, COM objects, such as MSOffice, which are invoked by the browser can access the files using clientless access. Use caution because the persistent cookie is stored on the disk.
* `reshdrrewritepolicylabel` - Name of the configured Response rewrite policy label.
* `urlrewritepolicylabel` - Name of the configured URL rewrite policy label. If you do not specify a policy label name, then URLs are not rewritten.
