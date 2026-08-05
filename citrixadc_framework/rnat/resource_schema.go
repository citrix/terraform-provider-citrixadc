package rnat

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
			// SDK v2: Optional + Computed + ForceNew.
			"aclname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An extended ACL defined for the RNAT entry.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			// SDK v2: Optional + Computed (updateable, no default).
			"connfailover": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Synchronize all connection-related information for the RNAT sessions with the secondary ADC in a high availability (HA) pair.",
			},
			// SDK v2: Required + ForceNew.
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the RNAT4 rule. Must begin with a letter, number, or the underscore character (_), and can consist of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore characters. Cannot be changed after the rule is created. Choose a name that helps identify the RNAT4 rule.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			// SDK v2: Optional + Computed + ForceNew.
			"natip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			// SDK v2: Optional + Computed + ForceNew.
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The subnet mask for the network address.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			// SDK v2: Optional + Computed + ForceNew.
			"network": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The network address defined for the RNAT entry.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
			},
			// newname drives an in-place NITRO rename (?action=rename). Not present
			// in SDK v2. Optional only: it is a pure user input never echoed by GET,
			// so Computed would cause known-after-apply churn, and RequiresReplace
			// would force recreation instead of letting the change reach Update.
			"newname": schema.StringAttribute{
				Optional:    true,
				Description: "New name for the RNAT4 rule. Must begin with an ASCII alphabetic or underscore (_) character, and must contain       only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			// SDK v2: Optional + Computed (updateable, no default).
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this rnat rule.",
			},
			// SDK v2: Optional + Computed (updateable, no default).
			"redirectport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number to which the IPv4 packets are redirected. Applicable to TCP and UDP protocols.",
			},
			// SDK v2: Optional + Computed (updateable, no default).
			"srcippersistency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables the Citrix ADC to use the same NAT IP address for all RNAT sessions initiated from a particular server.",
			},
			// SDK v2: Optional + Computed (updateable, no default).
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			// SDK v2: Optional + Computed (updateable, no default).
			"useproxyport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable source port proxying, which enables the Citrix ADC to use the RNAT ips using proxied source port.",
			},
		},
	}
}

