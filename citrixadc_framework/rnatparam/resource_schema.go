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
			// SDK v2 baseline: Optional+Computed with NO Default (value is read
			// from the ADC). The auto-gen wrongly added a Default here; removed to
			// preserve the SDK v2 backward-compatible contract.
			"srcippersistency": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Default added so removing the attribute from config produces a
				// plan diff, allowing Update to fire the NITRO unset. Value matches
				// the NITRO spec default for srcippersistency.
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable source ip persistency, which enables the Citrix ADC to use the RNAT ips using source ip.",
			},
			"tcpproxy": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Default added so removing the attribute from config produces a
				// plan diff, allowing Update to fire the NITRO unset. Value matches
				// the NITRO spec default for tcpproxy.
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable TCP proxy, which enables the Citrix ADC to optimize the RNAT TCP traffic by using Layer 4 features.",
			},
		},
	}
}

func rnatparamGetThePayloadFromthePlan(ctx context.Context, data *RnatparamResourceModel) network.Rnatparam {
	tflog.Debug(ctx, "In rnatparamGetThePayloadFromthePlan Function")

	// Create API request body from the model
	rnatparam := network.Rnatparam{}
	if !data.Srcippersistency.IsNull() && !data.Srcippersistency.IsUnknown() {
		rnatparam.Srcippersistency = data.Srcippersistency.ValueString()
	}
	if !data.Tcpproxy.IsNull() && !data.Tcpproxy.IsUnknown() {
		rnatparam.Tcpproxy = data.Tcpproxy.ValueString()
	}

	return rnatparam
}

func rnatparamSetAttrFromGet(ctx context.Context, data *RnatparamResourceModel, getResponseData map[string]interface{}) *RnatparamResourceModel {
	tflog.Debug(ctx, "In rnatparamSetAttrFromGet Function")

	// Convert API response to model.
	// Omit-on-default guard: NITRO may omit an attribute from GET when it holds a
	// default value. Only null the field when it was Unknown (Computed, unresolved);
	// never clobber a value the user explicitly configured.
	if val, ok := getResponseData["srcippersistency"]; ok && val != nil {
		data.Srcippersistency = types.StringValue(val.(string))
	} else if data.Srcippersistency.IsUnknown() {
		data.Srcippersistency = types.StringNull()
	}
	if val, ok := getResponseData["tcpproxy"]; ok && val != nil {
		data.Tcpproxy = types.StringValue(val.(string))
	} else if data.Tcpproxy.IsUnknown() {
		data.Tcpproxy = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("rnatparam-config")

	return data
}

// rnatparamSetAttrFromGetForDatasource copies all values from the GET response
// (nulling any absent field) and sets the datasource ID. Unlike the resource
// setter it never preserves prior values, so a datasource read always reflects
// the live ADC state.
func rnatparamSetAttrFromGetForDatasource(ctx context.Context, data *RnatparamResourceModel, getResponseData map[string]interface{}) *RnatparamResourceModel {
	tflog.Debug(ctx, "In rnatparamSetAttrFromGetForDatasource Function")

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

	data.Id = types.StringValue("rnatparam-config")

	return data
}
