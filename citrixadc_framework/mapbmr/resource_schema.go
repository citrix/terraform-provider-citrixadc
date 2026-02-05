package mapbmr

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// MapbmrResourceModel describes the resource data model.
type MapbmrResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Eabitlength    types.Int64  `tfsdk:"eabitlength"`
	Name           types.String `tfsdk:"name"`
	Psidlength     types.Int64  `tfsdk:"psidlength"`
	Psidoffset     types.Int64  `tfsdk:"psidoffset"`
	Ruleipv6prefix types.String `tfsdk:"ruleipv6prefix"`
}

func (r *MapbmrResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the mapbmr resource.",
			},
			"eabitlength": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(16),
				Description: "The Embedded Address (EA) bit field encodes the CE-specific IPv4 address and port information.  The EA bit field, which is unique for a\n			          given Rule IPv6 prefix.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the Basic Mapping Rule. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the  MAP Basic Mapping Rule is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"add network MapBmr bmr1 -natprefix 2005::/64 -EAbitLength 16 -psidoffset 6 -portsharingratio 8\" ).\n			The Basic Mapping Rule information allows a MAP BR to determine source IPv4 address from the IPv6 packet sent from MAP CE device.\n			Also it allows to determine destination IPv6 address of MAP CE before sending packets to MAP CE",
			},
			"psidlength": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(8),
				Description: "Length of Port Set IdentifierPort Set Identifier(PSID) in Embedded Address (EA) bits",
			},
			"psidoffset": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(6),
				Description: "Start bit position  of Port Set Identifier(PSID) value in Embedded Address (EA) bits.",
			},
			"ruleipv6prefix": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IPv6 prefix of Customer Edge(CE) device.MAP-T CE will send ipv6 packets with this ipv6 prefix as source ipv6 address prefix",
			},
		},
	}
}

func mapbmrGetThePayloadFromtheConfig(ctx context.Context, data *MapbmrResourceModel) network.Mapbmr {
	tflog.Debug(ctx, "In mapbmrGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	mapbmr := network.Mapbmr{}
	if !data.Eabitlength.IsNull() {
		mapbmr.Eabitlength = utils.IntPtr(int(data.Eabitlength.ValueInt64()))
	}
	if !data.Name.IsNull() {
		mapbmr.Name = data.Name.ValueString()
	}
	if !data.Psidlength.IsNull() {
		mapbmr.Psidlength = utils.IntPtr(int(data.Psidlength.ValueInt64()))
	}
	if !data.Psidoffset.IsNull() {
		mapbmr.Psidoffset = utils.IntPtr(int(data.Psidoffset.ValueInt64()))
	}
	if !data.Ruleipv6prefix.IsNull() {
		mapbmr.Ruleipv6prefix = data.Ruleipv6prefix.ValueString()
	}

	return mapbmr
}

func mapbmrSetAttrFromGet(ctx context.Context, data *MapbmrResourceModel, getResponseData map[string]interface{}) *MapbmrResourceModel {
	tflog.Debug(ctx, "In mapbmrSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["eabitlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Eabitlength = types.Int64Value(intVal)
		}
	} else {
		data.Eabitlength = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["psidlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Psidlength = types.Int64Value(intVal)
		}
	} else {
		data.Psidlength = types.Int64Null()
	}
	if val, ok := getResponseData["psidoffset"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Psidoffset = types.Int64Value(intVal)
		}
	} else {
		data.Psidoffset = types.Int64Null()
	}
	if val, ok := getResponseData["ruleipv6prefix"]; ok && val != nil {
		data.Ruleipv6prefix = types.StringValue(val.(string))
	} else {
		data.Ruleipv6prefix = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
