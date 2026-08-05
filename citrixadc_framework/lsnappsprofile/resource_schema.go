package lsnappsprofile

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

// LsnappsprofileResourceModel describes the resource data model.
type LsnappsprofileResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Appsprofilename   types.String `tfsdk:"appsprofilename"`
	Filtering         types.String `tfsdk:"filtering"`
	Ippooling         types.String `tfsdk:"ippooling"`
	L2info            types.String `tfsdk:"l2info"`
	Mapping           types.String `tfsdk:"mapping"`
	Tcpproxy          types.String `tfsdk:"tcpproxy"`
	Td                types.Int64  `tfsdk:"td"`
	Transportprotocol types.String `tfsdk:"transportprotocol"`
}

func (r *LsnappsprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnappsprofile resource.",
			},
			"appsprofilename": schema.StringAttribute{
				Required: true,
				// SDK v2 ForceNew: true
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the LSN application profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN application profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn application profile1\" or 'lsn application profile1').",
			},
			"filtering": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of filter to apply to packets originating from external hosts.\n\nConsider an example of an LSN mapping that includes the mapping of subscriber IP:port (X:x), NAT IP:port (N:n), and external host IP:port (Y:y).\n\nAvailable options function as follows:\n* ENDPOINT INDEPENDENT - Filters out only packets not destined to the subscriber IP address and port X:x, regardless of the external host IP address and port source (Z:z).  The Citrix ADC forwards any packets destined to X:x.  In other words, sending packets from the subscriber to any external IP address is sufficient to allow packets from any external hosts to the subscriber.\n\n* ADDRESS DEPENDENT - Filters out packets not destined to subscriber IP address and port X:x.  In addition, the ADC filters out packets from Y:y destined for the subscriber (X:x) if the client has not previously sent packets to Y:anyport (external port independent). In other words, receiving packets from a specific external host requires that the subscriber first send packets to that specific external host's IP address.\n\n* ADDRESS PORT DEPENDENT (the default) - Filters out  packets not destined to subscriber IP address and port (X:x).  In addition, the Citrix ADC filters out packets from Y:y destined for the subscriber (X:x) if the subscriber has not previously sent packets to Y:y.  In other words, receiving packets from a specific external host requires that the subscriber first send packets first to that external IP address and port.",
			},
			"ippooling": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "NAT IP address allocation options for sessions associated with the same subscriber.\n\nAvailable options function as follows:\n* Paired - The Citrix ADC allocates the same NAT IP address for all sessions associated with the same subscriber. When all the ports of a NAT IP address are used in LSN sessions (for same or multiple subscribers), the Citrix ADC ADC drops any new connection from the subscriber.\n* Random - The Citrix ADC allocates random NAT IP addresses, from the pool, for different sessions associated with the same subscriber.\n\nThis parameter is applicable to dynamic NAT allocation only.",
			},
			"l2info": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable l2info by creating natpcbs for LSN, which enables the Citrix ADC to use L2CONN/MBF with LSN.",
			},
			"mapping": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of LSN mapping to apply to subsequent packets originating from the same subscriber IP address and port.\n\nConsider an example of an LSN mapping that includes the mapping of the subscriber IP:port (X:x), NAT IP:port (N:n), and external host IP:port (Y:y).\n\nAvailable options function as follows: \n\n* ENDPOINT-INDEPENDENT - Reuse the LSN mapping for subsequent packets sent from the same subscriber IP address and port (X:x) to any external IP address and port. \n\n* ADDRESS-DEPENDENT - Reuse the LSN mapping for subsequent packets sent from the same subscriber IP address and port (X:x) to the same external IP address (Y), regardless of the external port.\n\n* ADDRESS-PORT-DEPENDENT - Reuse the LSN mapping for subsequent packets sent from the same internal IP address and port (X:x) to the same external IP address and port (Y:y) while the mapping is still active.",
			},
			"tcpproxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable TCP proxy, which enables the Citrix ADC to optimize the  TCP traffic by using Layer 4 features.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the traffic domain through which the Citrix ADC sends the outbound traffic after performing LSN. \n\nIf you do not specify an ID, the ADC sends the outbound traffic through the default traffic domain, which has an ID of 0.",
			},
			"transportprotocol": schema.StringAttribute{
				Required: true,
				// SDK v2 ForceNew: true
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the protocol for which the parameters of this LSN application profile applies.",
			},
		},
	}
}

