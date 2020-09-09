---
subcategory: "Responder"
---

# Resource: responderaction

The responderaction resource is used to create responder actions.


## Example usage

```hcl
resource "citrixadc_responderaction" "tf_responderaction" {
  name    = "tf_responderaction"
  type    = "respondwith"
  bypasssafetycheck = "YES"
  target  = "HTTP.REQ.URL.SUFFIX.EQ(\"goodbye\")"
  comment = "some comment"
}
```


## Argument Reference

* `name` - (Optional) Name for the responder action. 
* `type` - (Optional) Type of responder action. Available settings function as follows: * respondwith <target> - Respond to the request with the expression specified as the target. * respondwithhtmlpage - Respond to the request with the uploaded HTML page object specified as the target. * redirect - Redirect the request to the URL specified as the target. * sqlresponse_ok - Send an SQL OK response. * sqlresponse_error - Send an SQL ERROR response. Possible values: [ noop, respondwith, redirect, respondwithhtmlpage, sqlresponse_ok, sqlresponse_error ]
* `target` - (Optional) Expression specifying what to respond with. Typically a URL for redirect policies or a default-syntax expression.  In addition to Citrix ADC default-syntax expressions that refer to information in the request, a stringbuilder expression can contain text and HTML, and simple escape codes that define new lines and paragraphs. Enclose each stringbuilder expression element (either a Citrix ADC default-syntax expression or a string) in double quotation marks. Use the plus (+) character to join the elements. Examples: 1) Respondwith expression that sends an HTTP 1.1 200 OK response: "HTTP/1.1 200 OK\r\n\r\n" 2) Redirect expression that redirects user to the specified web host and appends the request URL to the redirect. "http://backupsite2.com" + HTTP.REQ.URL 3) Respondwith expression that sends an HTTP 1.1 404 Not Found response with the request URL included in the response: "HTTP/1.1 404 Not Found\r\n\r\n"+ "HTTP.REQ.URL.HTTP_URL_SAFE" + "does not exist on the web server." The following requirement applies only to the Citrix ADC CLI: Enclose the entire expression in single quotation marks. (Citrix ADC expression elements should be included inside the single quotation marks for the entire expression, but do not need to be enclosed in double quotation marks.).
* `htmlpage` - (Optional) For respondwithhtmlpage policies, name of the HTML page object to use as the response. You must first import the page object.
* `bypasssafetycheck` - (Optional) Bypass the safety check, allowing potentially unsafe expressions. An unsafe expression in a response is one that contains references to request elements that might not be present in all requests. If a response refers to a missing request element, an empty string is used instead. Possible values: [ YES, NO ]
* `comment` - (Optional) Comment. Any type of information about this responder action.
* `responsestatuscode` - (Optional) HTTP response status code, for example 200, 302, 404, etc. The default value for the redirect action type is 302 and for respondwithhtmlpage is 200.
* `reasonphrase` - (Optional) Expression specifying the reason phrase of the HTTP response. The reason phrase may be a string literal with quotes or a PI expression. For example: "Invalid URL: " + HTTP.REQ.URL.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the responderaction. It has the same value as the `name` attribute.


## Import

A responderaction can be imported using its name, e.g.

```shell
terraform import citrixadc_responderaction.tf_responderaction tf_responderaction
```
