package ipset_nsip6_binding

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

// IpsetNsip6BindingResourceModel describes the resource data model.
type IpsetNsip6BindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Ipaddress types.String `tfsdk:"ipaddress"`
	Name      types.String `tfsdk:"name"`
}

func (r *IpsetNsip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ipset_nsip6_binding resource.",
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

func ipset_nsip6_bindingGetThePayloadFromtheConfig(ctx context.Context, data *IpsetNsip6BindingResourceModel) network.Ipsetnsip6binding {
	tflog.Debug(ctx, "In ipset_nsip6_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ipset_nsip6_binding := network.Ipsetnsip6binding{}
	if !data.Ipaddress.IsNull() {
		ipset_nsip6_binding.Ipaddress = data.Ipaddress.ValueString()
	}
	if !data.Name.IsNull() {
		ipset_nsip6_binding.Name = data.Name.ValueString()
	}

	return ipset_nsip6_binding
}

func ipset_nsip6_bindingSetAttrFromGet(ctx context.Context, data *IpsetNsip6BindingResourceModel, getResponseData map[string]interface{}) *IpsetNsip6BindingResourceModel {
	tflog.Debug(ctx, "In ipset_nsip6_bindingSetAttrFromGet Function")

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
