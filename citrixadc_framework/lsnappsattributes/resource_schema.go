package lsnappsattributes

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
			// SDK v2: Required + ForceNew -> Required + RequiresReplace.
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the LSN Application Port ATTRIBUTES. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN application profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn application profile1\" or 'lsn application profile1').",
			},
			// SDK v2: Optional + ForceNew (no Computed). Computed added so a value
			// returned by the ADC when the attribute is not configured does not cause an
			// inconsistent-result-after-apply error. RequiresReplaceIfConfigured reproduces
			// ForceNew only when the user actually configured a change.
			"port": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "This is used for Displaying Port/Port range in CLI/Nitro.Lowport, Highport values are populated and used for displaying.Port numbers or range of port numbers to match against the destination port of the incoming packet from a subscriber. When the destination port is matched, the LSN application profile is applied for the LSN session. Separate a range of ports with a hyphen. For example, 40-90.",
			},
			// SDK v2: Optional + Computed, no Default (value read from ADC). The only
			// in-place updateable attribute.
			"sessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout, in seconds, for an idle LSN session. If an LSN session is idle for a time that exceeds this value, the Citrix ADC removes the session.This timeout does not apply for a TCP LSN session when a FIN or RST message is received from either of the endpoints.",
			},
			// SDK v2: Required + ForceNew -> Required + RequiresReplace.
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

// lsnappsattributesGetThePayloadFromthePlan builds the full create payload.
func lsnappsattributesGetThePayloadFromthePlan(ctx context.Context, data *LsnappsattributesResourceModel) lsn.Lsnappsattributes {
	tflog.Debug(ctx, "In lsnappsattributesGetThePayloadFromthePlan Function")

	lsnappsattributes := lsn.Lsnappsattributes{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		lsnappsattributes.Name = data.Name.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		lsnappsattributes.Port = data.Port.ValueString()
	}
	if !data.Sessiontimeout.IsNull() && !data.Sessiontimeout.IsUnknown() {
		lsnappsattributes.Sessiontimeout = utils.IntPtr(int(data.Sessiontimeout.ValueInt64()))
	}
	if !data.Transportprotocol.IsNull() && !data.Transportprotocol.IsUnknown() {
		lsnappsattributes.Transportprotocol = data.Transportprotocol.ValueString()
	}

	return lsnappsattributes
}

// lsnappsattributesGetTheUpdatablePayloadFromThePlan builds the minimal update
// payload. Matches SDK v2: only the name (identity) and the single updateable
// attribute (sessiontimeout) are sent; port/transportprotocol are ForceNew and
// never reach Update.
func lsnappsattributesGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *LsnappsattributesResourceModel) lsn.Lsnappsattributes {
	tflog.Debug(ctx, "In lsnappsattributesGetTheUpdatablePayloadFromThePlan Function")

	lsnappsattributes := lsn.Lsnappsattributes{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		lsnappsattributes.Name = data.Name.ValueString()
	}
	if !data.Sessiontimeout.IsNull() && !data.Sessiontimeout.IsUnknown() {
		lsnappsattributes.Sessiontimeout = utils.IntPtr(int(data.Sessiontimeout.ValueInt64()))
	}

	return lsnappsattributes
}

func lsnappsattributesSetAttrFromGet(ctx context.Context, data *LsnappsattributesResourceModel, getResponseData map[string]interface{}) *LsnappsattributesResourceModel {
	tflog.Debug(ctx, "In lsnappsattributesSetAttrFromGet Function")

	// Convert API response to model. Guard the else-branch so a value the ADC omits
	// from GET (omit-on-default) only nulls an Unknown (Computed) value and never
	// clobbers a known configured value.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["port"]; ok && val != nil {
		data.Port = types.StringValue(val.(string))
	} else if data.Port.IsUnknown() {
		data.Port = types.StringNull()
	}
	if val, ok := getResponseData["sessiontimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sessiontimeout = types.Int64Value(intVal)
		}
	} else if data.Sessiontimeout.IsUnknown() {
		data.Sessiontimeout = types.Int64Null()
	}
	if val, ok := getResponseData["transportprotocol"]; ok && val != nil {
		data.Transportprotocol = types.StringValue(val.(string))
	} else if data.Transportprotocol.IsUnknown() {
		data.Transportprotocol = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
