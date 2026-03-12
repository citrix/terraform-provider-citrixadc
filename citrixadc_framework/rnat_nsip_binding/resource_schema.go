package rnat_nsip_binding

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

// RnatNsipBindingResourceModel describes the resource data model.
type RnatNsipBindingResourceModel struct {
	Id    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Natip types.String `tfsdk:"natip"`
}

func (r *RnatNsipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rnat_nsip_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the RNAT rule to which to bind NAT IPs.",
			},
			"natip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range.",
			},
		},
	}
}

func rnat_nsip_bindingGetThePayloadFromtheConfig(ctx context.Context, data *RnatNsipBindingResourceModel) network.Rnatnsipbinding {
	tflog.Debug(ctx, "In rnat_nsip_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	rnat_nsip_binding := network.Rnatnsipbinding{}
	if !data.Name.IsNull() {
		rnat_nsip_binding.Name = data.Name.ValueString()
	}
	if !data.Natip.IsNull() {
		rnat_nsip_binding.Natip = data.Natip.ValueString()
	}

	return rnat_nsip_binding
}

func rnat_nsip_bindingSetAttrFromGet(ctx context.Context, data *RnatNsipBindingResourceModel, getResponseData map[string]interface{}) *RnatNsipBindingResourceModel {
	tflog.Debug(ctx, "In rnat_nsip_bindingSetAttrFromGet Function")

	// Convert API response to model
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

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("natip:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Natip.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
