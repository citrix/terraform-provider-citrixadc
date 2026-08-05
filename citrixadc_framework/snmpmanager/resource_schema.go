package snmpmanager

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/snmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SnmpmanagerResourceModel describes the resource data model.
type SnmpmanagerResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Domainresolveretry types.Int64  `tfsdk:"domainresolveretry"`
	Ipaddress          types.String `tfsdk:"ipaddress"`
	Netmask            types.String `tfsdk:"netmask"`
}

func (r *SnmpmanagerResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "The ID of the snmpmanager resource.",
			},
			// domainresolveretry: SDK v2 TypeInt, Optional+Computed (no Default, no ForceNew) - updateable.
			"domainresolveretry": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of time, in seconds, for which the Citrix ADC waits before sending another DNS query to resolve the host name of the SNMP manager if the last query failed. This parameter is valid for host-name based SNMP managers only. After a query succeeds, the TTL determines the wait time. The minimum and default value is 5.",
			},
			// ipaddress: SDK v2 TypeString, Required + ForceNew -> RequiresReplace().
			"ipaddress": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address of the SNMP manager. Can be an IPv4 or IPv6 address. You can instead specify an IPv4 network address or IPv6 network prefix if you want the Citrix ADC to respond to SNMP queries from any device on the specified network. Alternatively, instead of an IPv4 address, you can specify a host name that has been assigned to an SNMP manager. If you do so, you must add a DNS name server that resolves the host name of the SNMP manager to its IP address. \nNote: The Citrix ADC does not support host names for SNMP managers that have IPv6 addresses.",
			},
			// netmask: SDK v2 TypeString, Optional + Computed + ForceNew ->
			// UseStateForUnknown() (Computed) + RequiresReplaceIfConfigured() (Optional+Computed+ForceNew).
			"netmask": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Subnet mask associated with an IPv4 network address. If the IP address specifies the address or host name of a specific host, accept the default value of 255.255.255.255.",
			},
		},
	}
}

// snmpmanagerGetThePayloadFromthePlan builds the NITRO create payload from the plan.
func snmpmanagerGetThePayloadFromthePlan(ctx context.Context, data *SnmpmanagerResourceModel) snmp.Snmpmanager {
	tflog.Debug(ctx, "In snmpmanagerGetThePayloadFromthePlan Function")

	snmpmanager := snmp.Snmpmanager{}
	if !data.Domainresolveretry.IsNull() && !data.Domainresolveretry.IsUnknown() {
		snmpmanager.Domainresolveretry = utils.IntPtr(int(data.Domainresolveretry.ValueInt64()))
	}
	if !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() {
		snmpmanager.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		snmpmanager.Netmask = data.Netmask.ValueString()
	}

	return snmpmanager
}

func snmpmanagerSetAttrFromGet(ctx context.Context, data *SnmpmanagerResourceModel, getResponseData map[string]interface{}) *SnmpmanagerResourceModel {
	tflog.Debug(ctx, "In snmpmanagerSetAttrFromGet Function")

	// Convert API response to model.
	// else-branches only null the value when it is unknown, never clobbering a
	// known configured value that NITRO may omit from GET (omit-on-default trap).
	if val, ok := getResponseData["domainresolveretry"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Domainresolveretry = types.Int64Value(intVal)
		}
	} else if data.Domainresolveretry.IsUnknown() {
		data.Domainresolveretry = types.Int64Null()
	}
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else if data.Ipaddress.IsUnknown() {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else if data.Netmask.IsUnknown() {
		data.Netmask = types.StringNull()
	}

	// ID scheme matches SDK v2 (d.SetId(ipaddress)) and resource_id_mapping.json
	// ("snmpmanager": "ipaddress") - single-key, plain value.
	data.Id = types.StringValue(data.Ipaddress.ValueString())

	return data
}
