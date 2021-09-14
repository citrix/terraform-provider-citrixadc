---
subcategory: "Policy"
---

# Resource: policyhttpcallout

The policyhttpcallout resource is used to create an expression element that, when evaluated, sends an HTTP request to a specified service and receives an HTTP response from the service.


## Example usage

```hcl
resource "citrixadc_policyhttpcallout" "tf_policyhttpcallout" {
	name = "tf_policyhttpcallout"
	bodyexpr = "client.ip.src"
	cacheforsecs = 5
	comment = "Demo comment"
	headers = ["cip(client.ip.src)", "hdr(http.req.header(\"HDR\"))"]
	hostexpr = "http.req.header(\"Host\")"
	httpmethod = "GET"
	parameters = ["param1(\"name1\")", "param2(http.req.header(\"hdr\"))"]
	resultexpr = "http.res.body(10000).length"
	returntype = "TEXT"
	scheme = "http"
	vserver = citrixadc_lbvserver.tf_lbvserver.name
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	name = "tf_lbvserver"
	ipv46 = "10.202.11.11"
	port = 80
	servicetype = "HTTP"
}
```


## Argument Reference

* `name` - (Required) Name for the HTTP callout. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or HTTP callout.
* `ipaddress` - (Optional) IP Address of the server (callout agent) to which the callout is sent. Can be an IPv4 or IPv6 address. Mutually exclusive with the Virtual Server parameter. Therefore, you cannot set the <IP Address, Port> and the Virtual Server in the same HTTP callout.
* `port` - (Optional) Server port to which the HTTP callout agent is mapped. Mutually exclusive with the Virtual Server parameter. Therefore, you cannot set the <IP Address, Port> and the Virtual Server in the same HTTP callout.
* `vserver` - (Optional) Name of the load balancing, content switching, or cache redirection virtual server (the callout agent) to which the HTTP callout is sent. The service type of the virtual server must be HTTP. Mutually exclusive with the IP address and port parameters. Therefore, you cannot set the <IP Address, Port> and the Virtual Server in the same HTTP callout.
* `returntype` - (Optional) Type of data that the target callout agent returns in response to the callout. Available settings function as follows: * TEXT - Treat the returned value as a text string. * NUM - Treat the returned value as a number. * BOOL - Treat the returned value as a Boolean value. Note: You cannot change the return type after it is set. Possible values: [ BOOL, NUM, TEXT ]
* `httpmethod` - (Optional) Method used in the HTTP request that this callout sends.  Mutually exclusive with the full HTTP request expression. Possible values: [ GET, POST ]
* `hostexpr` - (Optional) String expression to configure the Host header. Can contain a literal value (for example, 10.101.10.11) or a derived value (for example, http.req.header("Host")). The literal value can be an IP address or a fully qualified domain name. Mutually exclusive with the full HTTP request expression.
* `urlstemexpr` - (Optional) String expression for generating the URL stem. Can contain a literal string (for example, "/mysite/index.html") or an expression that derives the value (for example, http.req.url). Mutually exclusive with the full HTTP request expression.
* `headers` - (Optional) One or more headers to insert into the HTTP request. Each header is specified as "name(expr)", where expr is an expression that is evaluated at runtime to provide the value for the named header. You can configure a maximum of eight headers for an HTTP callout. Mutually exclusive with the full HTTP request expression.
* `parameters` - (Optional) One or more query parameters to insert into the HTTP request URL (for a GET request) or into the request body (for a POST request). Each parameter is specified as "name(expr)", where expr is an expression that is evaluated at run time to provide the value for the named parameter (name=value). The parameter values are URL encoded. Mutually exclusive with the full HTTP request expression.
* `bodyexpr` - (Optional) An advanced string expression for generating the body of the request. The expression can contain a literal string or an expression that derives the value (for example, client.ip.src). Mutually exclusive with -fullReqExpr.
* `fullreqexpr` - (Optional) Exact HTTP request, in the form of an expression, which the Citrix ADC sends to the callout agent. If you set this parameter, you must not include HTTP method, host expression, URL stem expression, headers, or parameters. The request expression is constrained by the feature for which the callout is used. For example, an HTTP.RES expression cannot be used in a request-time policy bank or in a TCP content switching policy bank. The Citrix ADC does not check the validity of this request. You must manually validate the request.
* `scheme` - (Optional) Type of scheme for the callout server. Possible values: [ http, https ]
* `resultexpr` - (Optional) Expression that extracts the callout results from the response sent by the HTTP callout agent. Must be a response based expression, that is, it must begin with HTTP.RES. The operations in this expression must match the return type. For example, if you configure a return type of TEXT, the result expression must be a text based expression. If the return type is NUM, the result expression (resultExpr) must return a numeric value, as in the following example: http.res.body(10000).length.
* `cacheforsecs` - (Optional) Duration, in seconds, for which the callout response is cached. The cached responses are stored in an integrated caching content group named "calloutContentGroup". If no duration is configured, the callout responses will not be cached unless normal caching configuration is used to cache them. This parameter takes precedence over any normal caching configuration that would otherwise apply to these responses. Note that the calloutContentGroup definition may not be modified or removed nor may it be used with other cache policies.
* `comment` - (Optional) Any comments to preserve information about this HTTP callout.


## Attribute Reference

In addition to the arguments, the following attributes are available:

* `id` - The id of the policyhttpcallout. It has the same value as the `name` attribute.


## Import

A policyhttpcallout can be imported using its name, e.g.

```shell
terraform import citrixadc_policyhttpcallout.tf_policyhttpcallout tf_policyhttpcallout
```
