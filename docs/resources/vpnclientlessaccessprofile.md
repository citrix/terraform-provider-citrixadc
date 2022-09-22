---
subcategory: "VPN"
---

# Resource: vpnclientlessaccessprofile

The vpnclientlessaccessprofile resource is used to ads a collection of settings that allows clientless access to a given application..


## Example usage

```hcl
resource "citrixadc_vpnclientlessaccessprofile" "tf_vpnclientlessaccessprofile" {
	profilename = "tf_vpnclientlessaccessprofile"
	requirepersistentcookie = "ON"
}
```


## Argument Reference

* `profilename` - (Required) Name for the Citrix Gateway clientless access profile. Must begin with an ASCII alphabetic or underscore (_) character, and must consist only of ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my profile" or 'my profile').
* `urlrewritepolicylabel` - (Optional) Name of the configured URL rewrite policy label. If you do not specify a policy label name, then URLs are not rewritten.
* `javascriptrewritepolicylabel` - (Optional) Name of the configured JavaScript rewrite policy label.  If you do not specify a policy label name, then JAVA scripts are not rewritten.
* `reqhdrrewritepolicylabel` - (Optional) Name of the configured Request rewrite policy label.  If you do not specify a policy label name, then requests are not rewritten.
* `reshdrrewritepolicylabel` - (Optional) Name of the configured Response rewrite policy label.
* `regexforfindingurlinjavascript` - (Optional) Name of the pattern set that contains the regular expressions, which match the URL in Java script.
* `regexforfindingurlincss` - (Optional) Name of the pattern set that contains the regular expressions, which match the URL in the CSS.
* `regexforfindingurlinxcomponent` - (Optional) Name of the pattern set that contains the regular expressions, which match the URL in X Component.
* `regexforfindingurlinxml` - (Optional) Name of the pattern set that contains the regular expressions, which match the URL in XML.
* `regexforfindingcustomurls` - (Optional) Name of the pattern set that contains the regular expressions, which match the URLs in the custom content type other than HTML, CSS, XML, XCOMP, and JavaScript. The custom content type should be included in the patset ns_cvpn_custom_content_types.
* `clientconsumedcookies` - (Optional) Specify the name of the pattern set containing the names of the cookies, which are allowed between the client and the server. If a pattern set is not specified, Citrix Gateway does not allow any cookies between the client and the server. A cookie that is not specified in the pattern set is handled by Citrix Gateway on behalf of the client.
* `requirepersistentcookie` - (Optional) Specify whether a persistent session cookie is set and accepted for clientless access. If this parameter is set to ON, COM objects, such as MSOffice, which are invoked by the browser can access the files using clientless access. Use caution because the persistent cookie is stored on the disk. Possible values: [ on, off ]


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the vpnclientlessaccessprofile. It has the same value as the `name` attribute.


## Import

A vpnclientlessaccessprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_vpnclientlessaccessprofile.tf_vpnclientlessaccessprofile tf_vpnclientlessaccessprofile
```
