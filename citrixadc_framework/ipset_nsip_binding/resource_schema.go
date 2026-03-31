package ipset_nsip_binding

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

// IpsetNsipBindingResourceModel describes the resource data model.
type IpsetNsipBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Name      types.String `tfsdk:"name"`
}

func (r *IpsetNsipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ipset_nsip_binding resource.",
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "One or more IP addresses bound to the IP set.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the IP set to which to bind IP addresses.",
			},
		},
	}
}

func ipset_nsip_bindingGetThePayloadFromtheConfig(ctx context.Context, data *IpsetNsipBindingResourceModel) network.Ipsetnsipbinding {
	tflog.Debug(ctx, "In ipset_nsip_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ipset_nsip_binding := network.Ipsetnsipbinding{}
	if !data.Ipaddress.IsNull() {
		ipset_nsip_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() {
		ipset_nsip_binding.Name = data.Name.ValueString()
	}

	return ipset_nsip_binding
}

func ipset_nsip_bindingSetAttrFromGet(ctx context.Context, data *IpsetNsipBindingResourceModel, getResponseData map[string]interface{}) *IpsetNsipBindingResourceModel {
	tflog.Debug(ctx, "In ipset_nsip_bindingSetAttrFromGet Function")

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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipaddress.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
