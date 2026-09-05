package iptunnel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// IptunnelResourceModel describes the resource data model.
type IptunnelResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Destport         types.Int64  `tfsdk:"destport"`
	Grepayload       types.String `tfsdk:"grepayload"`
	Ipsecprofilename types.String `tfsdk:"ipsecprofilename"`
	Local            types.String `tfsdk:"local"`
	Name             types.String `tfsdk:"name"`
	Ownergroup       types.String `tfsdk:"ownergroup"`
	Protocol         types.String `tfsdk:"protocol"`
	Remote           types.String `tfsdk:"remote"`
	Remotesubnetmask types.String `tfsdk:"remotesubnetmask"`
	Tosinherit       types.String `tfsdk:"tosinherit"`
	Vlan             types.Int64  `tfsdk:"vlan"`
	Vlantagging      types.String `tfsdk:"vlantagging"`
	Vnid             types.Int64  `tfsdk:"vnid"`
}

func (r *IptunnelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the iptunnel resource.",
			},
			// Updateable in SDK v2 (not ForceNew).
			"destport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Description: "Specifies UDP destination port for Geneve packets. Default port is 6081.",
			},
			"grepayload": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "The payload GRE will carry",
			},
			"ipsecprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Name of IPSec profile to be associated.",
			},
			"local": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Type of Citrix ADC owned public IPv4 address, configured on the local Citrix ADC and used to set up the tunnel.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the IP tunnel. Leading character must be a number or letter. Other characters allowed, after the first character, are @ _ - . (period) : (colon) # and space ( ).",
			},
			"ownergroup": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "The owner node group in a Cluster for the iptunnel.",
			},
			"protocol": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Name of the protocol to be used on this tunnel.",
			},
			"remote": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Public IPv4 address, of the remote device, used to set up the tunnel. For this parameter, you can alternatively specify a network address.",
			},
			"remotesubnetmask": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Subnet mask of the remote IP address of the tunnel.",
			},
			// Updateable in SDK v2 (not ForceNew).
			"tosinherit": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Default behavior is to copy the ToS field of the internal IP Packet (Payload) to the outer IP packet (Transport packet). But the user can configure a new ToS field using this option.",
			},
			"vlan": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "The vlan for mulicast packets",
			},
			// Updateable in SDK v2 (not ForceNew).
			"vlantagging": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Option to select Vlan Tagging.",
			},
			"vnid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Virtual network identifier (VNID) is the value that identifies a specific virtual network in the data plane.",
			},
		},
	}
}

