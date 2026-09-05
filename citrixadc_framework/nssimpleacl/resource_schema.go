package nssimpleacl

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NssimpleaclResourceModel describes the resource data model.
type NssimpleaclResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Aclaction   types.String `tfsdk:"aclaction"`
	Aclname     types.String `tfsdk:"aclname"`
	Destport    types.Int64  `tfsdk:"destport"`
	Estsessions types.Bool   `tfsdk:"estsessions"`
	Protocol    types.String `tfsdk:"protocol"`
	Srcip       types.String `tfsdk:"srcip"`
	Td          types.Int64  `tfsdk:"td"`
	Ttl         types.Int64  `tfsdk:"ttl"`
}

func (r *NssimpleaclResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nssimpleacl resource.",
			},
			"aclaction": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Drop incoming IPv4 packets that match the simple ACL rule.",
			},
			"aclname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the simple ACL rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the simple ACL rule is created.",
			},
			"destport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// GH #1436
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Port number to match against the destination port number of an incoming IPv4 packet.\n\nDestPort is mandatory while setting Protocol. Omitting the port number and protocol creates an all-ports  and all protocols simple ACL rule, which matches any port and any protocol. In that case, you cannot create another simple ACL rule specifying a specific port and the same source IPv4 address.",
			},
			"estsessions": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					// GH #1436
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "0",
			},
			"protocol": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Protocol to match against the protocol of an incoming IPv4 packet. You must set this parameter if you have set the Destination Port parameter.",
			},
			"srcip": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address to match against the source IP address of an incoming IPv4 packet.",
			},
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// GH #1436
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"ttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					// GH #1436
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Number of seconds, in multiples of four, after which the simple ACL rule expires. If you do not want the simple ACL rule to expire, do not specify a TTL value.",
			},
		},
	}
}

func nssimpleaclGetThePayloadFromtheConfig(ctx context.Context, data *NssimpleaclResourceModel) ns.Nssimpleacl {
	tflog.Debug(ctx, "In nssimpleaclGetThePayloadFromtheConfig Function")

	// Create API request body from the model.
	// Only send values the user actually configured: for attributes read from the
	// plan, unconfigured Optional+Computed attrs are Unknown (not Null), so guard on
	// both to avoid pushing spurious zero-values to the ADC (matches the SDK v2
	// GetRawConfig-gated behavior).
	nssimpleacl := ns.Nssimpleacl{}
	if !data.Aclaction.IsNull() && !data.Aclaction.IsUnknown() {
		nssimpleacl.Aclaction = data.Aclaction.ValueString()
	}
	if !data.Aclname.IsNull() && !data.Aclname.IsUnknown() {
		nssimpleacl.Aclname = data.Aclname.ValueString()
	}
	if !data.Destport.IsNull() && !data.Destport.IsUnknown() {
		nssimpleacl.Destport = utils.IntPtr(int(data.Destport.ValueInt64()))
	}
	if !data.Estsessions.IsNull() && !data.Estsessions.IsUnknown() {
		nssimpleacl.Estsessions = data.Estsessions.ValueBool()
	}
	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		nssimpleacl.Protocol = data.Protocol.ValueString()
	}
	if !data.Srcip.IsNull() && !data.Srcip.IsUnknown() {
		nssimpleacl.Srcip = data.Srcip.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		nssimpleacl.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		nssimpleacl.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return nssimpleacl
}

func nssimpleaclSetAttrFromGet(ctx context.Context, data *NssimpleaclResourceModel, getResponseData map[string]interface{}) *NssimpleaclResourceModel {
	tflog.Debug(ctx, "In nssimpleaclSetAttrFromGet Function")

	// Convert API response to model.
	//
	// Hybrid setter: when NITRO returns the field, adopt it (enables drift
	// detection, matching SDK v2's d.Set). When NITRO OMITS the field (common for
	// zero/false values with json omitempty), do NOT clobber a value the user
	// configured — only null it when the current value is Unknown (omit-on-default
	// trap guard). This serves both the resource (preserves configured values) and
	// the datasource (its non-key fields arrive Null, so absent GET fields stay
	// Null while present ones are copied).
	if val, ok := getResponseData["aclaction"]; ok && val != nil {
		data.Aclaction = types.StringValue(val.(string))
	} else if data.Aclaction.IsUnknown() {
		data.Aclaction = types.StringNull()
	}
	if val, ok := getResponseData["aclname"]; ok && val != nil {
		data.Aclname = types.StringValue(val.(string))
	} else if data.Aclname.IsUnknown() {
		data.Aclname = types.StringNull()
	}
	if val, ok := getResponseData["destport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Destport = types.Int64Value(intVal)
		}
	} else if data.Destport.IsUnknown() {
		data.Destport = types.Int64Null()
	}
	if val, ok := getResponseData["estsessions"]; ok && val != nil {
		data.Estsessions = types.BoolValue(val.(bool))
	} else if data.Estsessions.IsUnknown() {
		data.Estsessions = types.BoolNull()
	}
	if val, ok := getResponseData["protocol"]; ok && val != nil {
		data.Protocol = types.StringValue(val.(string))
	} else if data.Protocol.IsUnknown() {
		data.Protocol = types.StringNull()
	}
	if val, ok := getResponseData["srcip"]; ok && val != nil {
		data.Srcip = types.StringValue(val.(string))
	} else if data.Srcip.IsUnknown() {
		data.Srcip = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}

	// ttl: the ADC returns a DECREMENTING time-to-live (e.g. a configured 600 comes
	// back as 599 on the next read), so a configured value must be preserved to
	// avoid perpetual diffs / "inconsistent result after apply". This mirrors the
	// SDK v2 resource, which deliberately re-set ttl from the user config, not the
	// API. Only resolve ttl from the GET response when the current value is
	// Unknown/Null (unconfigured Computed, or a datasource read).
	if data.Ttl.IsUnknown() || data.Ttl.IsNull() {
		if val, ok := getResponseData["ttl"]; ok && val != nil {
			if intVal, err := utils.ConvertToInt64(val); err == nil {
				data.Ttl = types.Int64Value(intVal)
			} else {
				data.Ttl = types.Int64Null()
			}
		} else {
			data.Ttl = types.Int64Null()
		}
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Aclname.ValueString())

	return data
}
