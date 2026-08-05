package inatparam

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

// InatparamResourceModel describes the resource data model.
type InatparamResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Nat46fragheader   types.String `tfsdk:"nat46fragheader"`
	Nat46ignoretos    types.String `tfsdk:"nat46ignoretos"`
	Nat46v6mtu        types.Int64  `tfsdk:"nat46v6mtu"`
	Nat46v6prefix     types.String `tfsdk:"nat46v6prefix"`
	Nat46zerochecksum types.String `tfsdk:"nat46zerochecksum"`
	Td                types.Int64  `tfsdk:"td"`
}

func (r *InatparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the inatparam resource.",
			},
			"nat46fragheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When disabled, translator will not insert IPv6 fragmentation header for non fragmented IPv4 packets",
			},
			"nat46ignoretos": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ignore TOS.",
			},
			"nat46v6mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MTU setting for the IPv6 side. If the incoming IPv4 packet greater than this, either fragment or send icmp need fragmentation error.",
			},
			"nat46v6prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The prefix used for translating packets received from private IPv6 servers into IPv4 packets. This prefix has a length of 96 bits (128-32 = 96). The IPv6 servers embed the destination IP address of the IPv4 servers or hosts in the last 32 bits of the destination IP address field of the IPv6 packets. The first 96 bits of the destination IP address field are set as the IPv6 NAT prefix. IPv6 packets addressed to this prefix have to be routed to the Citrix ADC to ensure that the IPv6-IPv4 translation is done by the appliance.",
			},
			"nat46zerochecksum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Calculate checksum for UDP packets with zero checksum",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},
		},
	}
}

func inatparamGetThePayloadFromtheConfig(ctx context.Context, data *InatparamResourceModel) network.Inatparam {
	tflog.Debug(ctx, "In inatparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	inatparam := network.Inatparam{}
	if !data.Nat46fragheader.IsNull() && !data.Nat46fragheader.IsUnknown() {
		inatparam.Nat46fragheader = data.Nat46fragheader.ValueString()
	}
	if !data.Nat46ignoretos.IsNull() && !data.Nat46ignoretos.IsUnknown() {
		inatparam.Nat46ignoretos = data.Nat46ignoretos.ValueString()
	}
	if !data.Nat46v6mtu.IsNull() && !data.Nat46v6mtu.IsUnknown() {
		inatparam.Nat46v6mtu = utils.IntPtr(int(data.Nat46v6mtu.ValueInt64()))
	}
	if !data.Nat46v6prefix.IsNull() && !data.Nat46v6prefix.IsUnknown() {
		inatparam.Nat46v6prefix = data.Nat46v6prefix.ValueString()
	}
	if !data.Nat46zerochecksum.IsNull() && !data.Nat46zerochecksum.IsUnknown() {
		inatparam.Nat46zerochecksum = data.Nat46zerochecksum.ValueString()
	}
	// td is a co-dependent key: NITRO rejects a payload that carries td unless
	// nat46v6prefix is also present (errorcode 1093 "Argument pre-requisite
	// missing [td, nat46v6Prefix]"). Match the SDK v2 d.GetOk("td") semantics,
	// which omit td for the default 0: only include td when the user explicitly
	// set a non-zero value. Leaving the pointer nil marshals to "td":null, which
	// NITRO accepts and treats as the default traffic domain (0).
	if !data.Td.IsNull() && !data.Td.IsUnknown() && data.Td.ValueInt64() != 0 {
		inatparam.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return inatparam
}

func inatparamSetAttrFromGet(ctx context.Context, data *InatparamResourceModel, getResponseData map[string]interface{}) *InatparamResourceModel {
	tflog.Debug(ctx, "In inatparamSetAttrFromGet Function")

	// Convert API response to model. For Optional+Computed attributes whose
	// NITRO default may be omitted from the GET response, only null the value
	// when it was unknown (never configured); this preserves a value the user
	// explicitly configured and avoids "inconsistent result after apply".
	if val, ok := getResponseData["nat46fragheader"]; ok && val != nil {
		data.Nat46fragheader = types.StringValue(val.(string))
	} else if data.Nat46fragheader.IsUnknown() {
		data.Nat46fragheader = types.StringNull()
	}
	if val, ok := getResponseData["nat46ignoretos"]; ok && val != nil {
		data.Nat46ignoretos = types.StringValue(val.(string))
	} else if data.Nat46ignoretos.IsUnknown() {
		data.Nat46ignoretos = types.StringNull()
	}
	if val, ok := getResponseData["nat46v6mtu"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nat46v6mtu = types.Int64Value(intVal)
		}
	} else if data.Nat46v6mtu.IsUnknown() {
		data.Nat46v6mtu = types.Int64Null()
	}
	if val, ok := getResponseData["nat46v6prefix"]; ok && val != nil {
		data.Nat46v6prefix = types.StringValue(val.(string))
	} else if data.Nat46v6prefix.IsUnknown() {
		data.Nat46v6prefix = types.StringNull()
	}
	if val, ok := getResponseData["nat46zerochecksum"]; ok && val != nil {
		data.Nat46zerochecksum = types.StringValue(val.(string))
	} else if data.Nat46zerochecksum.IsUnknown() {
		data.Nat46zerochecksum = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else if data.Td.IsUnknown() {
		data.Td = types.Int64Null()
	}

	// Set ID for the resource.
	// Case 2: Single unique attribute (the traffic domain, td).
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Td.ValueInt64()))

	return data
}
