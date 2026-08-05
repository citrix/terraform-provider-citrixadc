package ipv6

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ipv6ResourceModel describes the resource data model.
type Ipv6ResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Dodad                types.String `tfsdk:"dodad"`
	Natprefix            types.String `tfsdk:"natprefix"`
	Ndbasereachtime      types.Int64  `tfsdk:"ndbasereachtime"`
	Ndretransmissiontime types.Int64  `tfsdk:"ndretransmissiontime"`
	Ralearning           types.String `tfsdk:"ralearning"`
	Routerredirection    types.String `tfsdk:"routerredirection"`
	Td                   types.Int64  `tfsdk:"td"`
	Usipnatprefix        types.String `tfsdk:"usipnatprefix"`
}

func (r *Ipv6Resource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ipv6 resource.",
			},
			"dodad": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to do Duplicate Address\nDetection (DAD) for all the Citrix ADC owned IPv6 addresses regardless of whether they are obtained through stateless auto configuration, DHCPv6, or manual configuration.",
			},
			"natprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Prefix used for translating packets from private IPv6 servers to IPv4 packets. This prefix has a length of 96 bits (128-32 = 96). The IPv6 servers embed the destination IP address of the IPv4 servers or hosts in the last 32 bits of the destination IP address field of the IPv6 packets. The first 96 bits of the destination IP address field are set as the IPv6 NAT prefix. IPv6 packets addressed to this prefix have to be routed to the Citrix ADC to ensure that the IPv6-IPv4 translation is done by the appliance.",
			},
			"ndbasereachtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Base reachable time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, that the Citrix ADC assumes an adjacent device is reachable after receiving a reachability confirmation.",
			},
			"ndretransmissiontime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Retransmission time of the Neighbor Discovery (ND6) protocol. The time, in milliseconds, between retransmitted Neighbor Solicitation (NS) messages, to an adjacent device.",
			},
			"ralearning": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to learn about various routes from Router Advertisement (RA) and Router Solicitation (RS) messages sent by the routers.",
			},
			"routerredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the Citrix ADC to do Router Redirection.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
			"usipnatprefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPV6 NATPREFIX used in NAT46 scenario when USIP is turned on",
			},
		},
	}
}

func ipv6GetThePayloadFromtheConfig(ctx context.Context, data *Ipv6ResourceModel) network.Ipv6 {
	tflog.Debug(ctx, "In ipv6GetThePayloadFromtheConfig Function")

	// Create API request body from the model.
	// Guard against IsUnknown() so Optional+Computed attributes that are not
	// configured (and therefore unknown in the plan during Create) are not sent
	// with zero values — mirrors the SDK v2 d.GetOk() semantics.
	ipv6 := network.Ipv6{}
	if !data.Dodad.IsNull() && !data.Dodad.IsUnknown() {
		ipv6.Dodad = data.Dodad.ValueString()
	}
	if !data.Natprefix.IsNull() && !data.Natprefix.IsUnknown() {
		ipv6.Natprefix = data.Natprefix.ValueString()
	}
	if !data.Ndbasereachtime.IsNull() && !data.Ndbasereachtime.IsUnknown() {
		ipv6.Ndbasereachtime = utils.IntPtr(int(data.Ndbasereachtime.ValueInt64()))
	}
	if !data.Ndretransmissiontime.IsNull() && !data.Ndretransmissiontime.IsUnknown() {
		ipv6.Ndretransmissiontime = utils.IntPtr(int(data.Ndretransmissiontime.ValueInt64()))
	}
	if !data.Ralearning.IsNull() && !data.Ralearning.IsUnknown() {
		ipv6.Ralearning = data.Ralearning.ValueString()
	}
	if !data.Routerredirection.IsNull() && !data.Routerredirection.IsUnknown() {
		ipv6.Routerredirection = data.Routerredirection.ValueString()
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		ipv6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}
	if !data.Usipnatprefix.IsNull() && !data.Usipnatprefix.IsUnknown() {
		ipv6.Usipnatprefix = data.Usipnatprefix.ValueString()
	}

	return ipv6
}

func ipv6SetAttrFromGet(ctx context.Context, data *Ipv6ResourceModel, getResponseData map[string]interface{}) *Ipv6ResourceModel {
	tflog.Debug(ctx, "In ipv6SetAttrFromGet Function")

	// Convert API response to model.
	// omit-on-default guard: NITRO omits attributes left at their default (e.g. a
	// configured 0/false/DISABLED value). Only null an attribute when it is still
	// unknown; never clobber a known configured value that the GET simply omitted.
	if val, ok := getResponseData["dodad"]; ok && val != nil {
		data.Dodad = types.StringValue(val.(string))
	} else if data.Dodad.IsUnknown() {
		data.Dodad = types.StringNull()
	}
	if val, ok := getResponseData["natprefix"]; ok && val != nil {
		data.Natprefix = types.StringValue(val.(string))
	} else if data.Natprefix.IsUnknown() {
		data.Natprefix = types.StringNull()
	}
	if val, ok := getResponseData["ndbasereachtime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ndbasereachtime = types.Int64Value(intVal)
		}
	} else if data.Ndbasereachtime.IsUnknown() {
		data.Ndbasereachtime = types.Int64Null()
	}
	if val, ok := getResponseData["ndretransmissiontime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ndretransmissiontime = types.Int64Value(intVal)
		}
	} else if data.Ndretransmissiontime.IsUnknown() {
		data.Ndretransmissiontime = types.Int64Null()
	}
	if val, ok := getResponseData["ralearning"]; ok && val != nil {
		data.Ralearning = types.StringValue(val.(string))
	} else if data.Ralearning.IsUnknown() {
		data.Ralearning = types.StringNull()
	}
	if val, ok := getResponseData["routerredirection"]; ok && val != nil {
		data.Routerredirection = types.StringValue(val.(string))
	} else if data.Routerredirection.IsUnknown() {
		data.Routerredirection = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}
	if val, ok := getResponseData["usipnatprefix"]; ok && val != nil {
		data.Usipnatprefix = types.StringValue(val.(string))
	} else if data.Usipnatprefix.IsUnknown() {
		data.Usipnatprefix = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute (td)
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Td.ValueInt64()))

	return data
}
