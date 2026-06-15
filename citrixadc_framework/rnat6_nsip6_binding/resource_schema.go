package rnat6_nsip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Rnat6Nsip6BindingResourceModel describes the resource data model.
type Rnat6Nsip6BindingResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Natip6     types.String `tfsdk:"natip6"`
	Ownergroup types.String `tfsdk:"ownergroup"`
}

func (r *Rnat6Nsip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rnat6_nsip6_binding resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the RNAT6 rule to which to bind NAT IPs.",
			},
			"natip6": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Nat IP Address.",
			},
			"ownergroup": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The owner node group in a Cluster for this rnat rule.",
			},
		},
	}
}

func rnat6_nsip6_bindingGetThePayloadFromthePlan(ctx context.Context, data *Rnat6Nsip6BindingResourceModel) network.Rnat6nsip6binding {
	tflog.Debug(ctx, "In rnat6_nsip6_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	rnat6_nsip6_binding := network.Rnat6nsip6binding{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		rnat6_nsip6_binding.Name = data.Name.ValueString()
	}
	if !data.Natip6.IsNull() && !data.Natip6.IsUnknown() {
		rnat6_nsip6_binding.Natip6 = data.Natip6.ValueString()
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() {
		rnat6_nsip6_binding.Ownergroup = data.Ownergroup.ValueString()
	}

	return rnat6_nsip6_binding
}

func rnat6_nsip6_bindingSetAttrFromGet(ctx context.Context, data *Rnat6Nsip6BindingResourceModel, getResponseData map[string]interface{}) *Rnat6Nsip6BindingResourceModel {
	tflog.Debug(ctx, "In rnat6_nsip6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["natip6"]; ok && val != nil {
		data.Natip6 = types.StringValue(val.(string))
	} else {
		data.Natip6 = types.StringNull()
	}
	if val, ok := getResponseData["ownergroup"]; ok && val != nil {
		data.Ownergroup = types.StringValue(val.(string))
	} else {
		data.Ownergroup = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("natip6:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Natip6.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ownergroup:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ownergroup.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
