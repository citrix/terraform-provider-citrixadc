package mcpprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func McpprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the mcp profile.",
			},
			"proxymode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxy mode for the MCP profile. FORWARD mode replaces Host and URL in backend requests. REVERSE mode passes requests as-is. Possible values = FORWARD, REVERSE",
			},
			"profiletype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of MCP profile. Frontend profiles apply to the entity that receives requests from a client. Backend profiles apply to the entity that sends client requests to a server. Possible values = BACKEND, FRONTEND",
			},
			"hostreplacement": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the Host header should be replaced with the backend MCP server FQDN in FORWARD proxy mode. Possible values = ENABLED, DISABLED",
			},
			"urlreplacement": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the URL should be replaced with the backend MCP server URL in FORWARD proxy mode. Possible values = ENABLED, DISABLED",
			},
			"protocolversion": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MCP protocol version to advertise during monitoring of a mcp server.",
			},
			"tokenorapi": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "If you like to insert Bearer or API token, configure this parameter with full header.",
			},
			"tokenorapi_wo": schema.StringAttribute{
				Optional:    true,
				Description: "If you like to insert Bearer or API token, configure this parameter with full header. (write-only ephemeral variant of tokenorapi)",
			},
			"tokenorapi_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a tokenorapi_wo update.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any information about the MCP profile.",
			},
			"insertheaderinclientrequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether mcp_token_or_api configuration will be used for MCP requests coming from client. Possible values = ENABLED, DISABLED",
			},
		},
	}
}
