package lsnappsattributes

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LsnappsattributesResourceModel describes the resource data model.
type LsnappsattributesResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	Port              types.String `tfsdk:"port"`
	Sessiontimeout    types.Int64  `tfsdk:"sessiontimeout"`
	Transportprotocol types.String `tfsdk:"transportprotocol"`
}

func (r *LsnappsattributesResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnappsattributes resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN Application Port ATTRIBUTES. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN application profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn application profile1\" or 'lsn application profile1').",
			},
			"port": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "This is used for Displaying Port/Port range in CLI/Nitro.Lowport, Highport values are populated and used for displaying.Port numbers or range of port numbers to match against the destination port of the incoming packet from a subscriber. When the destination port is matched, the LSN application profile is applied for the LSN session. Separate a range of ports with a hyphen. For example, 40-90.",
			},
			"sessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(30),
				Description: "Timeout, in seconds, for an idle LSN session. If an LSN session is idle for a time that exceeds this value, the Citrix ADC removes the session.This timeout does not apply for a TCP LSN session when a FIN or RST message is received from either of the endpoints.",
			},
			"transportprotocol": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the protocol(TCP,UDP) for which the parameters of this LSN application port ATTRIBUTES applies",
			},
		},
	}
}

func lsnappsattributesGetThePayloadFromtheConfig(ctx context.Context, data *LsnappsattributesResourceModel) lsn.Lsnappsattributes {
	tflog.Debug(ctx, "In lsnappsattributesGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnappsattributes := lsn.Lsnappsattributes{}
	if !data.Name.IsNull() {
		lsnappsattributes.Name = data.Name.ValueString()
	}
	if !data.Port.IsNull() {
		lsnappsattributes.Port = data.Port.ValueString()
	}
	if !data.Sessiontimeout.IsNull() {
		lsnappsattributes.Sessiontimeout = utils.IntPtr(int(data.Sessiontimeout.ValueInt64()))
	}
	if !data.Transportprotocol.IsNull() {
		lsnappsattributes.Transportprotocol = data.Transportprotocol.ValueString()
	}

	return lsnappsattributes
}

func lsnappsattributesSetAttrFromGet(ctx context.Context, data *LsnappsattributesResourceModel, getResponseData map[string]interface{}) *LsnappsattributesResourceModel {
	tflog.Debug(ctx, "In lsnappsattributesSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		data.Port = types.StringValue(val.(string))
	} else {
		data.Port = types.StringNull()
	}
	if val, ok := getResponseData["sessiontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sessiontimeout = types.Int64Value(intVal)
		}
	} else {
		data.Sessiontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["transportprotocol"]; ok && val != nil {
		data.Transportprotocol = types.StringValue(val.(string))
	} else {
		data.Transportprotocol = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
