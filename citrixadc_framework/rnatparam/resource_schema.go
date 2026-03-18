package rnatparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RnatparamResourceModel describes the resource data model.
type RnatparamResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Srcippersistency types.String `tfsdk:"srcippersistency"`
	Tcpproxy         types.String `tfsdk:"tcpproxy"`
}

func (r *RnatparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rnatparam resource.",
			},
			"srcippersistency": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable source ip persistency, which enables the Citrix ADC to use the RNAT ips using source ip.",
			},
			"tcpproxy": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features.",
			},
		},
	}
}

func rnatparamGetThePayloadFromtheConfig(ctx context.Context, data *RnatparamResourceModel) network.Rnatparam {
	tflog.Debug(ctx, "In rnatparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	rnatparam := network.Rnatparam{}
	if !data.Srcippersistency.IsNull() {
		rnatparam.Srcippersistency = data.Srcippersistency.ValueString()
	}
	if !data.Tcpproxy.IsNull() {
		rnatparam.Tcpproxy = data.Tcpproxy.ValueString()
	}

	return rnatparam
}

func rnatparamSetAttrFromGet(ctx context.Context, data *RnatparamResourceModel, getResponseData map[string]interface{}) *RnatparamResourceModel {
	tflog.Debug(ctx, "In rnatparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["srcippersistency"]; ok && val != nil {
		data.Srcippersistency = types.StringValue(val.(string))
	} else {
		data.Srcippersistency = types.StringNull()
	}
	if val, ok := getResponseData["tcpproxy"]; ok && val != nil {
		data.Tcpproxy = types.StringValue(val.(string))
	} else {
		data.Tcpproxy = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("rnatparam-config")

	return data
}
