---
subcategory: "Responder"
---

# Data Source: citrixadc_responderaction

The `citrixadc_responderaction` data source is used to retrieve information about an existing responder action configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_responderaction" "example" {
  name = "my_responderaction"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the responder action to retrieve.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the responder action (same as name).
* `type` - Type of responder action. Possible values include:
  * `respondwith` - Respond to the request with the expression specified as the target.
  * `respondwithhtmlpage` - Respond to the request with the uploaded HTML page object specified as the target.
  * `redirect` - Redirect the request to the URL specified as the target.
  * `sqlresponse_ok` - Send an SQL OK response.
  * `sqlresponse_error` - Send an SQL ERROR response.
* `target` - Expression specifying what to respond with. Typically a URL for redirect policies or a default-syntax expression.
* `bypasssafetycheck` - Bypass the safety check, allowing potentially unsafe expressions. Possible values: `YES`, `NO`.
* `comment` - Comment. Any type of information about this responder action.
* `headers` - One or more headers to insert into the HTTP response. Each header is specified as "name(expr)", where expr is an expression that is evaluated at runtime to provide the value for the named header.
* `htmlpage` - For respondwithhtmlpage policies, name of the HTML page object to use as the response.
* `newname` - New name for the responder action.
* `reasonphrase` - Expression specifying the reason phrase of the HTTP response. The reason phrase may be a string literal with quotes or a PI expression.
* `responsestatuscode` - HTTP response status code, for example 200, 302, 404, etc. The default value for the redirect action type is 302 and for respondwithhtmlpage is 200.
