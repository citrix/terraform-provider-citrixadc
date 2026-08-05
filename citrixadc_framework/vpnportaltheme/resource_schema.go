package vpnportaltheme

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnportalthemeResourceModel describes the resource data model.
type VpnportalthemeResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Basetheme types.String `tfsdk:"basetheme"`
	Name      types.String `tfsdk:"name"`
}

func (r *VpnportalthemeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnportaltheme resource.",
			},
			"basetheme": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the uitheme",
			},
		},
	}
}

func vpnportalthemeGetThePayloadFromtheConfig(ctx context.Context, data *VpnportalthemeResourceModel) vpn.Vpnportaltheme {
	tflog.Debug(ctx, "In vpnportalthemeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnportaltheme := vpn.Vpnportaltheme{}
	if !data.Basetheme.IsNull() && !data.Basetheme.IsUnknown() {
		vpnportaltheme.Basetheme = data.Basetheme.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnportaltheme.Name = data.Name.ValueString()
	}

	return vpnportaltheme
}

func vpnportalthemeSetAttrFromGet(ctx context.Context, data *VpnportalthemeResourceModel, getResponseData map[string]interface{}) *VpnportalthemeResourceModel {
	tflog.Debug(ctx, "In vpnportalthemeSetAttrFromGet Function")

	// Convert API response to model.
	// Guard the else-branches so we only null a value when it is unknown; never
	// clobber a known configured value that NITRO may omit from GET (omit-on-default trap).
	if val, ok := getResponseData["basetheme"]; ok && val != nil {
		data.Basetheme = types.StringValue(val.(string))
	} else if data.Basetheme.IsUnknown() {
		data.Basetheme = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - plain value (name)
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
