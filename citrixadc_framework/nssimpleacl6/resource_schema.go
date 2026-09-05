package nssimpleacl6

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

// Nssimpleacl6ResourceModel describes the resource data model.
type Nssimpleacl6ResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Aclaction   types.String `tfsdk:"aclaction"`
	Aclname     types.String `tfsdk:"aclname"`
	Destport    types.Int64  `tfsdk:"destport"`
	Estsessions types.Bool   `tfsdk:"estsessions"`
	Protocol    types.String `tfsdk:"protocol"`
	Srcipv6     types.String `tfsdk:"srcipv6"`
	Td          types.Int64  `tfsdk:"td"`
	Ttl         types.Int64  `tfsdk:"ttl"`
}

func (r *Nssimpleacl6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nssimpleacl6 resource.",
			},
			// SDK v2: Required + ForceNew
			"aclaction": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Drop incoming IPv6 packets that match the simple ACL6 rule.",
			},
			// SDK v2: Required + ForceNew (primary key)
			"aclname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the simple ACL6 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the simple ACL6 rule is created.",
			},
			// SDK v2: Optional + Computed + ForceNew -> UseStateForUnknown + RequiresReplaceIfConfigured
			"destport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Port number to match against the destination port number of an incoming IPv6 packet.\n\nDestPort is mandatory while setting Protocol. Omitting the port number and protocol creates an all-ports  and all protocol simple ACL6 rule, which matches any port and any protocol. In that case, you cannot create another simple ACL6 rule specifying a specific port and the same source IPv6 address.",
			},
			// SDK v2: Optional + Computed + ForceNew -> UseStateForUnknown + RequiresReplaceIfConfigured
			"estsessions": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "0",
			},
			// SDK v2: Optional + Computed + ForceNew -> UseStateForUnknown + RequiresReplaceIfConfigured
			"protocol": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Protocol to match against the protocol of an incoming IPv6 packet. You must set this parameter if you set the Destination Port parameter.",
			},
			// SDK v2: Required + ForceNew
			"srcipv6": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address to match against the source IP address of an incoming IPv6 packet.",
			},
			// SDK v2: Optional + Computed + ForceNew -> UseStateForUnknown + RequiresReplaceIfConfigured
			"td": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			// SDK v2: Optional + Computed + ForceNew -> UseStateForUnknown + RequiresReplaceIfConfigured
			"ttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Number of seconds, in multiples of four, after which the simple ACL6 rule expires. If you do not want the simple ACL6 rule to expire, do not specify a TTL value.",
			},
		},
	}
}

func nssimpleacl6GetThePayloadFromthePlan(ctx context.Context, data *Nssimpleacl6ResourceModel) ns.Nssimpleacl6 {
	tflog.Debug(ctx, "In nssimpleacl6GetThePayloadFromthePlan Function")

	// Create API request body from the model. Only send explicitly-configured
	// values (mirrors SDK v2, which conditionally set destport/td/ttl and relied
	// on omitempty for the rest).
	nssimpleacl6 := ns.Nssimpleacl6{}
	if !data.Aclaction.IsNull() && !data.Aclaction.IsUnknown() {
		nssimpleacl6.Aclaction = data.Aclaction.ValueString()
	}
	if !data.Aclname.IsNull() && !data.Aclname.IsUnknown() {
		nssimpleacl6.Aclname = data.Aclname.ValueString()
	}
	if !data.Destport.IsNull() && !data.Destport.IsUnknown() {
		nssimpleacl6.Destport = utils.IntPtr(int(data.Destport.ValueInt64()))
	}
	if !data.Estsessions.IsNull() && !data.Estsessions.IsUnknown() {
		nssimpleacl6.Estsessions = data.Estsessions.ValueBool()
	}
	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		nssimpleacl6.Protocol = data.Protocol.ValueString()
	}
	if !data.Srcipv6.IsNull() && !data.Srcipv6.IsUnknown() {
		nssimpleacl6.Srcipv6 = data.Srcipv6.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		nssimpleacl6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Ttl.IsNull() && !data.Ttl.IsUnknown() {
		nssimpleacl6.Ttl = utils.IntPtr(int(data.Ttl.ValueInt64()))
	}

	return nssimpleacl6
}

// nssimpleacl6SetAttrFromGet populates the RESOURCE state from a GET response.
// It preserves the configured/state value of ttl because the NITRO API returns a
// continuously-decreasing time-to-live (e.g. 600 configured may read back as 599),
// which would otherwise trigger spurious diffs / inconsistent-result errors. This
// mirrors the SDK v2 behavior (setToInt("ttl", d, d.Get("ttl").(int))). The
// else-branches only null a value when it is still Unknown, so a configured
// 0/false value that NITRO omits from GET is never clobbered (omit-on-default guard).
func nssimpleacl6SetAttrFromGet(ctx context.Context, data *Nssimpleacl6ResourceModel, getResponseData map[string]interface{}) *Nssimpleacl6ResourceModel {
	tflog.Debug(ctx, "In nssimpleacl6SetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["srcipv6"]; ok && val != nil {
		data.Srcipv6 = types.StringValue(val.(string))
	} else if data.Srcipv6.IsUnknown() {
		data.Srcipv6 = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}
	// ttl is intentionally NOT read from the API: NITRO returns a decreasing
	// remaining-lifetime value which keeps changing. Preserve the configured/state
	// value. If it is still Unknown (e.g. bare import with no prior value), null it.
	if data.Ttl.IsUnknown() {
		data.Ttl = types.Int64Null()
	}

	// Set ID for the resource (single unique attribute -> plain aclname value).
	data.Id = types.StringValue(data.Aclname.ValueString())

	return data
}

// nssimpleacl6SetAttrFromGetForDatasource populates DATASOURCE state from a GET
// response. Unlike the resource setter, it copies every attribute (including ttl)
// straight from the API, since a datasource is a read-only view with no configured
// values to preserve.
func nssimpleacl6SetAttrFromGetForDatasource(ctx context.Context, data *Nssimpleacl6ResourceModel, getResponseData map[string]interface{}) *Nssimpleacl6ResourceModel {
	tflog.Debug(ctx, "In nssimpleacl6SetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["aclaction"]; ok && val != nil {
		data.Aclaction = types.StringValue(val.(string))
	} else {
		data.Aclaction = types.StringNull()
	}
	if val, ok := getResponseData["aclname"]; ok && val != nil {
		data.Aclname = types.StringValue(val.(string))
	} else {
		data.Aclname = types.StringNull()
	}
	if val, ok := getResponseData["destport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Destport = types.Int64Value(intVal)
		}
	} else {
		data.Destport = types.Int64Null()
	}
	if val, ok := getResponseData["estsessions"]; ok && val != nil {
		data.Estsessions = types.BoolValue(val.(bool))
	} else {
		data.Estsessions = types.BoolNull()
	}
	if val, ok := getResponseData["protocol"]; ok && val != nil {
		data.Protocol = types.StringValue(val.(string))
	} else {
		data.Protocol = types.StringNull()
	}
	if val, ok := getResponseData["srcipv6"]; ok && val != nil {
		data.Srcipv6 = types.StringValue(val.(string))
	} else {
		data.Srcipv6 = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["ttl"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ttl = types.Int64Value(intVal)
		}
	} else {
		data.Ttl = types.Int64Null()
	}

	// Set ID for the datasource (single unique attribute -> plain aclname value).
	data.Id = types.StringValue(data.Aclname.ValueString())

	return data
}
