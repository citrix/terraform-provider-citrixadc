---
subcategory: "NS"
---

# Data Source: citrixadc_nsicapprofile

This data source is used to retrieve information about an existing NS ICAP profile.

## Example Usage

```hcl
data "citrixadc_nsicapprofile" "example" {
  name = "my_icap_profile"
}
```

## Argument Reference

* `name` - (Required) Name for an ICAP profile. Must begin with a letter, number, or the underscore (_) character. Other characters allowed, after the first character, are the hyphen (-), period (.), hash (#), space ( ), at (@), colon (:), and equal (=) characters.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the ICAP profile (same as `name`).
* `allow204` - Enable or Disable sending Allow: 204 header in ICAP request.
* `connectionkeepalive` - If enabled, Citrix ADC keeps the ICAP connection alive after a transaction to reuse it to send next ICAP request.
* `hostheader` - ICAP Host Header.
* `inserthttprequest` - Exact HTTP request, in the form of an expression, which the Citrix ADC encapsulates and sends to the ICAP server.
* `inserticapheaders` - Insert custom ICAP headers in the ICAP request to send to ICAP server.
* `logaction` - Name of the audit message action which would be evaluated on receiving the ICAP response to emit the logs.
* `mode` - ICAP Mode of operation. It is a mandatory argument while creating an icapprofile.
* `preview` - Enable or Disable preview header with ICAP request.
* `previewlength` - Value of Preview Header field. Citrix ADC uses the minimum of this set value and the preview size received on OPTIONS response.
* `queryparams` - Query parameters to be included with ICAP request URI.
* `reqtimeout` - Time, in seconds, within which the remote server should respond to the ICAP-request.
* `reqtimeoutaction` - Name of the action to perform if the Vserver/Server representing the remote service does not respond with any response within the timeout value configured.
* `uri` - URI representing icap service. It is a mandatory argument while creating an icapprofile.
* `useragent` - ICAP User Agent Header String.
