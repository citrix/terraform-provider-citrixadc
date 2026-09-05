---
subcategory: "Basic"
---

# Data Source: mcpprofile

The `citrixadc_mcpprofile` data source is used to retrieve information about a specific MCP (Model Context Protocol) profile configured on the Citrix ADC.

## Example Usage

```hcl
data "citrixadc_mcpprofile" "example" {
  name = "my_mcpprofile"
}
```

## Argument Reference

* `name` - (Required) The name of the mcp profile.

## Attribute Reference

In addition to the argument, the following attributes are exported:

* `id` - The ID of the mcp profile.
* `proxymode` - Proxy mode for the MCP profile. FORWARD mode replaces Host and URL in backend requests. REVERSE mode passes requests as-is.
* `profiletype` - Type of MCP profile. Frontend profiles apply to the entity that receives requests from a client. Backend profiles apply to the entity that sends client requests to a server.
* `hostreplacement` - Whether the Host header should be replaced with the backend MCP server FQDN in FORWARD proxy mode.
* `urlreplacement` - Whether the URL should be replaced with the backend MCP server URL in FORWARD proxy mode.
* `protocolversion` - MCP protocol version to advertise during monitoring of a mcp server.
* `tokenorapi` - If you like to insert Bearer or API token, configure this parameter with full header.
* `comment` - Any information about the MCP profile.
* `insertheaderinclientrequest` - Whether mcp_token_or_api configuration will be used for MCP requests coming from client.
