package lsnip6profile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Lsnip6profileResourceModel describes the resource data model.
type Lsnip6profileResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Natprefix types.String `tfsdk:"natprefix"`
	Network6  types.String `tfsdk:"network6"`
	Type      types.String `tfsdk:"type"`
}

func (r *Lsnip6profileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnip6profile resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the LSN ip6 profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN ip6 profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn ip6 profile1\" or 'lsn ip6 profile1').",
			},
			"natprefix": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv6 address(es) of the LSN subscriber(s) or subscriber network(s) on whose traffic you want the Citrix ADC to perform Large Scale NAT.",
			},
			"network6": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv6 address of the Citrix ADC AFTR device",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv6 translation type for which to set the LSN IP6 profile parameters.",
			},
		},
	}
}

func lsnip6profileGetThePayloadFromtheConfig(ctx context.Context, data *Lsnip6profileResourceModel) lsn.Lsnip6profile {
	tflog.Debug(ctx, "In lsnip6profileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnip6profile := lsn.Lsnip6profile{}
	if !data.Name.IsNull() {
		lsnip6profile.Name = data.Name.ValueString()
	}
	if !data.Natprefix.IsNull() {
		lsnip6profile.Natprefix = data.Natprefix.ValueString()
	}
	if !data.Network6.IsNull() {
		lsnip6profile.Network6 = data.Network6.ValueString()
	}
	if !data.Type.IsNull() {
		lsnip6profile.Type = data.Type.ValueString()
	}

	return lsnip6profile
}

func lsnip6profileSetAttrFromGet(ctx context.Context, data *Lsnip6profileResourceModel, getResponseData map[string]interface{}) *Lsnip6profileResourceModel {
	tflog.Debug(ctx, "In lsnip6profileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["natprefix"]; ok && val != nil {
		data.Natprefix = types.StringValue(val.(string))
	} else {
		data.Natprefix = types.StringNull()
	}
	if val, ok := getResponseData["network6"]; ok && val != nil {
		data.Network6 = types.StringValue(val.(string))
	} else {
		data.Network6 = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
