---
subcategory: "Policy"
---

# Data Source `policyhttpcallout`

The policyhttpcallout data source allows you to retrieve information about an existing HTTP callout configuration.

## Example usage

```terraform
data "citrixadc_policyhttpcallout" "tf_policyhttpcallout" {
  name = "my_http_callout"
}

output "callout_name" {
  value = data.citrixadc_policyhttpcallout.tf_policyhttpcallout.name
}

output "callout_method" {
  value = data.citrixadc_policyhttpcallout.tf_policyhttpcallout.httpmethod
}
```

## Argument Reference

* `name` - (Required) Name for the HTTP callout. Not case sensitive. Must begin with an ASCII letter or underscore (_) character.

## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policyhttpcallout resource. It has the same value as the `name` attribute.
* `bodyexpr` - An advanced string expression for generating the body of the request.
* `cacheforsecs` - Duration, in seconds, for which the callout response is cached.
* `comment` - Any comments to preserve information about this HTTP callout.
* `fullreqexpr` - Exact HTTP request, in the form of an expression, which the Citrix ADC sends to the callout agent.
* `headers` - One or more headers to insert into the HTTP request.
* `hostexpr` - String expression to configure the Host header.
* `httpmethod` - Method used in the HTTP request that this callout sends. Possible values: [ GET, POST ]
* `ipaddress` - IP Address of the server (callout agent) to which the callout is sent.
* `parameters` - One or more query parameters to insert into the HTTP request URL.
* `port` - Port number on which the HTTP server (callout agent) listens.
* `resultexpr` - Expression that extracts the callout results from the response sent by the HTTP callout agent.
* `returntype` - Type of data that the target callout agent returns. Possible values: [ BOOL, NUM, TEXT ]
* `scheme` - Type of scheme for the callout server. Possible values: [ http, https ]
* `urlstemexpr` - String expression for generating the URL stem.
* `vserver` - Name of the load balancing or content switching virtual server or service to which the HTTP callout agent is bound.
