package netbridge_nsip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NetbridgeNsip6BindingResourceModel describes the resource data model.
type NetbridgeNsip6BindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Name      types.String `tfsdk:"name"`
	Netmask   types.String `tfsdk:"netmask"`
}

func (r *NetbridgeNsip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the netbridge_nsip6_binding resource.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The subnet that is extended by this network bridge.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the network bridge.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The network mask for the subnet.",
			},
		},
	}
}

func netbridge_nsip6_bindingGetThePayloadFromtheConfig(ctx context.Context, data *NetbridgeNsip6BindingResourceModel) network.Netbridgensip6binding {
	tflog.Debug(ctx, "In netbridge_nsip6_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	netbridge_nsip6_binding := network.Netbridgensip6binding{}
	if !data.Ipaddress.IsNull() {
		netbridge_nsip6_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() {
		netbridge_nsip6_binding.Name = data.Name.ValueString()
	}
	if !data.Netmask.IsNull() {
		netbridge_nsip6_binding.Netmask = data.Netmask.ValueString()
	}

	return netbridge_nsip6_binding
}

func netbridge_nsip6_bindingSetAttrFromGet(ctx context.Context, data *NetbridgeNsip6BindingResourceModel, getResponseData map[string]interface{}) *NetbridgeNsip6BindingResourceModel {
	tflog.Debug(ctx, "In netbridge_nsip6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ipaddress"]; ok && val != nil {
		data.Ipaddress = types.StringValue(val.(string))
	} else {
		data.Ipaddress = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipaddress.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
