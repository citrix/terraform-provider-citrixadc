package rnat

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// RnatResourceModel describes the resource data model.
type RnatResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Aclname          types.String `tfsdk:"aclname"`
	Connfailover     types.String `tfsdk:"connfailover"`
	Name             types.String `tfsdk:"name"`
	Natip            types.String `tfsdk:"natip"`
	Netmask          types.String `tfsdk:"netmask"`
	Network          types.String `tfsdk:"network"`
	Newname          types.String `tfsdk:"newname"`
	Ownergroup       types.String `tfsdk:"ownergroup"`
	Redirectport     types.Int64  `tfsdk:"redirectport"`
	Srcippersistency types.String `tfsdk:"srcippersistency"`
	Td               types.Int64  `tfsdk:"td"`
	Useproxyport     types.String `tfsdk:"useproxyport"`
}

func (r *RnatResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rnat resource.",
			},
			"aclname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An extended ACL defined for the RNAT entry.",
			},
			"connfailover": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Synchronize all connection-related information for the RNAT sessions with the secondary ADC in a high availability (HA) pair.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for the RNAT4 rule. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. Choose a name that helps identify the RNAT4 rule.",
			},
			"natip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The subnet mask for the network address.",
			},
			"network": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The network address defined for the RNAT entry.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the RNAT4 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain       only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DEFAULT_NG"),
				Description: "The owner node group in a Cluster for this rnat rule.",
			},
			"redirectport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number to which the IPv4 packets are redirected. Applicable to TCP and UDP protocols.",
			},
			"srcippersistency": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enables the Citrix ADC to use the same NAT IP address for all RNAT sessions initiated from a particular server.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable source port proxying, which enables the Citrix ADC to use the RNAT ips using proxied source port.",
			},
		},
	}
}

func rnatGetThePayloadFromtheConfig(ctx context.Context, data *RnatResourceModel) network.Rnat {
	tflog.Debug(ctx, "In rnatGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	rnat := network.Rnat{}
	if !data.Aclname.IsNull() {
		rnat.Aclname = data.Aclname.ValueString()
	}
	if !data.Connfailover.IsNull() {
		rnat.Connfailover = data.Connfailover.ValueString()
	}
	if !data.Name.IsNull() {
		rnat.Name = data.Name.ValueString()
	}
	if !data.Natip.IsNull() {
		rnat.Natip = data.Natip.ValueString()
	}
	if !data.Netmask.IsNull() {
		rnat.Netmask = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() {
		rnat.Network = data.Network.ValueString()
	}
	if !data.Newname.IsNull() {
		rnat.Newname = data.Newname.ValueString()
	}
	if !data.Ownergroup.IsNull() {
		rnat.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Redirectport.IsNull() {
		rnat.Redirectport = utils.IntPtr(int(data.Redirectport.ValueInt64()))
	}
	if !data.Srcippersistency.IsNull() {
		rnat.Srcippersistency = data.Srcippersistency.ValueString()
	}
	if !data.Td.IsNull() {
		rnat.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Useproxyport.IsNull() {
		rnat.Useproxyport = data.Useproxyport.ValueString()
	}

	return rnat
}

func rnatSetAttrFromGet(ctx context.Context, data *RnatResourceModel, getResponseData map[string]interface{}) *RnatResourceModel {
	tflog.Debug(ctx, "In rnatSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["aclname"]; ok && val != nil {
		data.Aclname = types.StringValue(val.(string))
	} else {
		data.Aclname = types.StringNull()
	}
	if val, ok := getResponseData["connfailover"]; ok && val != nil {
		data.Connfailover = types.StringValue(val.(string))
	} else {
		data.Connfailover = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["natip"]; ok && val != nil {
		data.Natip = types.StringValue(val.(string))
	} else {
		data.Natip = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["network"]; ok && val != nil {
		data.Network = types.StringValue(val.(string))
	} else {
		data.Network = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["redirectport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Redirectport = types.Int64Value(intVal)
		}
	} else {
		data.Redirectport = types.Int64Null()
	}
	if val, ok := getResponseData["srcippersistency"]; ok && val != nil {
		data.Srcippersistency = types.StringValue(val.(string))
	} else {
		data.Srcippersistency = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["useproxyport"]; ok && val != nil {
		data.Useproxyport = types.StringValue(val.(string))
	} else {
		data.Useproxyport = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