func iptunnelGetThePayloadFromtheConfig(ctx context.Context, data *IptunnelResourceModel) network.Iptunnel {
	tflog.Debug(ctx, "In iptunnelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	iptunnel := network.Iptunnel{}
	if !data.Destport.IsNull() && !data.Destport.IsUnknown() {
		iptunnel.Destport = utils.IntPtr(int(data.Destport.ValueInt64()))
	}
	if !data.Grepayload.IsNull() && !data.Grepayload.IsUnknown() {
		iptunnel.Grepayload = data.Grepayload.ValueString()
	}
	if !data.Ipsecprofilename.IsNull() && !data.Ipsecprofilename.IsUnknown() {
		iptunnel.Ipsecprofilename = data.Ipsecprofilename.ValueString()
	}
	if !data.Local.IsNull() && !data.Local.IsUnknown() {
		iptunnel.Local = data.Local.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		iptunnel.Name = data.Name.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		iptunnel.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Protocol.IsNull() && !data.Protocol.IsUnknown() {
		iptunnel.Protocol = data.Protocol.ValueString()
	}
	if !data.Remote.IsNull() && !data.Remote.IsUnknown() {
		iptunnel.Remote = data.Remote.ValueString()
	}
	if !data.Remotesubnetmask.IsNull() && !data.Remotesubnetmask.IsUnknown() {
		iptunnel.Remotesubnetmask = data.Remotesubnetmask.ValueString()
	}
	if !data.Tosinherit.IsNull() && !data.Tosinherit.IsUnknown() {
		iptunnel.Tosinherit = data.Tosinherit.ValueString()
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		iptunnel.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
	}
	if !data.Vlantagging.IsNull() && !data.Vlantagging.IsUnknown() {
		iptunnel.Vlantagging = data.Vlantagging.ValueString()
	}
	if !data.Vnid.IsNull() && !data.Vnid.IsUnknown() {
		iptunnel.Vnid = utils.IntPtr(int(data.Vnid.ValueInt64()))
	}

	return iptunnel
}

// iptunnelGetTheUpdatablePayloadFromThePlan builds the payload restricted to the
// NITRO-updatable fields (matching the SDK v2 update semantics: only destport,
// tosinherit and vlantagging are mutable in place; everything else is ForceNew).
func iptunnelGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *IptunnelResourceModel) network.Iptunnel {
	tflog.Debug(ctx, "In iptunnelGetTheUpdatablePayloadFromThePlan Function")

	iptunnel := network.Iptunnel{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		iptunnel.Name = data.Name.ValueString()
	}
	if !data.Destport.IsNull() && !data.Destport.IsUnknown() {
		iptunnel.Destport = utils.IntPtr(int(data.Destport.ValueInt64()))
	}
	if !data.Tosinherit.IsNull() && !data.Tosinherit.IsUnknown() {
		iptunnel.Tosinherit = data.Tosinherit.ValueString()
	}
	if !data.Vlantagging.IsNull() && !data.Vlantagging.IsUnknown() {
		iptunnel.Vlantagging = data.Vlantagging.ValueString()
	}

	return iptunnel
}

func iptunnelSetAttrFromGet(ctx context.Context, data *IptunnelResourceModel, getResponseData map[string]interface{}) *IptunnelResourceModel {
	tflog.Debug(ctx, "In iptunnelSetAttrFromGet Function")

	// Convert API response to model.
	// Guard the else-branches so a known/configured value is never clobbered with
	// null when NITRO omits an at-default field from GET (omit-on-default trap).
	if val, ok := getResponseData["destport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Destport = types.Int64Value(intVal)
		}
	} else if data.Destport.IsUnknown() {
		data.Destport = types.Int64Null()
	}
	if val, ok := getResponseData["grepayload"]; ok && val != nil {
		data.Grepayload = types.StringValue(val.(string))
	} else if data.Grepayload.IsUnknown() {
		data.Grepayload = types.StringNull()
	}
	if val, ok := getResponseData["ipsecprofilename"]; ok && val != nil {
		data.Ipsecprofilename = types.StringValue(val.(string))
	} else if data.Ipsecprofilename.IsUnknown() {
		data.Ipsecprofilename = types.StringNull()
	}
	if val, ok := getResponseData["local"]; ok && val != nil {
		data.Local = types.StringValue(val.(string))
	} else if data.Local.IsUnknown() {
		data.Local = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else if data.Name.IsUnknown() {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else if data.Ownergroup.IsUnknown() {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["protocol"]; ok && val != nil {
		data.Protocol = types.StringValue(val.(string))
	} else if data.Protocol.IsUnknown() {
		data.Protocol = types.StringNull()
	}
	if val, ok := getResponseData["remote"]; ok && val != nil {
		data.Remote = types.StringValue(val.(string))
	} else if data.Remote.IsUnknown() {
		data.Remote = types.StringNull()
	}
	if val, ok := getResponseData["remotesubnetmask"]; ok && val != nil {
		data.Remotesubnetmask = types.StringValue(val.(string))
	} else if data.Remotesubnetmask.IsUnknown() {
		data.Remotesubnetmask = types.StringNull()
	}
	if val, ok := getResponseData["tosinherit"]; ok && val != nil {
		data.Tosinherit = types.StringValue(val.(string))
	} else if data.Tosinherit.IsUnknown() {
		data.Tosinherit = types.StringNull()
	}
	if val, ok := getResponseData["vlan"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vlan = types.Int64Value(intVal)
		}
	} else if data.Vlan.IsUnknown() {
		data.Vlan = types.Int64Null()
	}
	if val, ok := getResponseData["vlantagging"]; ok && val != nil {
		data.Vlantagging = types.StringValue(val.(string))
	} else if data.Vlantagging.IsUnknown() {
		data.Vlantagging = types.StringNull()
	}
	if val, ok := getResponseData["vnid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Vnid = types.Int64Value(intVal)
		}
	} else if data.Vnid.IsUnknown() {
		data.Vnid = types.Int64Null()
	}

	// Set ID for the resource.
	// Single unique attribute (name) - matches the SDK v2 d.SetId(name) contract
	// and the resource_id_mapping.json "iptunnel": "name" order.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
