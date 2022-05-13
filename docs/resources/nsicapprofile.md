---
subcategory: "NS"
---

# Resource: nsicapprofile

The nsicapprofile resource is used to create ICAP profile resource.


## Example usage

```hcl
resource "citrixadc_nsicapprofile" "tf_nsicapprofile" {
  name             = "tf_nsicapprofile"
  uri              = "/example"
  mode             = "REQMOD"
  reqtimeout       = 4
  reqtimeoutaction = "RESET"
  preview          = "ENABLED"
  previewlength    = 4096
}
```


## Argument Reference

* `name` - (Required) Name for an ICAP profile. Must begin with a letter, number, or the underscore \(_\) character. Other characters allowed, after the first character, are the hyphen \(-\), period \(.\), hash \(\#\), space \( \), at \(@\), colon \(:\), and equal \(=\) characters. The name of a ICAP profile cannot be changed after it is created. CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \(for example, "my icap profile" or 'my icap profile'\). Minimum length =  1 Maximum length =  127
* `uri` - (Required) URI representing icap service. It is a mandatory argument while creating an icapprofile. Minimum length =  1
* `mode` - (Required) ICAP Mode of operation. It is a mandatory argument while creating an icapprofile. Possible values: [ REQMOD, RESPMOD ]
* `preview` - (Optional) Enable or Disable preview header with ICAP request. This feature allows an ICAP server to see the beginning of a transaction, then decide if it wants to opt-out of the transaction early instead of receiving the remainder of the request message. Possible values: [ ENABLED, DISABLED ]
* `previewlength` - (Optional) Value of Preview Header field. Citrix ADC uses the minimum of this set value and the preview size received on OPTIONS response. Minimum value =  0 Maximum value =  4294967294
* `hostheader` - (Optional) ICAP Host Header. Minimum length =  1
* `useragent` - (Optional) ICAP User Agent Header String. Minimum length =  1
* `queryparams` - (Optional) Query parameters to be included with ICAP request URI. Entered values should be in arg=value format. For more than one parameters, add & separated values. e.g.: arg1=val1&arg2=val2. Minimum length =  1
* `connectionkeepalive` - (Optional) If enabled, Citrix ADC keeps the ICAP connection alive after a transaction to reuse it to send next ICAP request. Possible values: [ ENABLED, DISABLED ]
* `allow204` - (Optional) Enable or Disable sending Allow: 204 header in ICAP request. Possible values: [ ENABLED, DISABLED ]
* `inserticapheaders` - (Optional) Insert custom ICAP headers in the ICAP request to send to ICAP server. The headers can be static or can be dynamically constructed using PI Policy Expression. For example, to send static user agent and Client's IP address, the expression can be specified as "User-Agent: NS-ICAP-Client/V1.0\r\nX-Client-IP: "+CLIENT.IP.SRC+"\r\n". The Citrix ADC does not check the validity of the specified header name-value. You must manually validate the specified header syntax. Minimum length =  1
* `inserthttprequest` - (Optional) Exact HTTP request, in the form of an expression, which the Citrix ADC encapsulates and sends to the ICAP server. If you set this parameter, the ICAP request is sent using only this header. This can be used when the HTTP header is not available to send or ICAP server only needs part of the incoming HTTP request. The request expression is constrained by the feature for which it is used. The Citrix ADC does not check the validity of this request. You must manually validate the request. Minimum length =  1
* `reqtimeout` - (Optional) Time, in seconds, within which the remote server should respond to the ICAP-request. If the Netscaler does not receive full response with this time, the specified request timeout action is performed. Zero value disables this timeout functionality. Minimum value =  0 Maximum value =  86400
* `reqtimeoutaction` - (Optional) Name of the action to perform if the Vserver/Server representing the remote service does not respond with any response within the timeout value configured. The Supported actions are * BYPASS - This Ignores the remote server response and sends the request/response to Client/Server. * If the ICAP response with Encapsulated headers is not received within the request-timeout value configured, this Ignores the remote ICAP server response and sends the Full request/response to Server/Client. * RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired. * DROP - Drop the request without sending a response to the user. Possible values: [ BYPASS, DROP, RESET ]
* `logaction` - (Optional) Name of the audit message action which would be evaluated on receiving the ICAP response to emit the logs.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the nsicapprofile. It has the same value as the `name` attribute.


## Import

A nsicapprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_nsicapprofile.tf_nsicapprofile tf_nsicapprofile
```
