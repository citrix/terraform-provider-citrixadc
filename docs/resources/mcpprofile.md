---
subcategory: "Basic"
---

# Resource: mcpprofile

The mcpprofile resource is used to create and manage MCP (Model Context Protocol) profiles on the Citrix ADC.

## Example usage

```hcl
resource "citrixadc_mcpprofile" "tf_mcpprofile" {
  name      = "my_mcpprofile"
  proxymode = "FORWARD"
  comment   = "example mcp profile"
}
```

## Argument Reference

* `name` - (Required) Name for the mcp profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Changing this forces a new resource to be created.
* `proxymode` - (Optional) Proxy mode for the MCP profile. FORWARD mode replaces Host and URL in backend requests. REVERSE mode passes requests as-is. Possible values: `FORWARD`, `REVERSE`. Default: `FORWARD`.
* `profiletype` - (Optional) Type of MCP profile. Frontend profiles apply to the entity that receives requests from a client. Backend profiles apply to the entity that sends client requests to a server. Possible values: `BACKEND`, `FRONTEND`. Default: `BACKEND`. Changing this forces a new resource to be created.
* `hostreplacement` - (Optional) Whether the Host header should be replaced with the backend MCP server FQDN in FORWARD proxy mode. If mcpProxyMode is FORWARD, this parameter is ENABLED by default. If mcpProxyMode is REVERSE, this parameter is DISABLED and cannot be ENABLED. Possible values: `ENABLED`, `DISABLED`.
* `urlreplacement` - (Optional) Whether the URL should be replaced with the backend MCP server URL in FORWARD proxy mode. If mcpProxyMode is FORWARD, this parameter is ENABLED by default. If mcpProxyMode is REVERSE, this parameter is DISABLED and cannot be ENABLED. Possible values: `ENABLED`, `DISABLED`.
* `protocolversion` - (Optional) MCP protocol version to advertise during monitoring of a mcp server.
* `tokenorapi` - (Optional) If you like to insert Bearer or API token, configure this parameter with full header.
* `comment` - (Optional) Any information about the MCP profile.
* `insertheaderinclientrequest` - (Optional) Whether mcp_token_or_api configuration will be used for MCP requests coming from client. Possible values: `ENABLED`, `DISABLED`. Default: `DISABLED`.
* `tokenorapi_wo` - (Optional, WriteOnly) If you like to insert Bearer or API token, configure this parameter with full header. (write-only ephemeral variant of tokenorapi)
* `tokenorapi_wo_version` - (Optional) Increment this version to signal a tokenorapi_wo update.

## Attribute Reference

In addition to the arguments, the following attributes are exported:

* `id` - The ID of the mcpprofile resource. It has the same value as the `name` attribute.

## Import

A mcpprofile can be imported using its name, e.g.

```shell
terraform import citrixadc_mcpprofile.tf_mcpprofile my_mcpprofile
```
