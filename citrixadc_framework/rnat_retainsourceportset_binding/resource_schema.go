package rnat_retainsourceportset_binding

import (
	"context"
	"strings"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// RnatRetainsourceportsetBindingResourceModel describes the resource data model.
type RnatRetainsourceportsetBindingResourceModel struct {
	Id types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Retainsourceportrange types.String `tfsdk:"retainsourceportrange"`
}

func (r *RnatRetainsourceportsetBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rnat_retainsourceportset_binding resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the RNAT rule to which to bind NAT IPs.",
			},
			"retainsourceportrange": schema.StringAttribute{
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "When the source port range is configured and associated with the RNAT rule, Citrix ADC will choose a port from the specified source port range configured for connection establishment at the backend servers.",
			},
		},
	}
}

func rnat_retainsourceportset_bindingGetThePayloadFromthePlan(ctx context.Context, data *RnatRetainsourceportsetBindingResourceModel) network.Rnatretainsourceportsetbinding {
	tflog.Debug(ctx, "In rnat_retainsourceportset_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	rnat_retainsourceportset_binding := network.Rnatretainsourceportsetbinding{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		rnat_retainsourceportset_binding.Name = data.Name.ValueString()
	}
	if !data.Retainsourceportrange.IsNull() && !data.Retainsourceportrange.IsUnknown() {
		rnat_retainsourceportset_binding.Retainsourceportrange = data.Retainsourceportrange.ValueString()
	}

	return rnat_retainsourceportset_binding
}

func rnat_retainsourceportset_bindingSetAttrFromGet(ctx context.Context, data *RnatRetainsourceportsetBindingResourceModel, getResponseData map[string]interface{}) *RnatRetainsourceportsetBindingResourceModel {
	tflog.Debug(ctx, "In rnat_retainsourceportset_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["retainsourceportrange"]; ok && val != nil {
		data.Retainsourceportrange = types.StringValue(val.(string))
	} else {
		data.Retainsourceportrange = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("retainsourceportrange:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Retainsourceportrange.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}