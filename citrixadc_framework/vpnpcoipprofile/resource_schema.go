package vpnpcoipprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnpcoipprofileResourceModel describes the resource data model.
type VpnpcoipprofileResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Conserverurl       types.String `tfsdk:"conserverurl"`
	Icvverification    types.String `tfsdk:"icvverification"`
	Name               types.String `tfsdk:"name"`
	Sessionidletimeout types.Int64  `tfsdk:"sessionidletimeout"`
}

func (r *VpnpcoipprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnpcoipprofile resource.",
			},
			"conserverurl": schema.StringAttribute{
				Required:    true,
				Description: "Connection server URL",
			},
			"icvverification": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ICV verification for PCOIP transport packets.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "name of PCoIP profile",
			},
			"sessionidletimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "PCOIP Idle Session timeout",
			},
		},
	}
}

func vpnpcoipprofileGetThePayloadFromthePlan(ctx context.Context, data *VpnpcoipprofileResourceModel) vpn.Vpnpcoipprofile {
	tflog.Debug(ctx, "In vpnpcoipprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnpcoipprofile := vpn.Vpnpcoipprofile{}
	if !data.Conserverurl.IsNull() && !data.Conserverurl.IsUnknown() {
		vpnpcoipprofile.Conserverurl = data.Conserverurl.ValueString()
	}
	if !data.Icvverification.IsNull() && !data.Icvverification.IsUnknown() {
		vpnpcoipprofile.Icvverification = data.Icvverification.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnpcoipprofile.Name = data.Name.ValueString()
	}
	if !data.Sessionidletimeout.IsNull() && !data.Sessionidletimeout.IsUnknown() {
		vpnpcoipprofile.Sessionidletimeout = utils.IntPtr(int(data.Sessionidletimeout.ValueInt64()))
	}

	return vpnpcoipprofile
}

func vpnpcoipprofileSetAttrFromGet(ctx context.Context, data *VpnpcoipprofileResourceModel, getResponseData map[string]interface{}) *VpnpcoipprofileResourceModel {
	tflog.Debug(ctx, "In vpnpcoipprofileSetAttrFromGet Function")

	// Convert API response to model.
	// Guard else-branches with IsUnknown() so a configured value that NITRO omits
	// from GET (omit-on-default trap) is preserved rather than nulled.
	if val, ok := getResponseData["conserverurl"]; ok && val != nil {
		data.Conserverurl = types.StringValue(val.(string))
	} else if data.Conserverurl.IsUnknown() {
		data.Conserverurl = types.StringNull()
	}
	if val, ok := getResponseData["icvverification"]; ok && val != nil {
		data.Icvverification = types.StringValue(val.(string))
	} else if data.Icvverification.IsUnknown() {
		data.Icvverification = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["sessionidletimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sessionidletimeout = types.Int64Value(intVal)
		} else if data.Sessionidletimeout.IsUnknown() {
			data.Sessionidletimeout = types.Int64Null()
		}
	} else if data.Sessionidletimeout.IsUnknown() {
		data.Sessionidletimeout = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