// lsnappsprofileGetThePayloadFromthePlan builds the full add/create payload.
func lsnappsprofileGetThePayloadFromthePlan(ctx context.Context, data *LsnappsprofileResourceModel) lsn.Lsnappsprofile {
	tflog.Debug(ctx, "In lsnappsprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	lsnappsprofile := lsn.Lsnappsprofile{}
	if !data.Appsprofilename.IsNull() && !data.Appsprofilename.IsUnknown() {
		lsnappsprofile.Appsprofilename = data.Appsprofilename.ValueString()
	}
	if !data.Filtering.IsNull() && !data.Filtering.IsUnknown() {
		lsnappsprofile.Filtering = data.Filtering.ValueString()
	}
	if !data.Ippooling.IsNull() && !data.Ippooling.IsUnknown() {
		lsnappsprofile.Ippooling = data.Ippooling.ValueString()
	}
	if !data.L2info.IsNull() && !data.L2info.IsUnknown() {
		lsnappsprofile.L2info = data.L2info.ValueString()
	}
	if !data.Mapping.IsNull() && !data.Mapping.IsUnknown() {
		lsnappsprofile.Mapping = data.Mapping.ValueString()
	}
	if !data.Tcpproxy.IsNull() && !data.Tcpproxy.IsUnknown() {
		lsnappsprofile.Tcpproxy = data.Tcpproxy.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		lsnappsprofile.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Transportprotocol.IsNull() && !data.Transportprotocol.IsUnknown() {
		lsnappsprofile.Transportprotocol = data.Transportprotocol.ValueString()
	}

	return lsnappsprofile
}

// lsnappsprofileGetTheUpdatablePayloadFromThePlan builds the update payload.
// It intentionally excludes the ForceNew attribute transportprotocol (SDK v2
// parity - the update PUT only carries appsprofilename plus updateable fields).
func lsnappsprofileGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *LsnappsprofileResourceModel) lsn.Lsnappsprofile {
	tflog.Debug(ctx, "In lsnappsprofileGetTheUpdatablePayloadFromThePlan Function")

	lsnappsprofile := lsn.Lsnappsprofile{}
	if !data.Appsprofilename.IsNull() && !data.Appsprofilename.IsUnknown() {
		lsnappsprofile.Appsprofilename = data.Appsprofilename.ValueString()
	}
	if !data.Filtering.IsNull() && !data.Filtering.IsUnknown() {
		lsnappsprofile.Filtering = data.Filtering.ValueString()
	}
	if !data.Ippooling.IsNull() && !data.Ippooling.IsUnknown() {
		lsnappsprofile.Ippooling = data.Ippooling.ValueString()
	}
	if !data.L2info.IsNull() && !data.L2info.IsUnknown() {
		lsnappsprofile.L2info = data.L2info.ValueString()
	}
	if !data.Mapping.IsNull() && !data.Mapping.IsUnknown() {
		lsnappsprofile.Mapping = data.Mapping.ValueString()
	}
	if !data.Tcpproxy.IsNull() && !data.Tcpproxy.IsUnknown() {
		lsnappsprofile.Tcpproxy = data.Tcpproxy.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		lsnappsprofile.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	// transportprotocol is ForceNew - excluded from the update payload

	return lsnappsprofile
}

func lsnappsprofileSetAttrFromGet(ctx context.Context, data *LsnappsprofileResourceModel, getResponseData map[string]interface{}) *LsnappsprofileResourceModel {
	tflog.Debug(ctx, "In lsnappsprofileSetAttrFromGet Function")

	// Convert API response to model. Guard the else-branches so that a value
	// NITRO omits from GET (omit-on-default) is only nulled when unknown - never
	// clobbering a known configured value.
	if val, ok := getResponseData["appsprofilename"]; ok && val != nil {
		data.Appsprofilename = types.StringValue(val.(string))
	} else if data.Appsprofilename.IsUnknown() {
		data.Appsprofilename = types.StringNull()
	}
	if val, ok := getResponseData["filtering"]; ok && val != nil {
		data.Filtering = types.StringValue(val.(string))
	} else if data.Filtering.IsUnknown() {
		data.Filtering = types.StringNull()
	}
	if val, ok := getResponseData["ippooling"]; ok && val != nil {
		data.Ippooling = types.StringValue(val.(string))
	} else if data.Ippooling.IsUnknown() {
		data.Ippooling = types.StringNull()
	}
	if val, ok := getResponseData["l2info"]; ok && val != nil {
		data.L2info = types.StringValue(val.(string))
	} else if data.L2info.IsUnknown() {
		data.L2info = types.StringNull()
	}
	if val, ok := getResponseData["mapping"]; ok && val != nil {
		data.Mapping = types.StringValue(val.(string))
	} else if data.Mapping.IsUnknown() {
		data.Mapping = types.StringNull()
	}
	if val, ok := getResponseData["tcpproxy"]; ok && val != nil {
		data.Tcpproxy = types.StringValue(val.(string))
	} else if data.Tcpproxy.IsUnknown() {
		data.Tcpproxy = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["transportprotocol"]; ok && val != nil {
		data.Transportprotocol = types.StringValue(val.(string))
	} else if data.Transportprotocol.IsUnknown() {
		data.Transportprotocol = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Appsprofilename.ValueString())

	return data
}