// rnatGetThePayloadFromthePlan builds the NITRO rnat object for the add (create)
// operation. It mirrors the SDK v2 create payload plus natip (a valid NITRO add
// parameter). newname is excluded: it is rename-only.
func rnatGetThePayloadFromthePlan(ctx context.Context, data *RnatResourceModel) network.Rnat {
	tflog.Debug(ctx, "In rnatGetThePayloadFromthePlan Function")

	rnat := network.Rnat{}
	if !data.Aclname.IsNull() && !data.Aclname.IsUnknown() {
		rnat.Aclname = data.Aclname.ValueString()
	}
	if !data.Connfailover.IsNull() && !data.Connfailover.IsUnknown() {
		rnat.Connfailover = data.Connfailover.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		rnat.Name = data.Name.ValueString()
	}
	if !data.Natip.IsNull() && !data.Natip.IsUnknown() {
		rnat.Natip = data.Natip.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		rnat.Netmask = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() && !data.Network.IsUnknown() {
		rnat.Network = data.Network.ValueString()
	}
	// newname is rename-only and must not be sent in the add payload.
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		rnat.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Redirectport.IsNull() && !data.Redirectport.IsUnknown() {
		rnat.Redirectport = utils.IntPtr(int(data.Redirectport.ValueInt64()))
	}
	if !data.Srcippersistency.IsNull() && !data.Srcippersistency.IsUnknown() {
		rnat.Srcippersistency = data.Srcippersistency.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		rnat.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Useproxyport.IsNull() && !data.Useproxyport.IsUnknown() {
		rnat.Useproxyport = data.Useproxyport.ValueString()
	}

	return rnat
}

// rnatGetTheUpdatablePayloadFromThePlan builds the NITRO rnat object for the
// update operation, restricted to the SDK v2 updateable fields (name identity +
// connfailover, ownergroup, redirectport, srcippersistency, td, useproxyport).
// The ForceNew fields (aclname, natip, netmask, network) and rename-only newname
// are excluded.
func rnatGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *RnatResourceModel) network.Rnat {
	tflog.Debug(ctx, "In rnatGetTheUpdatablePayloadFromThePlan Function")

	rnat := network.Rnat{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		rnat.Name = data.Name.ValueString()
	}
	if !data.Connfailover.IsNull() && !data.Connfailover.IsUnknown() {
		rnat.Connfailover = data.Connfailover.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		rnat.Ownergroup = data.Ownergroup.ValueString()
	}
	if !data.Redirectport.IsNull() && !data.Redirectport.IsUnknown() {
		rnat.Redirectport = utils.IntPtr(int(data.Redirectport.ValueInt64()))
	}
	if !data.Srcippersistency.IsNull() && !data.Srcippersistency.IsUnknown() {
		rnat.Srcippersistency = data.Srcippersistency.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		rnat.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Useproxyport.IsNull() && !data.Useproxyport.IsUnknown() {
		rnat.Useproxyport = data.Useproxyport.ValueString()
	}

	return rnat
}

// rnatSetAttrFromGet maps a NITRO GET response onto the resource model.
//
// Backward-compat guards:
//   - The `name` key attribute is only adopted from GET when the model value is
//     null/empty (import path). Otherwise the configured value is preserved so a
//     rename does not trigger a spurious RequiresReplace diff.
//   - else-branches only null a value when it is Unknown, never clobbering a known
//     configured value that NITRO omits from GET (omit-on-default trap).
//   - newname is rename-only, never returned by GET, and is left untouched.
//   - data.Id is NOT set here; it is owned by Create/Update/Import (the live name).
func rnatSetAttrFromGet(ctx context.Context, data *RnatResourceModel, getResponseData map[string]interface{}) *RnatResourceModel {
	tflog.Debug(ctx, "In rnatSetAttrFromGet Function")

	if val, ok := getResponseData["aclname"]; ok && val != nil {
		data.Aclname = types.StringValue(val.(string))
	} else if data.Aclname.IsUnknown() {
		data.Aclname = types.StringNull()
	}
	if val, ok := getResponseData["connfailover"]; ok && val != nil {
		data.Connfailover = types.StringValue(val.(string))
	} else if data.Connfailover.IsUnknown() {
		data.Connfailover = types.StringNull()
	}
	// Preserve the configured/known name; only adopt the GET value on import.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	if val, ok := getResponseData["natip"]; ok && val != nil {
		data.Natip = types.StringValue(val.(string))
	} else if data.Natip.IsUnknown() {
		data.Natip = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else if data.Netmask.IsUnknown() {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["network"]; ok && val != nil {
		data.Network = types.StringValue(val.(string))
	} else if data.Network.IsUnknown() {
		data.Network = types.StringNull()
	}
	// newname is rename-only; never returned by GET. Leave it untouched.
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else if data.Ownergroup.IsUnknown() {
		data.Ownergroup = types.StringNull()
	}
	if val, ok := getResponseData["redirectport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Redirectport = types.Int64Value(intVal)
		}
	} else if data.Redirectport.IsUnknown() {
		data.Redirectport = types.Int64Null()
	}
	if val, ok := getResponseData["srcippersistency"]; ok && val != nil {
		data.Srcippersistency = types.StringValue(val.(string))
	} else if data.Srcippersistency.IsUnknown() {
		data.Srcippersistency = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["useproxyport"]; ok && val != nil {
		data.Useproxyport = types.StringValue(val.(string))
	} else if data.Useproxyport.IsUnknown() {
		data.Useproxyport = types.StringNull()
	}

	return data
}

// rnatSetAttrFromGetForDatasource maps a NITRO GET response onto the shared model
// for the datasource: every attribute is copied verbatim, newname is nulled
// (rename-only), and the datasource ID is set to the rnat name.
func rnatSetAttrFromGetForDatasource(ctx context.Context, data *RnatResourceModel, getResponseData map[string]interface{}) *RnatResourceModel {
	tflog.Debug(ctx, "In rnatSetAttrFromGetForDatasource Function")

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
	// newname is rename-only and never a read-back field for the datasource.
	data.Newname = types.StringNull()
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

	// Datasource ID is the rnat name.
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
