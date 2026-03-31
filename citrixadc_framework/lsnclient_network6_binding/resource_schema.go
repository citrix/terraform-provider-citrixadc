package lsnclient_network6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/lsn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LsnclientNetwork6BindingResourceModel describes the resource data model.
type LsnclientNetwork6BindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Clientname types.String `tfsdk:"clientname"`
	Netmask    types.String `tfsdk:"netmask"`
	Network    types.String `tfsdk:"network"`
	Network6   types.String `tfsdk:"network6"`
	Td         types.Int64  `tfsdk:"td"`
}

func (r *LsnclientNetwork6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lsnclient_network6_binding resource.",
			},
			"clientname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN client entity. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN client is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn client1\" or 'lsn client1').",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subnet mask for the IPv4 address specified in the Network parameter.",
			},
			"network": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 address(es) of the LSN subscriber(s) or subscriber network(s) on whose traffic you want the Citrix ADC to perform Large Scale NAT.",
			},
			"network6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 address(es) of the LSN subscriber(s) or subscriber network(s) on whose traffic you want the Citrix ADC to perform Large Scale NAT.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "ID of the traffic domain on which this subscriber or the subscriber network (as specified by the network parameter) belongs. \nIf you do not specify an ID, the subscriber or the subscriber network becomes part of the default traffic domain.",
			},
		},
	}
}

func lsnclient_network6_bindingGetThePayloadFromtheConfig(ctx context.Context, data *LsnclientNetwork6BindingResourceModel) lsn.Lsnclientnetwork6binding {
	tflog.Debug(ctx, "In lsnclient_network6_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lsnclient_network6_binding := lsn.Lsnclientnetwork6binding{}
	if !data.Clientname.IsNull() {
		lsnclient_network6_binding.Clientname = data.Clientname.ValueString()
	}
	if !data.Netmask.IsNull() {
		lsnclient_network6_binding.Netmask = data.Netmask.ValueString()
	}
	if !data.Network.IsNull() {
		lsnclient_network6_binding.Network = data.Network.ValueString()
	}
	if !data.Network6.IsNull() {
		lsnclient_network6_binding.Network6 = data.Network6.ValueString()
	}
	if !data.Td.IsNull() {
		lsnclient_network6_binding.Td = utils.IntPtr(int(data.Td.ValueInt64()))
	}

	return lsnclient_network6_binding
}

func lsnclient_network6_bindingSetAttrFromGet(ctx context.Context, data *LsnclientNetwork6BindingResourceModel, getResponseData map[string]interface{}) *LsnclientNetwork6BindingResourceModel {
	tflog.Debug(ctx, "In lsnclient_network6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clientname"]; ok && val != nil {
		data.Clientname = types.StringValue(val.(string))
	} else {
		data.Clientname = types.StringNull()
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
	if val, ok := getResponseData["network6"]; ok && val != nil {
		data.Network6 = types.StringValue(val.(string))
	} else {
		data.Network6 = types.StringNull()
	}
	if val, ok := getResponseData["td"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Td = types.Int64Value(intVal)
		}
	} else {
		data.Td = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("clientname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Clientname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("network6:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Network6.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
