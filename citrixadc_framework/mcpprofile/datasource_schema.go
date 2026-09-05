package mcpprofile

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
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

type McpprofileDataSourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Comment                     types.String `tfsdk:"comment"`
	Hostreplacement             types.String `tfsdk:"hostreplacement"`
	Insertheaderinclientrequest types.String `tfsdk:"insertheaderinclientrequest"`
	Name                        types.String `tfsdk:"name"`
	Profiletype                 types.String `tfsdk:"profiletype"`
	Protocolversion             types.String `tfsdk:"protocolversion"`
	Proxymode                   types.String `tfsdk:"proxymode"`
	Tokenorapi                  types.String `tfsdk:"tokenorapi"`
	Urlreplacement              types.String `tfsdk:"urlreplacement"`
}

func mcpprofileDataSourceSetAttrFromGet(ctx context.Context, data *McpprofileDataSourceModel, getResponseData map[string]interface{}) *McpprofileDataSourceModel {
	tflog.Debug(ctx, "In mcpprofileDataSourceSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["proxymode"]; ok && val != nil {
		data.Proxymode = types.StringValue(val.(string))
	} else {
		data.Proxymode = types.StringNull()
	}
	if val, ok := getResponseData["profiletype"]; ok && val != nil {
		data.Profiletype = types.StringValue(val.(string))
	} else {
		data.Profiletype = types.StringNull()
	}
	if val, ok := getResponseData["hostreplacement"]; ok && val != nil {
		data.Hostreplacement = types.StringValue(val.(string))
	} else {
		data.Hostreplacement = types.StringNull()
	}
	if val, ok := getResponseData["urlreplacement"]; ok && val != nil {
		data.Urlreplacement = types.StringValue(val.(string))
	} else {
		data.Urlreplacement = types.StringNull()
	}
	if val, ok := getResponseData["protocolversion"]; ok && val != nil {
		data.Protocolversion = types.StringValue(val.(string))
	} else {
		data.Protocolversion = types.StringNull()
	}
	// tokenorapi is a secret returned by NITRO only in an encrypted form that does
	// not match the plaintext supplied in configuration. When the user configured a
	// plaintext value, retain it (avoids a "provider produced inconsistent result
	// after apply"); otherwise leave it null (the value may have been supplied via
	// value.
	if !data.Tokenorapi.IsNull() && !data.Tokenorapi.IsUnknown() {
		// retain the configured plaintext value
	} else {
		data.Tokenorapi = types.StringNull()
	}
	// returned by NITRO; retain them from config/state.
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["insertheaderinclientrequest"]; ok && val != nil {
		data.Insertheaderinclientrequest = types.StringValue(val.(string))
	} else {
		data.Insertheaderinclientrequest = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
